package plugin

import (
	"github.com/ParaServices/errgo"
	"github.com/ParaServices/kong/object"
	"github.com/ParaServices/paratils"
)

// NewJWTCredential ...
func NewJWTCredential(
	mutatorFuncs ...JWTCredentialMutatorFunc,
) (
	*JWTCredential,
	error,
) {
	jwtCred := &JWTCredential{}
	for i := range mutatorFuncs {
		if err := mutatorFuncs[i](jwtCred); err != nil {
			return nil, errgo.New(err)
		}

	}

	return jwtCred, nil
}

// JWTCredential ...
type JWTCredential struct {
	object.KongID
	object.Tags
	object.CreatedAt
	Consumer     *object.Consumer `json:"consumer,omitempty"`
	Key          string           `json:"key,omitempty"`
	Secret       string           `json:"secret,omitempty"`
	RSAPublicKey string           `json:"rsa_public_key,omitempty"`
	Algorithm    string           `json:"algorithm,omitempty"`
}

// GetConsumer ...
func (j *JWTCredential) GetConsumer() object.ConsumerAccessor {
	return j.Consumer
}

// HasConsumer ...
func (j *JWTCredential) HasConsumer() bool {
	return !paratils.IsNil(j.GetConsumer())
}

// HasConsumerID determines if the given JWTCredential has a consumer with a
// value to the ConsumerID. This determines the validaity of JWTCredential
// when creatinga consumer in the Kong API.
func (j *JWTCredential) HasConsumerID() bool {
	return j.HasConsumer() && j.GetConsumer().HasID()
}

// GetKey ...
func (j *JWTCredential) GetKey() string {
	return j.Key
}

// GetSecret ...
func (j *JWTCredential) GetSecret() string {
	return j.Secret
}

// GetRSAPublicKey ...
func (j *JWTCredential) GetRSAPublicKey() string {
	return j.RSAPublicKey
}

// GetAlgorithm ...
func (j *JWTCredential) GetAlgorithm() string {
	return j.Algorithm
}

// SetConsumer ...
func (j *JWTCredential) SetConsumer(getter object.ConsumerGetter) error {
	if paratils.IsNil(getter) {
		return nil
	}
	if paratils.IsNil(j.Consumer) {
		j.Consumer = &object.Consumer{}
	}

	return object.CopyConsumer(getter, j.Consumer)
}

// SetNewConsumer ...
func (j *JWTCredential) SetNewConsumer(
	newFn object.NewConsumerFunc,
	mutatorFuncs ...object.ConsumerMutatorFunc,
) error {
	if paratils.IsNil(newFn) || len(mutatorFuncs) == 0 {
		return nil
	}

	consumer, err := newFn(mutatorFuncs...)
	if err != nil {
		return errgo.New(err)
	}

	return j.SetConsumer(consumer)
}

// SetKey ...
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
	HasConsumer() bool
	HasConsumerID() bool
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
	SetNewConsumer(
		newFn object.NewConsumerFunc,
		mutatorFuncs ...object.ConsumerMutatorFunc,
	) error
	SetAlgorithm(algorithm string) error
	SetKey(key string) error
	SetRSAPublicKey(rsaPublicKey string) error
	SetSecret(secret string) error
}

type JWTCredentialAccessor interface {
	JWTCredentialGetter
	JWTCredentialSetter
}

// CopyJWTCredential ...
func CopyJWTCredential(
	getter JWTCredentialGetter,
	setter JWTCredentialSetter,
) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	if err := object.CopyKongID(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := object.CopyCreatedAt(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := object.CopyTags(getter, setter); err != nil {
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
