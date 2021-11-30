package object

import (
	"github.com/ParaServices/errgo"
	"github.com/ParaServices/paratils"
)

// ConsumerMutatorFunc ...
type ConsumerMutatorFunc func(accessor ConsumerAccessor) error

// NewConsumerFunc defines the function type that initializes a new consmer
type NewConsumerFunc func(mutatorFuncs ...ConsumerMutatorFunc) (*Consumer, error)

// NewConsumer ...
func NewConsumer(mutatorFuncs ...ConsumerMutatorFunc) (*Consumer, error) {
	c := &Consumer{}
	for i := range mutatorFuncs {
		if err := mutatorFuncs[i](c); err != nil {
			return nil, errgo.New(err)
		}
	}
	return c, nil
}

// Consumer ...
type Consumer struct {
	KongID
	CreatedAt
	Tags
	Username string `json:"username,omitempty"`
	CustomID string `json:"custom_id,omitempty"`
}

// GetUsername ...
func (c Consumer) GetUsername() string {
	return c.Username
}

// GetCustomID ...
func (c Consumer) GetCustomID() string {
	return c.CustomID
}

func (c *Consumer) SetUsername(username string) error {
	c.Username = username
	return nil
}

func (c *Consumer) SetCustomID(customID string) error {
	c.CustomID = customID
	return nil
}

type ConsumerGetter interface {
	KongIDGetter
	CreatedAtGetter
	TagsGetter
	GetCustomID() string
	GetUsername() string
}

type ConsumerSetter interface {
	KongIDSetter
	CreatedAtSetter
	TagsSetter
	SetCustomID(customID string) error
	SetUsername(username string) error
}

type ConsumerAccessor interface {
	ConsumerGetter
	ConsumerSetter
}

var _ ConsumerAccessor = (*Consumer)(nil)

// CopyConsumer ...
func CopyConsumer(getter ConsumerGetter, setter ConsumerSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	if err := CopyKongID(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := CopyCreatedAt(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := CopyTags(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetUsername(getter.GetUsername()); err != nil {
		return errgo.New(err)
	}

	return setter.SetCustomID(getter.GetCustomID())
}
