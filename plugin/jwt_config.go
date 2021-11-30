package plugin

import (
	"encoding/json"

	"github.com/ParaServices/paratils"
)

// JWTConfigMutatorFunc ...
type JWTConfigMutatorFunc func(acessor JWTConfigAccessor) error

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

// MarshalToBytes ...
func (j JWTConfig) MarshalToBytes() ([]byte, error) {
	return json.Marshal(j)
}

func (j JWTConfig) GetURIParamNames() []string {
	return j.URIParamNames
}

func (j JWTConfig) GetCookieNames() []string {
	return j.CookieNames
}

func (j JWTConfig) GetHeaderNames() []string {
	return j.HeaderNames
}

func (j JWTConfig) GetClaimsToVerify() []string {
	return j.ClaimsToVerify
}

func (j JWTConfig) GetKeyClaimName() string {
	return j.KeyClaimName
}

func (j JWTConfig) GetSecretIsBase64() bool {
	return j.SecretIsBase64
}

func (j JWTConfig) GetRunOnPreFlight() bool {
	return j.RunOnPreFlight
}

func (j JWTConfig) GetMaximumExpiration() int {
	return j.MaximumExpiration
}

func (j *JWTConfig) SetURIParamNames(paramNames ...string) error {
	if len(paramNames) < 1 {
		return nil
	}
	j.URIParamNames = make([]string, len(paramNames))
	copy(j.URIParamNames, paramNames)
	return nil
}

// AddURIParamNames ...
func (j *JWTConfig) AddURIParamNames(paramNames ...string) error {
	if len(paramNames) < 1 {
		return nil
	}
	if paratils.IsNil(j.URIParamNames) || len(j.URIParamNames) == 0 {
		j.URIParamNames = make([]string, len(paramNames))
		copy(j.URIParamNames, paramNames)
	} else {
		j.URIParamNames = append(j.URIParamNames, paramNames...)
	}
	return nil
}

func (j *JWTConfig) SetCookieNames(cookieNames ...string) error {
	if len(cookieNames) < 1 {
		return nil
	}
	j.CookieNames = make([]string, len(cookieNames))
	copy(j.CookieNames, cookieNames)
	return nil
}

// AddCookieNames ...
func (j *JWTConfig) AddCookieNames(cookieNames ...string) error {
	if len(cookieNames) < 1 {
		return nil
	}
	if paratils.IsNil(j.CookieNames) || len(j.CookieNames) == 0 {
		j.CookieNames = make([]string, len(cookieNames))
		copy(j.CookieNames, cookieNames)
	} else {
		j.CookieNames = append(j.CookieNames, cookieNames...)
	}
	return nil
}

func (j *JWTConfig) SetHeaderNames(headerNames ...string) error {
	if len(headerNames) < 1 {
		return nil
	}
	j.HeaderNames = make([]string, len(headerNames))
	copy(j.HeaderNames, headerNames)
	return nil
}

// AddHeaderNames ...
func (j *JWTConfig) AddHeaderNames(headerNames ...string) error {
	if len(headerNames) < 1 {
		return nil
	}
	if paratils.IsNil(j.HeaderNames) || len(j.HeaderNames) == 0 {
		j.HeaderNames = make([]string, len(headerNames))
		copy(j.HeaderNames, headerNames)
	} else {
		j.HeaderNames = append(j.HeaderNames, headerNames...)
	}
	return nil
}

// SetClaimsToVerify ...
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

// SetKeyClaimName ...
func (j *JWTConfig) SetKeyClaimName(claimName string) error {
	j.KeyClaimName = claimName
	return nil
}

// SetSecretIsBase64 ...
func (j *JWTConfig) SetSecretIsBase64(isBase64 bool) error {
	j.SecretIsBase64 = isBase64
	return nil
}

// SetRunOnPreFlight ...
func (j *JWTConfig) SetRunOnPreFlight(run bool) error {
	j.RunOnPreFlight = run
	return nil
}

// SetMaximumExpiration ...
func (j *JWTConfig) SetMaximumExpiration(maxExp int) error {
	j.MaximumExpiration = maxExp
	return nil
}

var _ JWTConfigAccessor = (*JWTConfig)(nil)

type JWTConfigGetter interface {
	GetClaimsToVerify() []string
	GetCookieNames() []string
	GetHeaderNames() []string
	GetKeyClaimName() string
	GetMaximumExpiration() int
	GetRunOnPreFlight() bool
	GetSecretIsBase64() bool
	GetURIParamNames() []string
	MarshalToBytes() ([]byte, error)
}

type JWTConfigSetter interface {
	SetClaimsToVerify(claims ...string) error
	SetCookieNames(cookieNames ...string) error
	AddCookieNames(cookieNames ...string) error
	SetHeaderNames(headers ...string) error
	AddHeaderNames(headers ...string) error
	SetKeyClaimName(claimName string) error
	SetMaximumExpiration(maxExp int) error
	SetRunOnPreFlight(run bool) error
	SetSecretIsBase64(isBase64 bool) error
	SetURIParamNames(paramNames ...string) error
	AddURIParamNames(paramNames ...string) error
}

type JWTConfigAccessor interface {
	JWTConfigGetter
	JWTConfigSetter
}
