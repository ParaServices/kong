package plugin

import (
	"github.com/ParaServices/errgo"
	"github.com/ParaServices/kong/object"
	"github.com/ParaServices/paratils"
)

type JWTCredentialSetterFunc func(setter JWTCredentialSetter) error

func SetJWTCredentialKey(key string) JWTCredentialSetterFunc {
	return func(setter JWTCredentialSetter) error {
		return setter.SetKey(key)
	}
}

func SetJWTCredentialSecret(secret string) JWTCredentialSetterFunc {
	return func(setter JWTCredentialSetter) error {
		return setter.SetSecret(secret)
	}
}

func SetJWTCredentialRSAPublicKey(rsapublickey string) JWTCredentialSetterFunc {
	return func(setter JWTCredentialSetter) error {
		return setter.SetRSAPublicKey(rsapublickey)
	}
}

func SetJWTCredentialAlgorithm(algorithm string) JWTCredentialSetterFunc {
	return func(setter JWTCredentialSetter) error {
		return setter.SetAlgorithm(algorithm)
	}
}

func NewJWTCredential(getter object.ConsumerGetter, setterFuncs ...JWTCredentialSetterFunc) (*JWTCredential, error) {
	if paratils.IsNil(getter) {
		return nil, errgo.NewF("consumer is nil")
	}

	if paratils.StringIsEmpty(getter.GetCustomID()) {
		return nil, errgo.NewF("custom ID is empty")
	}
	if paratils.StringIsEmpty(getter.GetUsername()) {
		return nil, errgo.NewF("username is empty")
	}

	consumer := object.Consumer{}
	if err := object.MarshalConsumer(getter, &consumer); err != nil {
		return nil, errgo.New(err)
	}

	jwtCred := &JWTCredential{
		Consumer: &consumer,
	}
	for i := range setterFuncs {
		if err := setterFuncs[i](jwtCred); err != nil {
			return nil, errgo.New(err)
		}

	}

	return jwtCred, nil
}

type JWTCredential struct {
	object.KongID
	object.Tags
	object.CreatedAt
	Consumer     *object.Consumer `json:"consumer,omitempty"`
	Key          string           `json:"key,omitempty"`
	Secret       string           `json:"secret,omptempty"`
	RSAPublicKey string           `json:"rsa_public_key,omitempty"`
	Algorithm    string           `json:"algorithm,omitempty"`
}

func (j *JWTCredential) GetConsumer() object.ConsumerAccessor {
	return j.Consumer
}

func (j *JWTCredential) GetKey() string {
	return j.Key
}

func (j *JWTCredential) GetSecret() string {
	return j.Secret
}

func (j *JWTCredential) GetRSAPublicKey() string {
	return j.RSAPublicKey
}

func (j *JWTCredential) GetAlgorithm() string {
	return j.Algorithm
}

func (j *JWTCredential) SetConsumer(getter object.ConsumerGetter) error {
	if paratils.IsNil(getter) {
		return nil
	}
	if paratils.IsNil(j.Consumer) {
		j.Consumer = &object.Consumer{}
	}

	return object.MarshalConsumer(getter, j.Consumer)
}

func (j *JWTCredential) SetKey(key string) error {
	j.Key = key
	return nil
}

func (j *JWTCredential) SetSecret(secret string) error {
	j.Secret = secret
	return nil
}

func (j *JWTCredential) SetRSAPublicKey(rsaPublicKey string) error {
	j.RSAPublicKey = rsaPublicKey
	return nil
}
func (j *JWTCredential) SetAlgorithm(algorithm string) error {
	j.Algorithm = algorithm
	return nil
}

var _ JWTCredentialAccessor = (*JWTCredential)(nil)

type JWTCredentialGetter interface {
	object.KongIDGetter
	object.TagsGetter
	object.CreatedAtGetter
	GetConsumer() object.ConsumerAccessor
	GetAlgorithm() string
	GetKey() string
	GetRSAPublicKey() string
	GetSecret() string
}

type JWTCredentialSetter interface {
	object.KongIDSetter
	object.TagsSetter
	object.CreatedAtSetter
	SetConsumer(getter object.ConsumerGetter) error
	SetAlgorithm(algorithm string) error
	SetKey(key string) error
	SetRSAPublicKey(rsaPublicKey string) error
	SetSecret(secret string) error
}

type JWTCredentialAccessor interface {
	JWTCredentialGetter
	JWTCredentialSetter
}

func MarshalJWTCredential(getter JWTCredentialGetter, setter JWTCredentialSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	if err := object.MarshalKongID(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := object.MarshalCreatedAt(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := object.MarshalTags(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetConsumer(getter.GetConsumer()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetAlgorithm(getter.GetAlgorithm()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetKey(getter.GetKey()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetRSAPublicKey(getter.GetRSAPublicKey()); err != nil {
		return errgo.New(err)
	}
	return setter.SetSecret(getter.GetSecret())
}
