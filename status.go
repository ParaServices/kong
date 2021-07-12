package kong

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/ParaServices/errgo"
)

// StatusResponse ...
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

func (c *Client) GetStatus() (*StatusResponse, error) {
	// Build URL
	relURL, err := url.Parse(path.Join("status"))
	if err != nil {
		return nil, errgo.New(err)
	}
	url := c.baseURL.ResolveReference(relURL)
	// Create Request
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, errgo.New(err)
	}
	req.Header.Add("Content-Type", "application/json")

	// Send Request
	resp, err := c.doRequest(req)
	if err != nil {
		return nil, errgo.New(err)
	}
	defer resp.Body.Close()
	// Register Response Body Decoder
	decoder := json.NewDecoder(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("kong status check failed, status: %s, url: %s", resp.Status, url.String())
	}
	// Decode Status Response
	var statusResponse StatusResponse
	if err := decoder.Decode(&statusResponse); err != nil {
		return nil, errgo.New(err)
	}
	return &statusResponse, nil
}
