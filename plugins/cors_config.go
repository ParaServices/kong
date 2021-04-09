package plugins

import "encoding/json"

// COCORSConfig ...
type CORSConfig struct {
	Origins           []string `json:"origins,omitempty"`
	Methods           []string `json:"methods,omitempty"`
	Headers           []string `json:"headers,omitempty"`
	Credentials       *bool    `json:"credentials,omitempty"`
	ExposedHeaders    []string `json:"exposed_headers,omitempty"`
	MaxAge            int      `json:"max_age,omitempty"`
	PreFlightContinue *bool    `json:"preflight_continue,omitempty"`
}

// Marshal ...
func (c *CORSConfig) Marshal() ([]byte, error) {
	return json.Marshal(c)
}
