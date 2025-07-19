
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
	"github.com/emicklei/go-restful/v3"
	clientapi "github.com/kubernetes/dashboard/src/app/backend/client/api"
	"github.com/kubernetes/dashboard/src/app/backend/errors"
	authv1 "k8s.io/api/authorization/v1"
)

func Authorize(clientManager clientapi.ClientManager, req *restful.Request, resourceAttributes *authv1.ResourceAttributes) error {
	isAuthorized, err := clientManager.CanI(req, &authv1.SelfSubjectAccessReview{
		Spec: authv1.SelfSubjectAccessReviewSpec{
			ResourceAttributes: resourceAttributes,
		},
	})

	if err != nil {
		return err
	}

	if !isAuthorized {
		return errors.NewForbidden("user is not authorized to access this resource")
	}

	return nil
}
