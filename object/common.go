package object

import (
	"github.com/ParaServices/paratils"
)

func NewKongID(id string) *KongID {
	return &KongID{
		ID: id,
	}
}

type KongID struct {
	ID string `json:"id,omitempty"`
}

func (k *KongID) GetID() string {
	return k.ID
}

func (k *KongID) SetID(id string) error {
	k.ID = id
	return nil
}

var _ KongIDAccessor = (*KongID)(nil)

type KongIDGetter interface {
	GetID() string
}

type KongIDSetter interface {
	SetID(id string) error
}

type KongIDAccessor interface {
	KongIDGetter
	KongIDSetter
}

func MarshalKongID(getter KongIDGetter, setter KongIDSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetID(getter.GetID())
}

type Tags struct {
	Tags []string `json:"tags,omitempty"`
}

func (t *Tags) GetTags() []string {
	return t.Tags
}

func (t *Tags) SetTags(tags ...string) error {
	if paratils.IsNil(t.Tags) {
		t.Tags = make([]string, 0)
	}
	if len(tags) < 1 {
		return nil
	}

	copy(tags, t.Tags)
	return nil
}

var _ TagsAccessor = (*Tags)(nil)

type TagsGetter interface {
	GetTags() []string
}

type TagsSetter interface {
	SetTags(tags ...string) error
}

type TagsAccessor interface {
	TagsGetter
	TagsSetter
}

func MarshalTags(getter TagsGetter, setter TagsSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetTags(getter.GetTags()...)
}

type CreatedAt struct {
	CreatedAt int64 `json:"created_at,omitempty"`
}

func (c *CreatedAt) GetCreatedAt() int64 {
	return c.CreatedAt
}

func (c *CreatedAt) SetCreatedAt(createdat int64) error {
	c.CreatedAt = createdat
	return nil
}

type CreatedAtGetter interface {
	GetCreatedAt() int64
}

type CreatedAtSetter interface {
	SetCreatedAt(createdAt int64) error
}

type CreatedAtAccessor interface {
	CreatedAtGetter
	CreatedAtSetter
}

func MarshalCreatedAt(getter CreatedAtGetter, setter CreatedAtSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetCreatedAt(getter.GetCreatedAt())
}

type UpdatedAt struct {
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

func (u *UpdatedAt) GetUpdatedAt() int64 {
	return u.UpdatedAt
}

func (u *UpdatedAt) SetUpdatedAt(updatedat int64) error {
	u.UpdatedAt = updatedat
	return nil
}

type UpdatedAtGetter interface {
	GetUpdatedAt() int64
}

type UpdatedAtSetter interface {
	SetUpdatedAt(updatedAt int64) error
}

type UpdatedAtAccessor interface {
	UpdatedAtGetter
	UpdatedAtSetter
}

func MarshalUpdatedAt(getter UpdatedAtGetter, setter UpdatedAtSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetUpdatedAt(getter.GetUpdatedAt())
}

type Name struct {
	Name string `json:"name,omtiempty"`
}

func (n *Name) GetName() string {
	return n.Name
}

func (n *Name) SetName(name string) error {
	n.Name = name
	return nil
}

var _ NameAccessor = (*Name)(nil)

type NameGetter interface {
	GetName() string
}

type NameSetter interface {
	SetName(name string) error
}

type NameAccessor interface {
	NameGetter
	NameSetter
}

func MarshalName(getter NameGetter, setter NameSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetName(getter.GetName())
}
