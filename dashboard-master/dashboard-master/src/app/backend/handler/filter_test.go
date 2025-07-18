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
	"flag"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRequestAndResponseLogger(t *testing.T) {
	flag.Set("api-log-level", "INFO")

	var buf bytes.Buffer
	logrus.SetOutput(&buf)

	req, _ := http.NewRequest("GET", "/api/v1/login", nil)
	req.Header.Set("Content-Length", "0")
	restfulReq := restful.NewRequest(req)

	resp := httptest.NewRecorder()
	restfulResp := restful.NewResponse(resp)

	chain := &restful.FilterChain{
		Target: func(req *restful.Request, resp *restful.Response) {
			// do nothing
		},
	}

	requestAndResponseLogger(restfulReq, restfulResp, chain)

	logOutput := buf.String()
	assert.True(t, strings.Contains(logOutput, "Incoming request"))
	assert.True(t, strings.Contains(logOutput, "Outgoing response"))
	assert.True(t, strings.Contains(logOutput, "/api/v1/login"))
}
