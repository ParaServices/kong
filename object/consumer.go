package object

import (
	"github.com/ParaServices/errgo"
	"github.com/ParaServices/paratils"
)

func NewConsumer(username, customID string, tags ...string) *Consumer {
	return &Consumer{
		Username: username,
		CustomID: customID,
		Tags: Tags{
			Tags: tags,
		},
	}
}

// Consumer ...
type Consumer struct {
	KongID
	CreatedAt
	Tags
	Username string `json:"username,omitempty"`
	CustomID string `json:"custom_id,omitempty"`
}

func (c *Consumer) GetUsername() string {
	return c.Username
}

func (c *Consumer) GetCustomID() string {
	return c.CustomID
}

func (c *Consumer) SetUsername(username string) error {
	c.Username = username
	return nil
}

func (c *Consumer) SetCustomID(customid string) error {
	c.CustomID = customid
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
	SetCustomID(customid string) error
	SetUsername(username string) error
}

type ConsumerAccessor interface {
	ConsumerGetter
	ConsumerSetter
}

var _ ConsumerAccessor = (*Consumer)(nil)

func MarshalConsumer(getter ConsumerGetter, setter ConsumerSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	if err := MarshalKongID(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := MarshalCreatedAt(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := MarshalTags(getter, setter); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetUsername(getter.GetUsername()); err != nil {
		return errgo.New(err)
	}
	return setter.SetCustomID(getter.GetCustomID())
}
