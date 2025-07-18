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
	"io"

	restful "github.com/emicklei/go-restful/v3"

	"github.com/kubernetes/dashboard/src/app/backend/errors"
)

func handleDownload(response *restful.Response, request *restful.Request, result io.ReadCloser) {
	response.AddHeader(restful.HEADER_ContentType, "text/plain")
	defer result.Close()
	_, err := io.Copy(response, result)
	if err != nil {
		errors.HandleInternalError(response, request, err)
	}
}
