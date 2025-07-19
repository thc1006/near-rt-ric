package federatedlearning

import (
	"encoding/json"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

// API provides the HTTP API for the federated learning coordinator.
type API struct {
	coordinator *Coordinator
}

// NewAPI creates a new API.
func NewAPI(coordinator *Coordinator) *API {
	return &API{coordinator: coordinator}
}

// RegisterRoutes registers the API routes.
func (a *API) RegisterRoutes(ws *restful.WebService) {
	ws.Route(ws.POST("/fl/register").To(a.registerClient))
	ws.Route(ws.GET("/fl/model/{modelId}").To(a.getModel))
	ws.Route(ws.POST("/fl/model/{modelId}/update").To(a.submitModelUpdate))
}

func (a *API) registerClient(req *restful.Request, resp *restful.Response) {
	var registrationReq struct {
		ModelID string `json:"modelId"`
	}
	if err := req.ReadEntity(&registrationReq); err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}

	client, err := a.coordinator.RegisterClient(req.Request.RemoteAddr)
	if err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}

	resp.WriteEntity(client)
}

func (a *API) getModel(req *restful.Request, resp *restful.Response) {
	modelID := req.PathParameter("modelId")

	model, err := a.coordinator.GetModel(modelID)
	if err != nil {
		resp.WriteError(http.StatusNotFound, err)
		return
	}

	resp.WriteEntity(model)
}

func (a *API) submitModelUpdate(req *restful.Request, resp *restful.Response) {
	var update ModelUpdate
	if err := req.ReadEntity(&update); err != nil {
		resp.WriteError(http.StatusBadRequest, err)
		return
	}

	if err := a.coordinator.SubmitModelUpdate(&update); err != nil {
		resp.WriteError(http.StatusInternalServerError, err)
		return
	}

	resp.WriteHeader(http.StatusAccepted)
}