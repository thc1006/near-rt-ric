// Copyright 2017 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/xsrftoken"
	utilnet "k8s.io/apimachinery/pkg/util/net"

	"github.com/kubernetes/dashboard/src/app/backend/args"
	authApi "github.com/kubernetes/dashboard/src/app/backend/auth/api"
	clientapi "github.com/kubernetes/dashboard/src/app/backend/client/api"
	"github.com/kubernetes/dashboard/src/app/backend/errors"
)

const (
	originalForwardedForHeader = "X-Original-Forwarded-For"
	forwardedForHeader         = "X-Forwarded-For"
	realIPHeader               = "X-Real-Ip"
)

// InstallFilters installs defined filter for given web service
func InstallFilters(ws *restful.WebService, manager clientapi.ClientManager) {
	ws.Filter(requestAndResponseLogger)
	ws.Filter(metricsFilter)
	ws.Filter(validateXSRFFilter(manager.CSRFKey()))
	ws.Filter(restrictedResourcesFilter)
}

// Filter used to restrict access to dashboard exclusive resource, i.e. secret used to store dashboard encryption key.
func restrictedResourcesFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	if !authApi.ShouldRejectRequest(request.Request.URL.String()) {
		chain.ProcessFilter(request, response)
		return
	}

	err := errors.NewUnauthorized(errors.MsgDashboardExclusiveResourceError)
	response.WriteHeaderAndEntity(int(err.ErrStatus.Code), err.Error())
}

// web-service filter function used for request and response logging.
func requestAndResponseLogger(request *restful.Request, response *restful.Response,
	chain *restful.FilterChain) {
	if level, err := logrus.ParseLevel(strings.ToLower(args.Holder.GetAPILogLevel())); err == nil {
		logrus.SetLevel(level)
	}

	entry := logrus.WithFields(logrus.Fields{
		"proto":       request.Request.Proto,
		"method":      request.Request.Method,
		"remote_addr": getRemoteAddr(request.Request),
	})

	if request.Request.URL != nil {
		entry = entry.WithField("uri", request.Request.URL.RequestURI())
	}

	if request.Request.Body != nil {
		byteArr, err := io.ReadAll(request.Request.Body)
		if err == nil {
			// Restore request body so we can read it again in regular request handlers
			request.Request.Body = io.NopCloser(bytes.NewReader(byteArr))

			content := string(byteArr)
			if checkSensitiveURL(request.Request.URL.RequestURI()) {
				content = "{ contents hidden }"
			}

			entry = entry.WithField("body", content)
		}
	}

	entry.Info("Incoming request")

	chain.ProcessFilter(request, response)

	entry.WithFields(logrus.Fields{
		"status_code": response.StatusCode(),
	}).Info("Outgoing response")
}

// checkSensitiveURL checks if a string matches against a sensitive URL
// true if sensitive. false if not.
func checkSensitiveURL(url string) bool {
	sensitiveURLs := []string{"/api/v1/login", "/api/v1/csrftoken/login", "/api/v1/token/refresh"}
	for _, sensitiveURL := range sensitiveURLs {
		if strings.HasPrefix(url, sensitiveURL) {
			return true
		}
	}
	return false
}

func metricsFilter(req *restful.Request, resp *restful.Response,
	chain *restful.FilterChain) {
	resource := mapUrlToResource(req.SelectedRoutePath())
	httpClient := utilnet.GetHTTPClient(req.Request)

	chain.ProcessFilter(req, resp)

	if resource != nil {
		monitor(
			req.Request.Method,
			*resource, httpClient,
			resp.Header().Get("Content-Type"),
			resp.StatusCode(),
			time.Now(),
		)
	}
}

func validateXSRFFilter(csrfKey string) restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		resource := mapUrlToResource(req.SelectedRoutePath())

		if resource == nil || (shouldDoCsrfValidation(req) &&
			!xsrftoken.Valid(req.HeaderParameter("X-CSRF-TOKEN"), csrfKey, "none",
				*resource)) {
			err := errors.NewInvalid("CSRF validation failed")
			logrus.Error(err)
			resp.AddHeader("Content-Type", "text/plain")
			resp.WriteErrorString(http.StatusUnauthorized, err.Error()+"\n")
			return
		}

		chain.ProcessFilter(req, resp)
	}
}

// Post requests should set correct X-CSRF-TOKEN header, all other requests
// should either not edit anything or be already safe to CSRF attacks (PUT
// and DELETE)
func shouldDoCsrfValidation(req *restful.Request) bool {
	if req.Request.Method != http.MethodPost {
		return false
	}

	// Validation handlers are idempotent functions, and not actual data
	// modification operations
	if strings.HasPrefix(req.SelectedRoutePath(), "/api/v1/appdeployment/validate/") {
		return false
	}

	return true
}

// mapUrlToResource extracts the resource from the URL path /api/v1/<resource>.
// Ignores potential subresources.
func mapUrlToResource(url string) *string {
	parts := strings.Split(url, "/")
	if len(parts) < 3 {
		return nil
	}
	return &parts[3]
}

// getRemoteAddr extracts the remote address of the request, taking into
// account proxy headers.
func getRemoteAddr(r *http.Request) string {
	if ip := getRemoteIPFromForwardHeader(r, originalForwardedForHeader); ip != "" {
		return ip
	}

	if ip := getRemoteIPFromForwardHeader(r, forwardedForHeader); ip != "" {
		return ip
	}

	if realIP := strings.TrimSpace(r.Header.Get(realIPHeader)); realIP != "" {
		return realIP
	}

	return r.RemoteAddr
}

func getRemoteIPFromForwardHeader(r *http.Request, header string) string {
	ips := strings.Split(r.Header.Get(header), ",")
	return strings.TrimSpace(ips[0])
}
