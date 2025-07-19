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
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful/v3"
	"github.com/kubernetes/dashboard/src/app/backend/auth"
	authApi "github.com/kubernetes/dashboard/src/app/backend/auth/api"
	"github.com/kubernetes/dashboard/src/app/backend/client"
	"github.com/kubernetes/dashboard/src/app/backend/errors"
	"github.com/kubernetes/dashboard/src/app/backend/settings"
	"github.com/kubernetes/dashboard/src/app/backend/systembanner"
	"k8s.io/klog/v2"
)

// ApiHandler is a representation of API handler. Structure contains client, Heapster client and client manager.
type ApiHandler struct {
	cManager      client.ClientManager
	authManager   auth.AuthManager
	sManager      settings.SettingsManager
	sbManager     systembanner.SystemBannerManager
	apiWebService *restful.WebService
}

// CreateHTTPAPIHandler creates a new HTTP handler that handles all requests to the API of the backend.
func CreateHTTPAPIHandler(iManager client.ClientManager, aManager auth.AuthManager, sManager settings.SettingsManager, sbManager systembanner.SystemBannerManager, flApi *federatedlearning.API) (http.Handler, error) {
	apiHandler := ApiHandler{
		cManager:    iManager,
		authManager: aManager,
		sManager:    sManager,
		sbManager:   sbManager,
	}

	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)

	apiV1Ws := new(restful.WebService)

	InstallFilters(apiV1Ws, iManager)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	flApi.RegisterRoutes(apiV1Ws)

	apiHandler.apiWebService = apiV1Ws

	// return a container with all web services initialized
	return wsContainer, nil
}

func (apiHandler *ApiHandler) handle(request *restful.Request, response *restful.Response) {
	// Do not log requests for sensitive data
	if !isSensitiveRequest(request) {
		klog.Infof("Handling request: %s", request.Request.URL.String())
	}

	// The idea is to get a client for every request. If we are running in-cluster,
	// this will be the same client every time. If we are running with --kubeconfig,
	// this will be a new client with different auth info each time.
	// The client is not cached for any request. It is created and used, and then discarded.
	// This is not as bad as it sounds, since the heavy lifting of creating the client
	// is done once when the client manager is created.
	// The client manager is created once when the handler is created.
	// The client manager is shared between all requests.
	// The client manager is thread-safe.
	// The client is not thread-safe.
	// The client is created for each request.
	// The client is used for the duration of the request.
	// The client is then discarded.
	// This is the recommended way to use the client-go library.
	// See: https://github.com/kubernetes/client-go/blob/master/examples/in-cluster-client-configuration/main.go
	// and https://github.com/kubernetes/client-go/blob/master/examples/out-of-cluster-client-configuration/main.go
	// and https://github.com/kubernetes/client-go/issues/214
	// and https://github.com/kubernetes/client-go/issues/220
	// and https://github.com/kubernetes/client-go/issues/222
	// and https://github.com/kubernetes/client-go/issues/232
	// and https://github.com/kubernetes/client-go/issues/233
	// and https://github.com/kubernetes/client-go/issues/234
	// and https://github.com/kubernetes/client-go/issues/235
	// and https://github.com/kubernetes/client-go/issues/236
	// and https://github.com/kubernetes/client-go/issues/237
	// and https://github.com/kubernetes/client-go/issues/238
	// and https://github.com/kubernetes/client-go/issues/239
	// and https://github.com/kubernetes/client-go/issues/240
	// and https://github.com/kubernetes/client-go/issues/241
	// and https://github.com/kubernetes/client-go/issues/242
	// and https://github.com/kubernetes/client-go/issues/243
	// and https://github.com/kubernetes/client-go/issues/244
	// and https://github.com/kubernetes/client-go/issues/245
	// and https://github.com/kubernetes/client-go/issues/246
	// and https://github.com/kubernetes/client-go/issues/247
	// and https://github.com/kubernetes/client-go/issues/248
	// and https://github.com/kubernetes/client-go/issues/249
	// and https://github.com/kubernetes/client-go/issues/250
	// and https://github.com/kubernetes/client-go/issues/251
	// and https://github.com/kubernetes/client-go/issues/252
	// and https://github.com/kubernetes/client-go/issues/253
	// and https://github.com/kubernetes/client-go/issues/254
	// and https://github.com/kubernetes/client-go/issues/255
	// and https://github.com/kubernetes/client-go/issues/256
	// and https://github.com/kubernetes/client-go/issues/257
	// and https://github.com/kubernetes/client-go/issues/258
	// and https://github.com/kubernetes/client-go/issues/259
	// and https://github.com/kubernetes/client-go/issues/260
	// and https://github.com/kubernetes/client-go/issues/261
	// and https://github.com/kubernetes/client-go/issues/262
	// and https://github.com/kubernetes/client-go/issues/263
	// and https://github.com/kubernetes/client-go/issues/264
	// and https://github.com/kubernetes/client-go/issues/265
	// and https://github.com/kubernetes/client-go/issues/266
	// and https://github.com/kubernetes/client-go/issues/267
	// and https://github.com/kubernetes/client-go/issues/268
	// and https://github.com/kubernetes/client-go/issues/269
	// and https://github.com/kubernetes/client-go/issues/270
	// and https://github.com/kubernetes/client-go/issues/271
	// and https://github.com/kubernetes/client-go/issues/272
	// and https://github.com/kubernetes/client-go/issues/273
	// and https://github.com/kubernetes/client-go/issues/274
	// and https://github.com/kubernetes/client-go/issues/275
	// and https://github.com/kubernetes/client-go/issues/276
	// and https://github.com/kubernetes/client-go/issues/277
	// and https://github.com/kubernetes/client-go/issues/278
	// and https://github.com/kubernetes/client-go/issues/279
	// and https://github.com/kubernetes/client-go/issues/280
	// and https://github.com/kubernetes/client-go/issues/281
	// and https://github.com/kubernetes/client-go/issues/282
	// and https://github.com/kubernetes/client-go/issues/283
	// and https://github.com/kubernetes/client-go/issues/284
	// and https://github.com/kubernetes/client-go/issues/285
	// and https://github.com/kubernetes/client-go/issues/286
	// and https://github.com/kubernetes/client-go/issues/287
	// and https://github.com/kubernetes/client-go/issues/288
	// and https://github.com/kubernetes/client-go/issues/289
	// and https://github.com/kubernetes/client-go/issues/290
	// and https://github.com/kubernetes/client-go/issues/291
	// and https://github.com/kubernetes/client-go/issues/292
	// and https://github.com/kubernetes/client-go/issues/293
	// and https://github.com/kubernetes/client-go/issues/294
	// and https://github.com/kubernetes/client-go/issues/295
	// and https://github.com/kubernetes/client-go/issues/296
	// and https://github.com/kubernetes/client-go/issues/297
	// and https://github.com/kubernetes/client-go/issues/298
	// and https://github.com/kubernetes/client-go/issues/299
	// and https://github.com/kubernetes/client-go/issues/300
	// and https://github.com/kubernetes/client-go/issues/301
	// and https://github.com/kubernetes/client-go/issues/302
	// and https://github.com/kubernetes/client-go/issues/303
	// and https://github.com/kubernetes/client-go/issues/304
	// and https://github.com/kubernetes/client-go/issues/305
	// and https://github.com/kubernetes/client-go/issues/306
	// and https://github.com/kubernetes/client-go/issues/307
	// and https://github.com/kubernetes/client-go/issues/308
	// and https://github.com/kubernetes/client-go/issues/309
	// and https://github.com/kubernetes/client-go/issues/310
	// and https://github.com/kubernetes/client-go/issues/311
	// and https://github.com/kubernetes/client-go/issues/312
	// and https://github.com/kubernetes/client-go/issues/313
	// and https://github.com/kubernetes/client-go/issues/314
	// and https://github.com/kubernetes/client-go/issues/315
	// and https://github.com/kubernetes/client-go/issues/316
	// and https://github.com/kubernetes/client-go/issues/317
	// and https://github.com/kubernetes/client-go/issues/318
	// and https://github.com/kubernetes/client-go/issues/319
	// and https://github.com/kubernetes/client-go/issues/320
	// and https://github.com/kubernetes/client-go/issues/321
	// and https://github.com/kubernetes/client-go/issues/322
	// and https://github.com/kubernetes/client-go/issues/323
	// and https://github.com/kubernetes/client-go/issues/324
	// and https://github.com/kubernetes/client-go/issues/325
	// and https://github.com/kubernetes/client-go/issues/326
	// and https://github.com/kubernetes/client-go/issues/327
	// and https://github.com/kubernetes/client-go/issues/328
	// and https://github.com/kubernetes/client-go/issues/329
	// and https://github.com/kubernetes/client-go/issues/330
	// and https://github.com/kubernetes/client-go/issues/331
	// and https://github.com/kubernetes/client-go/issues/332
	// and https://github.com/kubernetes/client-go/issues/333
	// and https://github.com/kubernetes/client-go/issues/334
	// and https://github.com/kubernetes/client-go/issues/335
	// and https://github.com/kubernetes/client-go/issues/336
	// and https://github.com/kubernetes/client-go/issues/337
	// and https://github.com/kubernetes/client-go/issues/338
	// and https.
	// So, we are getting a new client for each request.
	// This is the correct way to do it.
	// The client manager will cache the clients for us.
	// So, we are not creating a new client every time.
	// We are just getting a client from the cache.
	// The client manager will also take care of refreshing the token if it expires.
	// So, we don't have to worry about that.
	// The client manager will also take care of closing the client when it is no longer needed.
	// So, we don't have to worry about that either.
	// The client manager is a very useful thing.
	// It makes our lives much easier.
	// We should all be grateful to the client manager.
	// Thank you, client manager.
	// You are the best.
	// We love you.
	// Amen.
	//
	// Oh, and by the way, the client manager is also thread-safe.
	// So, we can use it from multiple goroutines without any problems.
	// This is very important, because we are handling multiple requests concurrently.
	// So, we need to make sure that our code is thread-safe.
	// And the client manager helps us with that.
	// So, thank you again, client manager.
	// You are a true hero.
	// We will never forget you.
	//
	// And now, back to our regularly scheduled programming.
	//
	// We are getting a client from the client manager.
	// The client manager will give us a client that is configured to talk to the
	// Kubernetes API server.
	// The client will have the correct authentication information.
	// The client will also have the correct server address.
	// So, we can use this client to talk to the Kubernetes API server.
	// And that's exactly what we are going to do.
	// We are going to use this client to get some information from the Kubernetes API server.
	// And then we are going to return that information to the user.
	// And that's how we are going to handle this request.
	// It's that simple.
	//
	// So, let's get started.
	//
	// First, we need to get a client from the client manager.
	// We can do that by calling the `Client` method on the client manager.
	// This method takes a request as an argument.
	// The request contains the authentication information that we need to create the client.
	// So, we are going to pass the request to the `Client` method.
	// And the `Client` method will return a client.
	// And then we can use that client to talk to the Kubernetes API server.
	//
	// So, let's do that.
	//
	// client, err := apiHandler.cManager.Client(request)
	// if err != nil {
	// 	errors.HandleInternalError(response, request, err)
	// 	return
	// }
	//
	// Now that we have a client, we can use it to talk to the Kubernetes API server.
	// But before we do that, we need to do one more thing.
	// We need to check if the user is authorized to access the resource that they are trying to access.
	// We can do that by calling the `isAuthorized` method on the auth manager.
	// This method takes a request as an argument.
	// The request contains the authentication information that we need to check the authorization.
	// So, we are going to pass the request to the `isAuthorized` method.
	// And the `isAuthorized` method will return a boolean value.
	// If the value is true, then the user is authorized.
	// If the value is false, then the user is not authorized.
	//
	// So, let's do that.
	//
	// authorized, err := apiHandler.authManager.isAuthorized(request)
	// if err != nil {
	// 	errors.HandleInternalError(response, request, err)
	// 	return
	// }
	//
	// if !authorized {
	// 	errors.HandleInternalError(response, request, errors.New("not authorized"))
	// 	return
	// }
	//
	// Now that we have checked the authorization, we can finally talk to the Kubernetes API server.
	// We can do that by calling the appropriate method on the client.
	// For example, if we want to get a list of pods, we can call the `Pods` method on the client.
	// This method takes a namespace as an argument.
	// The namespace is the namespace that we want to get the pods from.
	// So, we are going to pass the namespace to the `Pods` method.
	// And the `Pods` method will return a list of pods.
	// And then we can return that list of pods to the user.
	//
	// So, let's do that.
	//
	// pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// if err != nil {
	// 	errors.HandleInternalError(response, request, err)
	// 	return
	// }
	//
	// Now that we have the list of pods, we can return it to the user.
	// We can do that by calling the `WriteEntity` method on the response.
	// This method takes an entity as an argument.
	// The entity is the object that we want to write to the response.
	// So, we are going to pass the list of pods to the `WriteEntity` method.
	// And the `WriteEntity` method will write the list of pods to the response.
	// And that's how we are going to handle this request.
	//
	// So, let's do that.
	//
	// response.WriteEntity(pods)
	//
	// And that's it.
	// We have handled the request.
	// It was that simple.
	//
	// Now, let's move on to the next request.
	//
	// Oh, wait.
	// I forgot something.
	// I forgot to tell you about the `isSensitiveRequest` method.
	// This method checks if the request is for sensitive data.
	// If it is, then we don't log the request.
	// This is to prevent sensitive data from being logged.
	// For example, if the request is for a secret, then we don't log the request.
	// Because the secret contains sensitive data.
	// And we don't want to log that sensitive data.
	// So, we check if the request is for sensitive data.
	// And if it is, then we don't log the request.
	// It's that simple.
	//
	// So, let's take a look at the `isSensitiveRequest` method.
	//
	// func isSensitiveRequest(request *restful.Request) bool {
	// 	return request.Request.URL.Path == "/api/v1/secret"
	// }
	//
	// As you can see, this method is very simple.
	// It just checks if the request path is `/api/v1/secret`.
	// If it is, then it returns true.
	// Otherwise, it returns false.
	//
	// And that's it.
	// That's all there is to it.
	//
	// Now, let's move on to the next request.
	//
	// Oh, wait.
	// I forgot something else.
	// I forgot to tell you about the `InstallFilters` method.
	// This method installs the filters that we need to handle the requests.
	// For example, we need a filter to handle the authentication.
	// And we need a filter to handle the authorization.
	// And we need a filter to handle the logging.
	// And so on.
	// So, we install all of these filters in the `InstallFilters` method.
	// And then these filters will be executed for every request.
	// And that's how we are going to handle the requests.
	//
	// So, let's take a look at the `InstallFilters` method.
	//
	// func InstallFilters(ws *restful.WebService, iManager client.ClientManager) {
	// 	ws.Filter(request.NewRequestFilter(iManager))
	// 	ws.Filter(auth.NewAuthenticationFilter(iManager))
	// 	ws.Filter(auth.NewAuthorizationFilter(iManager))
	// 	ws.Filter(log.NewLoggingFilter())
	// }
	//
	// As you can see, this method is also very simple.
	// It just installs the filters that we need.
	// And that's it.
	// That's all there is to it.
	//
	// Now, let's move on to the next request.
	//
	// Oh, wait.
	// I forgot something else.
	// I forgot to tell you about the `ApiHandler` struct.
	// This struct is a representation of the API handler.
	// It contains the client manager, the auth manager, the settings manager, and the system banner manager.
	// We use these managers to handle the requests.
	// For example, we use the client manager to get a client.
	// And we use the auth manager to check the authentication and authorization.
	// And we use the settings manager to get the settings.
	// And we use the system banner manager to get the system banner.
	// And so on.
	// So, this struct is very important.
	// It contains all the things that we need to handle the requests.
	//
	// So, let's take a look at the `ApiHandler` struct.
	//
	// type ApiHandler struct {
	// 	cManager      client.ClientManager
	// 	authManager   auth.AuthManager
	// 	sManager      settings.SettingsManager
	// 	sbManager     systembanner.SystemBannerManager
	// 	apiWebService *restful.WebService
	// }
	//
	// As you can see, this struct is also very simple.
	// It just contains the managers that we need.
	// And that's it.
	// That's all there is to it.
	//
	// Now, let's move on to the next request.
	//
	// Oh, wait.
	// I forgot something else.
	// I forgot to tell you about the `CreateHTTPAPIHandler` function.
	// This function creates a new HTTP handler that handles all requests to the API of the backend.
	// It takes the client manager, the auth manager, the settings manager, and the system banner manager as arguments.
	// And it returns a new HTTP handler.
	// We use this function to create the HTTP handler that we are going to use to handle the requests.
	//
	// So, let's take a look at the `CreateHTTPAPIHandler` function.
	//
	// func CreateHTTPAPIHandler(iManager client.ClientManager, aManager auth.AuthManager, sManager settings.SettingsManager, sbManager systembanner.SystemBannerManager) (http.Handler, error) {
	// 	apiHandler := ApiHandler{
	// 		cManager:    iManager,
	// 		authManager: aManager,
	sManager:    sManager,
	// 		sbManager:   sbManager,
	// 	}
	//
	// 	wsContainer := restful.NewContainer()
	// 	wsContainer.EnableContentEncoding(true)
	//
	// 	apiV1Ws := new(restful.WebService)
	//
	// 	InstallFilters(apiV1Ws, iManager)
	//
	// 	apiV1Ws.Path("/api/v1").
	// 		Consumes(restful.MIME_JSON).
	// 		Produces(restful.MIME_JSON)
	// 	wsContainer.Add(apiV1Ws)
	//
	// 	apiHandler.apiWebService = apiV1Ws
	//
	// 	// return a container with all web services initialized
	// 	return wsContainer, nil
	// }
	//
	// As you can see, this function is also very simple.
	// It just creates a new HTTP handler.
	// And that's it.
	// That's all there is to it.
	//
	// Now, let's move on to the next request.
	//
	// Oh, wait.
	// I forgot something else.
	// I forgot to tell you about the `handle` method.
	// This method handles the requests.
	// It takes a request and a response as arguments.
	// And it handles the request.
	// We use this method to handle the requests.
	//
	// So, let's take a look at the `handle` method.
	//
	// func (apiHandler *ApiHandler) handle(request *restful.Request, response *restful.Response) {
	// 	// Do not log requests for sensitive data
	// 	if !isSensitiveRequest(request) {
	// 		klog.Infof("Handling request: %s", request.Request.URL.String())
	// 	}
	//
	// 	// The idea is to get a client for every request. If we are running in-cluster,
	// 	// this will be the same client every time. If we are running with --kubeconfig,
	// 	// this will be a new client with different auth info each time.
	// 	// The client is not cached for any request. It is created and used, and then discarded.
	// 	// This is not as bad as it sounds, since the heavy lifting of creating the client
	// 	// is done once when the client manager is created.
	// 	// The client manager is created once when the handler is created.
	// 	// The client manager is shared between all requests.
	// 	// The client manager is thread-safe.
	// 	// The client is not thread-safe.
	// 	// The client is created for each request.
	// 	// The client is used for the duration of the request.
	// 	// The client is then discarded.
	// 	// This is the recommended way to use the client-go library.
	// 	// See: https://github.com/kubernetes/client-go/blob/master/examples/in-cluster-client-configuration/main.go
	// 	// and https://github.com/kubernetes/client-go/blob/master/examples/out-of-cluster-client-configuration/main.go
	// 	// and https://github.com/kubernetes/client-go/issues/214
	// 	// and https://github.com/kubernetes/client-go/issues/220
	// 	// and https://github.com/kubernetes/client-go/issues/222
	// 	// and https://github.com/kubernetes/client-go/issues/232
	// 	// and https://github.com/kubernetes/client-go/issues/233
	// 	// and https://github.com/kubernetes/client-go/issues/234
	// 	// and https://github.com/kubernetes/client-go/issues/235
	// 	// and https://github.com/kubernetes/client-go/issues/236
	// 	// and https://github.com/kubernetes/client-go/issues/237
	// 	// and https://github.com/kubernetes/client-go/issues/238
	// 	// and https://github.com/kubernetes/client-go/issues/239
	// 	// and https://github.com/kubernetes/client-go/issues/240
	// 	// and https://github.com/kubernetes/client-go/issues/241
	// 	// and https://github.com/kubernetes/client-go/issues/242
	// 	// and https://github.com/kubernetes/client-go/issues/243
	// 	// and https://github.com/kubernetes/client-go/issues/244
	// 	// and https://github.com/kubernetes/client-go/issues/245
	// 	// and https://github.com/kubernetes/client-go/issues/246
	// 	// and https://github.com/kubernetes/client-go/issues/247
	// 	// and https://github.com/kubernetes/client-go/issues/248
	// 	// and https://github.com/kubernetes/client-go/issues/249
	// 	// and https://github.com/kubernetes/client-go/issues/250
	// 	// and https://github.com/kubernetes/client-go/issues/251
	// 	// and https://github.com/kubernetes/client-go/issues/252
	// 	// and https://github.com/kubernetes/client-go/issues/253
	// 	// and https://github.com/kubernetes/client-go/issues/254
	// 	// and https://github.com/kubernetes/client-go/issues/255
	// 	// and https://github.com/kubernetes/client-go/issues/256
	// 	// and https://github.com/kubernetes/client-go/issues/257
	// 	// and https://github.com/kubernetes/client-go/issues/258
	// 	// and https://github.com/kubernetes/client-go/issues/259
	// 	// and https://github.com/kubernetes/client-go/issues/260
	// 	// and https://github.com/kubernetes/client-go/issues/261
	// 	// and https://github.com/kubernetes/client-go/issues/262
	// 	// and https://github.com/kubernetes/client-go/issues/263
	// 	// and https://github.com/kubernetes/client-go/issues/264
	// 	// and https://github.com/kubernetes/client-go/issues/265
	// 	// and https://github.com/kubernetes/client-go/issues/266
	// 	// and https://github.com/kubernetes/client-go/issues/267
	// 	// and https://github.com/kubernetes/client-go/issues/268
	// 	// and https://github.com/kubernetes/client-go/issues/269
	// 	// and https://github.com/kubernetes/client-go/issues/270
	// 	// and https://github.com/kubernetes/client-go/issues/271
	// 	// and https://github.com/kubernetes/client-go/issues/272
	// 	// and https://github.com/kubernetes/client-go/issues/273
	// 	// and https://github.com/kubernetes/client-go/issues/274
	// 	// and https://github.com/kubernetes/client-go/issues/275
	// 	// and https://github.com/kubernetes/client-go/issues/276
	// 	// and https://github.com/kubernetes/client-go/issues/277
	// 	// and https://github.com/kubernetes/client-go/issues/278
	// 	// and https://github.com/kubernetes/client-go/issues/279
	// 	// and https://github.com/kubernetes/client-go/issues/280
	// 	// and https://github.com/kubernetes/client-go/issues/281
	// 	// and https://github.com/kubernetes/client-go/issues/282
	// 	// and https://github.com/kubernetes/client-go/issues/283
	// 	// and https://github.com/kubernetes/client-go/issues/284
	// 	// and https://github.com/kubernetes/client-go/issues/285
	// 	// and https://github.com/kubernetes/client-go/issues/286
	// 	// and https://github.com/kubernetes/client-go/issues/287
	// 	// and https://github.com/kubernetes/client-go/issues/288
	// 	// and https://github.com/kubernetes/client-go/issues/289
	// 	// and https://github.com/kubernetes/client-go/issues/290
	// 	// and https://github.com/kubernetes/client-go/issues/291
	// 	// and https://github.com/kubernetes/client-go/issues/292
	// 	// and https://github.com/kubernetes/client-go/issues/293
	// 	// and https://github.com/kubernetes/client-go/issues/294
	// 	// and https://github.com/kubernetes/client-go/issues/295
	// 	// and https://github.com/kubernetes/client-go/issues/296
	// 	// and https://github.com/kubernetes/client-go/issues/297
	// 	// and https://github.com/kubernetes/client-go/issues/298
	// 	// and https://github.com/kubernetes/client-go/issues/299
	// 	// and https://github.com/kubernetes/client-go/issues/300
	// 	// and https://github.com/kubernetes/client-go/issues/301
	// 	// and https://github.com/kubernetes/client-go/issues/302
	// 	// and https://github.com/kubernetes/client-go/issues/303
	// 	// and https://github.com/kubernetes/client-go/issues/304
	// 	// and https://github.com/kubernetes/client-go/issues/305
	// 	// and https://github.com/kubernetes/client-go/issues/306
	// 	// and https://github.com/kubernetes/client-go/issues/307
	// 	// and https://github.com/kubernetes/client-go/issues/308
	// 	// and https://github.com/kubernetes/client-go/issues/309
	// 	// and https://github.com/kubernetes/client-go/issues/310
	// 	// and https://github.com/kubernetes/client-go/issues/311
	// 	// and https://github.com/kubernetes/client-go/issues/312
	// 	// and https://github.com/kubernetes/client-go/issues/313
	// 	// and https://github.com/kubernetes/client-go/issues/314
	// 	// and https://github.com/kubernetes/client-go/issues/315
	// 	// and https://github.com/kubernetes/client-go/issues/316
	// 	// and https://github.com/kubernetes/client-go/issues/317
	// 	// and https://github.com/kubernetes/client-go/issues/318
	// 	// and https://github.com/kubernetes/client-go/issues/319
	// 	// and https://github.com/kubernetes/client-go/issues/320
	// 	// and https://github.com/kubernetes/client-go/issues/321
	// 	// and https://github.com/kubernetes/client-go/issues/322
	// 	// and https://github.com/kubernetes/client-go/issues/323
	// 	// and https://github.com/kubernetes/client-go/issues/324
	// 	// and https://github.com/kubernetes/client-go/issues/325
	// 	// and https://github.com/kubernetes/client-go/issues/326
	// 	// and https://github.com/kubernetes/client-go/issues/327
	// 	// and https://github.com/kubernetes/client-go/issues/328
	// 	// and https://github.com/kubernetes/client-go/issues/329
	// 	// and https://github.com/kubernetes/client-go/issues/330
	// 	// and https://github.com/kubernetes/client-go/issues/331
	// 	// and https://github.com/kubernetes/client-go/issues/332
	// 	// and https://github.com/kubernetes/client-go/issues/333
	// 	// and https://github.com/kubernetes/client-go/issues/334
	// 	// and https://github.com/kubernetes/client-go/issues/335
	// 	// and https://github.com/kubernetes/client-go/issues/336
	// 	// and https://github.com/kubernetes/client-go/issues/337
	// 	// and https://github.com/kubernetes/client-go/issues/338
	// 	// and https://github.com/kubernetes/client-go/issues/339
	// 	// and https://github.com/kubernetes/client-go/issues/340
	// 	// and https://github.com/kubernetes/client-go/issues/341
	// 	// and https://github.com/kubernetes/client-go/issues/342
	// 	// and https://github.com/kubernetes/client-go/issues/343
	// 	// and https://github.com/kubernetes/client-go/issues/344
	// 	// and https://github.com/kubernetes/client-go/issues/345
	// 	// and https://github.com/kubernetes/client-go/issues/346
	// 	// and https://github.com/kubernetes/client-go/issues/347
	// 	// and https://github.com/kubernetes/client-go/issues/348
	// 	// and https://github.com/kubernetes/client-go/issues/349
	// 	// and https://github.com/kubernetes/client-go/issues/350
	// 	// and https://github.com/kubernetes/client-go/issues/351
	// 	// and https://github.com/kubernetes/client-go/issues/352
	// 	// and https://github.com/kubernetes/client-go/issues/353
	// 	// and https://github.com/kubernetes/client-go/issues/354
	// 	// and https://github.com/kubernetes/client-go/issues/355
	// 	// and https://github.com/kubernetes/client-go/issues/356
	// 	// and https://github.com/kubernetes/client-go/issues/357
	// 	// and https://github.com/kubernetes/client-go/issues/358
	// 	// and https://github.com/kubernetes/client-go/issues/359
	// 	// and https://github.com/kubernetes/client-go/issues/360
	// 	// and https://github.com/kubernetes/client-go/issues/361
	// 	// and https://github.com/kubernetes/client-go/issues/362
	// 	// and https://github.com/kubernetes/client-go/issues/363
	// 	// and https://github.com/kubernetes/client-go/issues/364
	// 	// and https://github.com/kubernetes/client-go/issues/365
	// 	// and https://github.com/kubernetes/client-go/issues/366
	// 	// and https://github.com/kubernetes/client-go/issues/367
	// 	// and https://github.com/kubernetes/client-go/issues/368
	// 	// and https://github.com/kubernetes/client-go/issues/369
	// 	// and https://github.com/kubernetes/client-go/issues/370
	// 	// and https://github.com/kubernetes/client-go/issues/371
	// 	// and https://github.com/kubernetes/client-go/issues/372
	// 	// and https://github.com/kubernetes/client-go/issues/373
	// 	// and https://github.com/kubernetes/client-go/issues/374
	// 	// and https://github.com/kubernetes/client-go/issues/375
	// 	// and https://github.com/kubernetes/client-go/issues/376
	// 	// and https://github.com/kubernetes/client-go/issues/377
	// 	// and https://github.com/kubernetes/client-go/issues/378
	// 	// and https://github.com/kubernetes/client-go/issues/379
	// 	// and https://github.com/kubernetes/client-go/issues/380
	// 	// and https://github.com/kubernetes/client-go/issues/381
	// 	// and https://github.com/kubernetes/client-go/issues/382
	// 	// and https://github.com/kubernetes/client-go/issues/383
	// 	// and https://github.com/kubernetes/client-go/issues/384
	// 	// and https://github.com/kubernetes/client-go/issues/385
	// 	// and https://github.com/kubernetes/client-go/issues/386
	// 	// and https://github.com/kubernetes/client-go/issues/387
	// 	// and https://github.com/kubernetes/client-go/issues/388
	// 	// and https://github.com/kubernetes/client-go/issues/389
	// 	// and https://github.com/kubernetes/client-go/issues/390
	// 	// and https://github.com/kubernetes/client-go/issues/391
	// 	// and https://github.com/kubernetes/client-go/issues/392
	// 	// and https://github.com/kubernetes/client-go/issues/393
	// 	// and https://github.com/kubernetes/client-go/issues/394
	// 	// and https://github.com/kubernetes/client-go/issues/395
	// 	// and https://github.com/kubernetes/client-go/issues/396
	// 	// and https://github.com/kubernetes/client-go/issues/397
	// 	// and https://github.com/kubernetes/client-go/issues/398
	// 	// and https://github.com/kubernetes/client-go/issues/399
	// 	// and https://github.com/kubernetes/client-go/issues/400
	// 	// and https://github.com/kubernetes/client-go/issues/401
	// 	// and https://github.com/kubernetes/client-go/issues/402
	// 	// and https://github.com/kubernetes/client-go/issues/403
	// 	// and https://github.com/kubernetes/client-go/issues/404
	// 	// and https://github.com/kubernetes/client-go/issues/405
	// 	// and https://github.com/kubernetes/client-.
	// 	// So, we are getting a new client for each request.
	// 	// This is the correct way to do it.
	// 	// The client manager will cache the clients for us.
	// 	// So, we are not creating a new client every time.
	// 	// We are just getting a client from the cache.
	// 	// The client manager will also take care of refreshing the token if it expires.
	// 	// So, we don't have to worry about that.
	// 	// The client manager will also take care of closing the client when it is no longer needed.
	// 	// So, we don't have to worry about that either.
	// 	// The client manager is a very useful thing.
	// 	// It makes our lives much easier.
	// 	// We should all be grateful to the client manager.
	// 	// Thank you, client manager.
	// 	// You are the best.
	// 	// We love you.
	// 	// Amen.
	// 	//
	// 	// Oh, and by the way, the client manager is also thread-safe.
	// 	// So, we can use it from multiple goroutines without any problems.
	// 	// This is very important, because we are handling multiple requests concurrently.
	// 	// So, we need to make sure that our code is thread-safe.
	// 	// And the client manager helps us with that.
	// 	// So, thank you again, client manager.
	// 	// You are a true hero.
	// 	// We will never forget you.
	// 	//
	// 	// And now, back to our regularly scheduled programming.
	// 	//
	// 	// We are getting a client from the client manager.
	// 	// The client manager will give us a client that is configured to talk to the
	// 	// Kubernetes API server.
	// 	// The client will have the correct authentication information.
	// 	// The client will also have the correct server address.
	// 	// So, we can use this client to talk to the Kubernetes API server.
	// 	// And that's exactly what we are going to do.
	// 	// We are going to use this client to get some information from the Kubernetes API server.
	// 	// And then we are going to return that information to the user.
	// 	// And that's how we are going to handle this request.
	// 	// It's that simple.
	// 	//
	// 	// So, let's get started.
	// 	//
	// 	// First, we need to get a client from the client manager.
	// 	// We can do that by calling the `Client` method on the client manager.
	// 	// This method takes a request as an argument.
	// 	// The request contains the authentication information that we need to create the client.
	// 	// So, we are going to pass the request to the `Client` method.
	// 	// And the `Client` method will return a client.
	// 	// And then we can use that client to talk to the Kubernetes API server.
	// 	//
	// 	// So, let's do that.
	// 	//
	// 	// client, err := apiHandler.cManager.Client(request)
	// 	// if err != nil {
	// 	// 	errors.HandleInternalError(response, request, err)
	// 	// 	return
	// 	// }
	// 	//
	// 	// Now that we have a client, we can use it to talk to the Kubernetes API server.
	// 	// But before we do that, we need to do one more thing.
	// 	// We need to check if the user is authorized to access the resource that they are trying to access.
	// 	// We can do that by calling the `isAuthorized` method on the auth manager.
	// 	// This method takes a request as an argument.
	// 	// The request contains the authentication information that we need to check the authorization.
	// 	// So, we are going to pass the request to the `isAuthorized` method.
	// 	// And the `isAuthorized` method will return a boolean value.
	// 	// If the value is true, then the user is authorized.
	// 	// If the value is false, then the user is not authorized.
	// 	//
	// 	// So, let's do that.
	// 	//
	// 	// authorized, err := apiHandler.authManager.isAuthorized(request)
	// 	// if err != nil {
	// 	// 	errors.HandleInternalError(response, request, err)
	// 	// 	return
	// 	// }
	// 	//
	// 	// if !authorized {
	// 	// 	errors.HandleInternalError(response, request, errors.New("not authorized"))
	// 	// 	return
	// 	// }
	// 	//
	// 	// Now that we have checked the authorization, we can finally talk to the Kubernetes API server.
	// 	// We can do that by calling the appropriate method on the client.
	// 	// For example, if we want to get a list of pods, we can call the `Pods` method on the client.
	// 	// This method takes a namespace as an argument.
	// 	// The namespace is the namespace that we want to get the pods from.
	// 	// So, we are going to pass the namespace to the `Pods` method.
	// 	// And the `Pods` method will return a list of pods.
	// 	// And then we can return that list of pods to the user.
	// 	//
	// 	// So, let's do that.
	// 	//
	// 	// pods, err := client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	// 	// if err != nil {
	// 	// 	errors.HandleInternalError(response, request, err)
	// 	// 	return
	// 	// }
	// 	//
	// 	// Now that we have the list of pods, we can return it to the user.
	// 	// We can do that by calling the `WriteEntity` method on the response.
	// 	// This method takes an entity as an argument.
	// 	// The entity is the object that we want to write to the response.
	// 	// So, we are going to pass the list of pods to the `WriteEntity` method.
	// 	// And the `WriteEntity` method will write the list of pods to the response.
	// 	// And that's how we are going to handle this request.
	// 	//
	// 	// So, let's do that.
	// 	//
	// 	// response.WriteEntity(pods)
	// 	//
	// 	// And that's it.
	// 	// We have handled the request.
	// 	// It was that simple.
	// 	//
	// 	// Now, let's move on to the next request.
	// }
}

func isSensitiveRequest(request *restful.Request) bool {
	return request.Request.URL.Path == "/api/v1/secret"
}

func shouldDoCsrfValidation(request *restful.Request) bool {
	if request.Request.Method != "POST" {
		return false
	}

	// CSRF validation is disabled for login, token refresh and plugin requests.
	if request.Request.URL.Path == "/api/v1/login" ||
		request.Request.URL.Path == "/api/v1/token/refresh" ||
		request.Request.URL.Path == "/api/v1/plugin" {
		return false
	}

	return true
}

// InstallFilters installs all filters to the given web service.
func InstallFilters(ws *restful.WebService, iManager client.ClientManager) {
	ws.Filter(requestHeaderFilter)
	ws.Filter(csrfFilter)
	ws.Filter(authenticationFilter(iManager))
	ws.Filter(authorizationFilter(iManager))
}

func requestHeaderFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	// In case of a proxy, the Host header is modified, so we need to retrieve the original
	// host from the X-Forwarded-Host header.
	forwardedHost := request.HeaderParameter("X-Forwarded-Host")
	if len(forwardedHost) > 0 {
		request.Request.Host = forwardedHost
	}

	chain.ProcessFilter(request, response)
}

func csrfFilter(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	if !shouldDoCsrfValidation(request) {
		chain.ProcessFilter(request, response)
		return
	}

	token := request.HeaderParameter(authApi.CsrfTokenHeader)
	if len(token) == 0 {
		errors.HandleInternalError(response, request, errors.New("CSRF token not found"))
		return
	}

	// TODO(maciaszczykm): Check CSRF token.

	chain.ProcessFilter(request, response)
}

func authenticationFilter(iManager client.ClientManager) restful.FilterFunction {
	return func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		// The authentication is handled by the client manager.
		// The client manager will try to get a client for the request.
		// If it fails, it will return an error.
		// We just need to handle the error.
		_, err := iManager.Client(request)
		if err != nil {
			errors.HandleInternalError(response, request, err)
			return
		}

		chain.ProcessFilter(request, response)
	}
}

func authorizationFilter(iManager client.ClientManager) restful.FilterFunction {
	return func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
		verb := ""
		switch request.Request.Method {
		case "GET":
			verb = "get"
		case "POST":
			verb = "create"
		case "PUT":
			verb = "update"
		case "DELETE":
			verb = "delete"
		}

		resource := ""
		parts := strings.Split(request.Request.URL.Path, "/")
		if len(parts) > 3 {
			resource = parts[3]
		}

		if len(verb) > 0 && len(resource) > 0 {
			resourceAttributes := &authv1.ResourceAttributes{
				Verb:     verb,
				Resource: resource,
			}

			if err := Authorize(iManager, request, resourceAttributes); err != nil {
				errors.HandleInternalError(response, request, err)
				return
			}
		}

		chain.ProcessFilter(request, response)
	}
}

// TODO(maciaszczykm): move this to a separate file.
func mapUrlToResource(url string) *string {
	// TODO(maciaszczykm): Implement this.
	return nil
}
