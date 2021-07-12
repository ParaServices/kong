package plugin

import (
	"encoding/json"

	"github.com/ParaServices/paratils"
)

// JWTConfig ...
type JWTConfig struct {
	URIParamNames     []string `json:"uri_param_names,omitempty"`
	CookieNames       []string `json:"cookie_names,omitempty"`
	HeaderNames       []string `json:"header_names,omitempty"`
	ClaimsToVerify    []string `json:"claims_to_verify,omitempty"`
	KeyClaimName      string   `json:"key_claim_name,omitempty"`
	SecretIsBase64    bool     `json:"secret_is_base64"`
	RunOnPreFlight    bool     `json:"run_on_preflight"`
	MaximumExpiration int      `json:"maximum_expiration,omitempty"`
}

// Marshal ...
func (j *JWTConfig) Marshal() ([]byte, error) {
	return json.Marshal(j)
}

func (j *JWTConfig) GetURIParamNames() []string {
	return j.URIParamNames
}

func (j *JWTConfig) GetCookieNames() []string {
	return j.CookieNames
}

func (j *JWTConfig) GetHeaderNames() []string {
	return j.HeaderNames
}

func (j *JWTConfig) GetClaimsToVerify() []string {
	return j.ClaimsToVerify
}

func (j *JWTConfig) GetKeyClaimName() string {
	return j.KeyClaimName
}

func (j *JWTConfig) GetSecretIsBase64() bool {
	return j.SecretIsBase64
}

func (j *JWTConfig) GetRunOnPreFlight() bool {
	return j.RunOnPreFlight
}

func (j *JWTConfig) GetMaximumExpiration() int {
	return j.MaximumExpiration
}

func (j *JWTConfig) SetURIParamNames(paramNames ...string) error {
	if len(paramNames) < 1 {
		return nil
	}
	if paratils.IsNil(j.URIParamNames) {
		j.URIParamNames = make([]string, 0)
	}
	copy(j.URIParamNames, paramNames)
	return nil
}

func (j *JWTConfig) SetCookieNames(cookieNames ...string) error {
	if len(cookieNames) < 1 {
		return nil
	}
	if paratils.IsNil(j.CookieNames) {
		j.CookieNames = make([]string, 0)
	}
	copy(j.CookieNames, cookieNames)
	return nil
}

func (j *JWTConfig) SetHeaderNames(headers ...string) error {
	if len(headers) < 1 {
		return nil
	}
	if paratils.IsNil(j.HeaderNames) {
		j.HeaderNames = make([]string, 0)
	}
	copy(j.HeaderNames, headers)
	return nil
}

func (j *JWTConfig) SetClaimsToVerify(claims ...string) error {
	if len(claims) < 1 {
		return nil
	}
	if paratils.IsNil(j.ClaimsToVerify) {
		j.ClaimsToVerify = make([]string, 0)
	}
	copy(j.ClaimsToVerify, claims)
	return nil
}

func (j *JWTConfig) SetKeyClaimName(claimName string) error {
	j.KeyClaimName = claimName
	return nil
}

func (j *JWTConfig) SetSecretIsBase64(isBase64 bool) error {
	j.SecretIsBase64 = isBase64
	return nil
}

func (j *JWTConfig) SetRunOnPreFlight(run bool) error {
	j.RunOnPreFlight = run
	return nil
}

func (j *JWTConfig) SetMaximumExpiration(maxExp int) error {
	j.MaximumExpiration = maxExp
	return nil
}

type JWTConfigGetter interface {
	GetClaimsToVerify() []string
	GetCookieNames() []string
	GetHeaderNames() []string
	GetKeyClaimName() string
	GetMaximumExpiration() int
	GetRunOnPreFlight() bool
	GetSecretIsBase64() bool
	GetURIParamNames() []string
	Marshal() ([]byte, error)
}

type JWTConfigSetter interface {
	SetClaimsToVerify(claims ...string) error
	SetCookieNames(cookieNames ...string) error
	SetHeaderNames(headers ...string) error
	SetKeyClaimName(claimName string) error
	SetMaximumExpiration(maxExp int) error
	SetRunOnPreFlight(run bool) error
	SetSecretIsBase64(isBase64 bool) error
	SetURIParamNames(paramNames ...string) error
}

type JWTConfigAccessor interface {
	JWTConfigGetter
	JWTConfigSetter
}
