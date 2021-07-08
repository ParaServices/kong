package plugin

import "encoding/json"

// JWTConfig ...
type JWTConfig struct {
	URIParamNames     string `json:"uri_param_names,omitempty"`
	CookieNames       string `json:"cookie_names,omitempty"`
	HeaderNames       string `json:"header_names,omitempty"`
	ClaimsToVerify    string `json:"claims_to_verify,omitempty"`
	KeyClaimName      string `json:"key_claim_name,omitempty"`
	SecretIsBase64    *bool  `json:"secret_is_base64,omitempty"`
	RunOnPreFlight    *bool  `json:"run_on_preflight,omitempty"`
	MaximumExpiration int    `json:"maximum_expiration,omitempty"`
}

// Marshal ...
func (j *JWTConfig) Marshal() ([]byte, error) {
	return json.Marshal(j)
}
