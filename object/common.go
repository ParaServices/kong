package object

import (
	"github.com/ParaServices/paratils"
)

// NewKongID ...
func NewKongID(id string) *KongID {
	return &KongID{
		ID: id,
	}
}

// KongID ...
type KongID struct {
	ID string `json:"id,omitempty"`
}

func (k KongID) GetID() string {
	return k.ID
}

func (k KongID) HasID() bool {
	return k.GetID() != ""
}

func (k *KongID) SetID(id string) error {
	k.ID = id
	return nil
}

var _ KongIDAccessor = (*KongID)(nil)

type KongIDGetter interface {
	GetID() string
	HasID() bool
}

type KongIDSetter interface {
	SetID(id string) error
}

type KongIDAccessor interface {
	KongIDGetter
	KongIDSetter
}

// CopyKongID ...
func CopyKongID(getter KongIDGetter, setter KongIDSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetID(getter.GetID())
}

// Tags ...
type Tags struct {
	Tags []string `json:"tags,omitempty"`
}

// GetTags ...
func (t Tags) GetTags() []string {
	return t.Tags
}

// SetTags ...
func (t *Tags) SetTags(tags ...string) error {
	if len(tags) < 1 {
		return nil
	}
	t.Tags = make([]string, len(tags))

	copy(tags, t.Tags)
	return nil
}

// AddTags ...
func (t *Tags) AddTags(tags ...string) error {
	if len(tags) < 1 {
		return nil
	}
	if paratils.IsNil(t.Tags) || len(t.Tags) == 0 {
		t.Tags = make([]string, len(tags))
		copy(tags, t.Tags)
	} else {
		t.Tags = append(t.Tags, tags...)
	}
	return nil
}

var _ TagsAccessor = (*Tags)(nil)

type TagsGetter interface {
	GetTags() []string
}

type TagsSetter interface {
	SetTags(tags ...string) error
	AddTags(tags ...string) error
}

type TagsAccessor interface {
	TagsGetter
	TagsSetter
}

// CopyTags ...
func CopyTags(getter TagsGetter, setter TagsSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetTags(getter.GetTags()...)
}

// CreatedAt ...
type CreatedAt struct {
	CreatedAt int64 `json:"created_at,omitempty"`
}

func (c CreatedAt) GetCreatedAt() int64 {
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

// CopyCreatedAt ...
func CopyCreatedAt(getter CreatedAtGetter, setter CreatedAtSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetCreatedAt(getter.GetCreatedAt())
}

// UpdatedAt ...
type UpdatedAt struct {
	UpdatedAt int64 `json:"updated_at,omitempty"`
}

func (u UpdatedAt) GetUpdatedAt() int64 {
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

// CopyUpdatedAt ...
func CopyUpdatedAt(getter UpdatedAtGetter, setter UpdatedAtSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetUpdatedAt(getter.GetUpdatedAt())
}

// Name ...
type Name struct {
	Name string `json:"name,omtiempty"`
}

func (n Name) GetName() string {
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

// CopyName ...
func CopyName(getter NameGetter, setter NameSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	return setter.SetName(getter.GetName())
}
