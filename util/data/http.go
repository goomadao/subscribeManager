package data

import "go.uber.org/zap/zapcore"

// HTTP struct
type HTTP struct {
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
func (http HTTP) GetType() string {
	return "http"
}

// GetName returns node's name
func (http HTTP) GetName() string {
	return http.Name
}

// SetName sets node's name
func (http HTTP) SetName(name string) {
	http.Name = name
}

// ClashSupport determines whether the node is supported in clash
func (http HTTP) ClashSupport() bool {
	return true
}

// MarshalLogObject provides a method to marshal zap object
func (http HTTP) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", http.Type)
	enc.AddString("server", http.Server)
	enc.AddString("name", http.Name)
	enc.AddInt("port", http.Port)
	enc.AddString("password", http.Password)
	enc.AddString("password", http.Password)
	enc.AddBool("tls", http.ClashTLS)
	enc.AddBool("skipCertVerify", http.SkipCertVerify)
	return nil
}
