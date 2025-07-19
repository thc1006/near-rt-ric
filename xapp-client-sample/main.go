
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	coordinatorURL = "http://localhost:9090/api/v1/fl"
	modelID        = "rrm-power-control"
)

// GlobalModel represents the master model managed by the coordinator.
type GlobalModel struct {
	ID          string    `json:"id"`
	Version     int       `json:"version"`
	Weights     []byte    `json:"-"`
	CreatedAt   time.Time `json:"createdAt"`
	Description string    `json:"description"`
}

// ModelUpdate is sent by an xApp client to the coordinator after a local training round.
type ModelUpdate struct {
	ClientID     string `json:"clientId"`
	ModelID      string `json:"modelId"`
	BaseVersion  int    `json:"baseVersion"`
	WeightUpdate []byte `json:"weightUpdate"`
}

func main() {
	// 1. Register with the coordinator
	clientID, err := register()
	if err != nil {
		fmt.Printf("Error registering with coordinator: %v\n", err)
		return
	}
	fmt.Printf("Registered with coordinator, client ID: %s\n", clientID)

	// 2. Get the global model
	model, err := getModel()
	if err != nil {
		fmt.Printf("Error getting model: %v\n", err)
		return
	}
	fmt.Printf("Got model version %d\n", model.Version)

	// 3. Train the model (conceptual)
	fmt.Println("Training model...")
	time.Sleep(5 * time.Second) // Simulate training
	newWeights := []byte("new weights")

	// 4. Submit the model update
	if err := submitUpdate(clientID, model.Version, newWeights); err != nil {
		fmt.Printf("Error submitting update: %v\n", err)
		return
	}
	fmt.Println("Submitted model update")
}

func register() (string, error) {
	reqBody, err := json.Marshal(map[string]string{"modelId": modelID})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(fmt.Sprintf("%s/register", coordinatorURL), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var client struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&client); err != nil {
		return "", err
	}

	return client.ID, nil
}

func getModel() (*GlobalModel, error) {
	resp, err := http.Get(fmt.Sprintf("%s/model/%s", coordinatorURL, modelID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	var model GlobalModel
	if err := json.NewDecoder(resp.Body).Decode(&model); err != nil {
		return nil, err
	}

	return &model, nil
}

func submitUpdate(clientID string, baseVersion int, newWeights []byte) error {
	update := ModelUpdate{
		ClientID:     clientID,
		ModelID:      modelID,
		BaseVersion:  baseVersion,
		WeightUpdate: newWeights,
	}

	reqBody, err := json.Marshal(update)
	if err != nil {
		return err
	}

	resp, err := http.Post(fmt.Sprintf("%s/model/%s/update", coordinatorURL, modelID), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	return nil
}
