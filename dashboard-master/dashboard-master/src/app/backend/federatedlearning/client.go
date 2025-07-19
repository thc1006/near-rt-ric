
package federatedlearning

import "time"

// Client represents a registered xApp that can participate in training.
type Client struct {
    ID         string    `json:"id"`         // Kubernetes Pod Name/UID for uniqueness
    Status     string    `json:"status"`     // e.g., "available", "training", "offline"
    RegisteredAt time.Time `json:"registeredAt"`
    LastSeen   time.Time `json:"lastSeen"`
}
