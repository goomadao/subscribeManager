package data

import (
	"go.uber.org/zap/zapcore"
)

// Socks5 struct
type Socks5 struct {
	Type           string `yaml:"type" json:"nodeType"`
	Server         string `yaml:"server" json:"server"`
	Name           string `yaml:"name" json:"remarks,omitempty"`
	Port           int    `yaml:"port" json:"port,omitempty"`
	Username       string `yaml:"username,omitempty" json:"username,omitempty"`
	Password       string `yaml:"password,omitempty" json:"password,omitempty"`
	ClashTLS       bool   `yaml:"tls,omitempty" json:"-"`
	TLS            string `yaml:"-" json:"tls,omitempty"`
	SkipCertVerify bool   `yaml:"skipCertVerify,omitempty" json:"skipCertVerify,omitempty"`
}

// GetType returns node's type
func (s Socks5) GetType() string {
	return "socks5"
}

// GetName returns node's name
func (s Socks5) GetName() string {
	return s.Name
}

// SetName sets node's name
func (s *Socks5) SetName(name string) {
	s.Name = name
}

// ClashSupport determines whether the node is supported in clash
func (s Socks5) ClashSupport() bool {
	return true
}

// MarshalLogObject provides a method to marshal zap object
func (s *Socks5) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", s.Type)
	enc.AddString("server", s.Server)
	enc.AddString("name", s.Name)
	enc.AddInt("port", s.Port)
	enc.AddString("password", s.Password)
	enc.AddString("password", s.Password)
	enc.AddBool("tls", s.ClashTLS)
	enc.AddBool("skipCertVerify", s.SkipCertVerify)
	return nil
}
