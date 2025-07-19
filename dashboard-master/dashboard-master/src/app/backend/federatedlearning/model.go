
package federatedlearning

import "time"

// GlobalModel represents the master model managed by the coordinator.
// The weights are stored as a byte slice to remain agnostic to the specific
// ML framework (e.g., TensorFlow, PyTorch) used by the xApps.
type GlobalModel struct {
    ID          string    `json:"id"`           // Unique identifier for the model, e.g., "rrm-power-control"
    Version     int       `json:"version"`      // Monotonically increasing version number
    Weights     []byte    `json:"-"`            // The actual model weights (omitted from standard JSON responses)
    CreatedAt   time.Time `json:"createdAt"`
    Description string    `json:"description"`
}

// ModelUpdate is sent by an xApp client to the coordinator after a local training round.
type ModelUpdate struct {
    ClientID     string `json:"clientId"`     // The unique ID of the xApp client (e.g., pod name)
    ModelID      string `json:"modelId"`      // The ID of the model being updated
    BaseVersion  int    `json:"baseVersion"`  // The version of the global model the update is based on
    WeightUpdate []byte `json:"weightUpdate"` // The new weights or gradients from the client
}
