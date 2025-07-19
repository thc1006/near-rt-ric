
package federatedlearning

import (
	"fmt"
	"sync"
)

// Coordinator manages the federated learning process.
type Coordinator struct {
	mu           sync.Mutex
	models       map[string]*GlobalModel
	clients      map[string]*Client
	modelUpdates map[string][]*ModelUpdate
}

// NewCoordinator creates a new Coordinator.
func NewCoordinator() *Coordinator {
	return &Coordinator{
		models:       make(map[string]*GlobalModel),
		clients:      make(map[string]*Client),
		modelUpdates: make(map[string][]*ModelUpdate),
	}
}

// RegisterClient registers a new client.
func (c *Coordinator) RegisterClient(id string) (*Client, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.clients[id]; exists {
		return nil, fmt.Errorf("client already registered: %s", id)
	}

	client := &Client{ID: id, Status: "available"}
	c.clients[id] = client
	return client, nil
}

// GetModel returns the latest global model.
func (c *Coordinator) GetModel(id string) (*GlobalModel, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	model, exists := c.models[id]
	if !exists {
		return nil, fmt.Errorf("model not found: %s", id)
	}

	return model, nil
}

// SubmitModelUpdate submits a model update from a client.
func (c *Coordinator) SubmitModelUpdate(update *ModelUpdate) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.clients[update.ClientID]; !exists {
		return fmt.Errorf("client not registered: %s", update.ClientID)
	}

	c.modelUpdates[update.ModelID] = append(c.modelUpdates[update.ModelID], update)
	return nil
}
