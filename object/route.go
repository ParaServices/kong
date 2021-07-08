package object

import (
	"github.com/ParaServices/errgo"
	"github.com/ParaServices/paratils"
)

// Route ...
type Route struct {
	KongID
	CreatedAt
	UpdatedAt
	Service       *KongID  `json:"service,omitempty"`
	Protocols     []string `json:"protocols,omitempty"`
	Paths         []string `json:"paths,omitempty"`
	Methods       []string `json:"methods,omitempty"`
	Hosts         []string `json:"hosts,omitempty"`
	PreserveHost  bool     `json:"preserve_host,omitempty"`
	StripPath     bool     `json:"strip_path,omtiempty"`
	RegexPriority int      `json:"regex_priority,omitempty"`
}

func (r *Route) GetService() KongIDGetter {
	return r.Service
}

func (r *Route) GetProtocols() []string {
	return r.Protocols
}

func (r *Route) GetPaths() []string {
	return r.Paths
}

func (r *Route) GetMethods() []string {
	return r.Methods
}

func (r *Route) GetHosts() []string {
	return r.Hosts
}

func (r *Route) SetService(getter KongIDGetter) error {
	if paratils.IsNil(getter) {
		return nil
	}
	if paratils.IsNil(r.Service) {
		r.Service = &KongID{}
	}

	return MarshalKongID(getter, r.Service)
}

func (r *Route) SetProtocols(protocols ...string) error {
	if paratils.IsNil(r.Protocols) {
		r.Protocols = make([]string, 0)
	}
	if len(protocols) < 1 {
		return nil
	}

	copy(protocols, r.Protocols)
	return nil
}

func (r *Route) SetPaths(paths ...string) error {
	if paratils.IsNil(r.Paths) {
		r.Paths = make([]string, 0)
	}
	if len(paths) < 1 {
		return nil
	}

	copy(paths, r.Paths)
	return nil
}

func (r *Route) SetMethods(methods ...string) error {
	if paratils.IsNil(r.Methods) {
		r.Methods = make([]string, 0)
	}
	if len(methods) < 1 {
		return nil
	}

	copy(methods, r.Methods)
	return nil
}

func (r *Route) SetHosts(hosts ...string) error {
	if paratils.IsNil(r.Hosts) {
		r.Hosts = make([]string, 0)
	}
	if len(hosts) < 1 {
		return nil
	}

	copy(hosts, r.Hosts)
	return nil
}

func (r *Route) SetPreserveHost(preserveHost bool) error {
	r.PreserveHost = preserveHost
	return nil
}

func (r *Route) SetStripPath(preserveHost bool) error {
	r.StripPath = preserveHost
	return nil
}

func (r *Route) SetRegexPriority(preserveHost int) error {
	r.RegexPriority = preserveHost
	return nil
}

type RouteGetter interface {
	KongIDGetter
	CreatedAtGetter
	UpdatedAtGetter
	GetHosts() []string
	GetMethods() []string
	GetPaths() []string
	GetProtocols() []string
	GetService() KongIDGetter
}

type RouteSetter interface {
	KongIDSetter
	CreatedAtSetter
	UpdatedAtSetter
	SetHosts(hosts ...string) error
	SetMethods(methods ...string) error
	SetPaths(paths ...string) error
	SetPreserveHost(preserveHost bool) error
	SetProtocols(protocols ...string) error
	SetRegexPriority(preserveHost int) error
	SetService(getter KongIDGetter) error
	SetStripPath(preserveHost bool) error
}

type RouteAccessor interface {
	RouteGetter
	RouteSetter
}

func MarshalRoute(getter RouteGetter, setter RouteSetter) error {
	if paratils.OneIsNil(getter, setter) {
		return nil
	}

	if err := setter.SetID(getter.GetID()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetCreatedAt(getter.GetCreatedAt()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetUpdatedAt(getter.GetUpdatedAt()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetHosts(getter.GetHosts()...); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetMethods(getter.GetMethods()...); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetPaths(getter.GetPaths()...); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetProtocols(getter.GetProtocols()...); err != nil {
		return errgo.New(err)
	}
	return setter.SetService(getter.GetService())
}

type Routes []Route

func (r Routes) GetLength() int {
	return len(r)
}
