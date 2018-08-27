package kong

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type StatusResponse struct {
	Server struct {
		TotalRequests       int `json:"total_requests"`
		ConnectionsActive   int `json:"connections_active"`
		ConnectionsAccepted int `json:"connections_accepted"`
		ConnectionsHandled  int `json:"connections_handled"`
		ConnectionsReading  int `json:"connections_reading"`
		ConnectionsWriting  int `json:"connections_writing"`
		ConnectionsWaiting  int `json:"connections_waiting"`
	} `json:"server"`
	Database struct {
		Reachable bool `json:"reachable"`
	} `json:"database"`
}

func (c *client) GetStatus() (*StatusResponse, error) {
	// Build URL
	relURL, err := url.Parse(path.Join("status"))
	if err != nil {
		return nil, err
	}
	url := c.BaseURL.ResolveReference(relURL)
	// Create Request
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	// Send Request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Register Response Body Decoder
	decoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Kong Status Check Failed, status: %s, url: %s", resp.Status, url.String())
	}
	// Decode Status Response
	var statusResponse StatusResponse
	if err := decoder.Decode(&statusResponse); err != nil {
		return nil, err
	}
	return &statusResponse, nil
}
