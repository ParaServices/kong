package object

import (
	"encoding/json"

	"github.com/ParaServices/errgo"
	"github.com/ParaServices/paratils"
	"github.com/ParaServices/paratils/typeconv/byteconv"
)

// PluginConfig is an interface that needs to be implemented if configuration
// for the given plugin is needed.
type PluginConfig interface {
	Marshal() ([]byte, error)
}

// Plugin ...
type Plugin struct {
	Name
	Service  *KongID          `json:"service,omitempty"`
	Route    *KongID          `json:"route,omitempty"`
	Consumer *KongID          `json:"consumer,omitempty"`
	Enabled  *bool            `json:"enabled,omitempty"`
	Config   *json.RawMessage `json:"config,omitempty"`
}

func (p *Plugin) GetService() KongIDAccessor {
	return p.Service
}

func (p *Plugin) GetRoute() KongIDAccessor {
	return p.Route
}

func (p *Plugin) GetConsumer() KongIDAccessor {
	return p.Consumer
}

func (p *Plugin) GetEnabled() *bool {
	return p.Enabled
}

func (p *Plugin) IsEnabled() bool {
	return *p.Enabled
}

func (p *Plugin) GetConfig() *json.RawMessage {
	return p.Config
}

func (p *Plugin) SetService(getter KongIDGetter) error {
	if paratils.IsNil(getter) {
		return nil
	}
	if paratils.IsNil(p.Service) {
		p.Service = &KongID{}
	}

	return MarshalKongID(getter, p.Service)
}

func (p *Plugin) SetRoute(getter KongIDGetter) error {
	if paratils.IsNil(getter) {
		return nil
	}
	if paratils.IsNil(p.Route) {
		p.Route = &KongID{}
	}

	return MarshalKongID(getter, p.Route)
}

func (p *Plugin) SetConsumer(getter KongIDGetter) error {
	if paratils.IsNil(getter) {
		return nil
	}
	if paratils.IsNil(p.Consumer) {
		p.Consumer = &KongID{}
	}

	return MarshalKongID(getter, p.Consumer)
}

func (p *Plugin) SetEnabled(enabled *bool) error {
	p.Enabled = enabled
	return nil
}

func (p *Plugin) SetConfig(config *json.RawMessage) error {
	p.Config = config
	return nil
}

func (p *Plugin) MarshalConfig(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return errgo.New(err)
	}
	return p.SetConfig(byteconv.ToJSONRawMessagePtr(b))
}

var _ PluginAccessor = (*Plugin)(nil)

type PluginGetter interface {
	NameGetter
	GetConfig() *json.RawMessage
	GetConsumer() KongIDAccessor
	GetEnabled() *bool
	GetRoute() KongIDAccessor
	GetService() KongIDAccessor
	IsEnabled() bool
}

type PluginSetter interface {
	NameSetter
	SetConfig(config *json.RawMessage) error
	MarshalConfig(v interface{}) error
	SetConsumer(getter KongIDGetter) error
	SetEnabled(enabled *bool) error
	SetRoute(getter KongIDGetter) error
	SetService(getter KongIDGetter) error
}

type PluginAccessor interface {
	PluginGetter
	PluginSetter
}

func MarshalPlugin(getter PluginGetter, setter PluginSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	if err := setter.SetName(getter.GetName()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetConfig(getter.GetConfig()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetConsumer(getter.GetConsumer()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetEnabled(getter.GetEnabled()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetRoute(getter.GetRoute()); err != nil {
		return errgo.New(err)
	}
	return setter.SetService(getter.GetService())
}

// Plugins ...
type Plugins []Plugin

func (p Plugins) GetLength() int {
	return len(p)
}

// HasPlugin ...
func (p Plugins) HasPlugin(plugin string) bool {
	for i := range p {
		if plugin == p[i].GetName() {
			return true
		}
	}
	return false
}
