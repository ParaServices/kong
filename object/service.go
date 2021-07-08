package object

import (
	"github.com/ParaServices/errgo"
	"github.com/ParaServices/paratils"
)

// Service ...
type Service struct {
	KongID
	CreatedAt
	UpdatedAt
	Name
	Protocol       string `json:"protocol,omitempty"`
	Host           string `json:"host,omitempty"`
	Path           string `json:"path,omitempty"`
	Port           int    `json:"port,omitempty"`
	Retries        int    `json:"retries,omitempty"`
	ConnectTimeout int    `json:"connect_timeout,omitempty"`
	WriteTimeout   int    `json:"write_timeout,omitempty"`
	ReadTimeout    int    `json:"read_timeout,omitempty"`
}

func (s *Service) GetProtocol() string {
	return s.Protocol
}

func (s *Service) GetHost() string {
	return s.Host
}

func (s *Service) GetPath() string {
	return s.Path
}

func (s *Service) GetPort() int {
	return s.Port
}

func (s *Service) GetRetries() int {
	return s.Retries
}

func (s *Service) GetConnectTimeout() int {
	return s.ConnectTimeout
}

func (s *Service) GetWriteTimeout() int {
	return s.WriteTimeout
}

func (s *Service) GetReadTimeout() int {
	return s.ReadTimeout
}

func (s *Service) SetProtocol(protocol string) error {
	s.Protocol = protocol
	return nil
}

func (s *Service) SetHost(host string) error {
	s.Host = host
	return nil
}

func (s *Service) SetPath(path string) error {
	s.Path = path
	return nil
}

func (s *Service) SetPort(port int) error {
	s.Port = port
	return nil
}

func (s *Service) SetRetries(retries int) error {
	s.Retries = retries
	return nil
}

func (s *Service) SetConnectTimeout(connecttimeout int) error {
	s.ConnectTimeout = connecttimeout
	return nil
}

func (s *Service) SetWriteTimeout(writetimeout int) error {
	s.WriteTimeout = writetimeout
	return nil
}

func (s *Service) SetReadTimeout(readtimeout int) error {
	s.ReadTimeout = readtimeout
	return nil
}

var _ ServiceAccessor = (*Service)(nil)

type ServiceGetter interface {
	KongIDGetter
	CreatedAtGetter
	UpdatedAtGetter
	NameGetter
	GetConnectTimeout() int
	GetHost() string
	GetPath() string
	GetPort() int
	GetProtocol() string
	GetReadTimeout() int
	GetRetries() int
	GetWriteTimeout() int
}

type ServiceSetter interface {
	KongIDSetter
	CreatedAtSetter
	UpdatedAtSetter
	NameSetter
	SetConnectTimeout(connecttimeout int) error
	SetHost(host string) error
	SetPath(path string) error
	SetPort(port int) error
	SetProtocol(protocol string) error
	SetReadTimeout(readtimeout int) error
	SetRetries(retries int) error
	SetWriteTimeout(writetimeout int) error
}

type ServiceAccessor interface {
	ServiceGetter
	ServiceSetter
}

func MarshalService(getter ServiceGetter, setter ServiceSetter) error {
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
	if err := setter.SetName(getter.GetName()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetConnectTimeout(getter.GetConnectTimeout()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetHost(getter.GetHost()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetPath(getter.GetPath()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetPort(getter.GetPort()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetProtocol(getter.GetProtocol()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetReadTimeout(getter.GetReadTimeout()); err != nil {
		return errgo.New(err)
	}
	if err := setter.SetRetries(getter.GetRetries()); err != nil {
		return errgo.New(err)
	}
	return setter.SetWriteTimeout(getter.GetWriteTimeout())
}

type Services []Service

func (s Services) GetLength() int {
	return len(s)
}
