package data

import "go.uber.org/zap/zapcore"

// Trojan struct
type Trojan struct {
	Type           string `yaml:"type" json:"nodeType"`
	Server         string `yaml:"server" json:"server"`
	Name           string `yaml:"name" json:"remarks,omitempty"`
	Port           int    `yaml:"port" json:"port,omitempty"`
	Password       string `yaml:"password" json:"password,omitempty"`
	Sni            string `yaml:"sni,omitempty" json:"sni,omitempty"`
	SkipCertVerify bool   `yaml:"skip-cert-verify,omitempty" json:"verify"`
}

// GetType returns node's type
func (t Trojan) GetType() string {
	return t.Type
}

// GetName returns node's name
func (t Trojan) GetName() string {
	return t.Name
}

// SetName sets node's name
func (t *Trojan) SetName(name string) {
	t.Name = name
}

// ClashSupport determines whether the node is supported in clash
func (t Trojan) ClashSupport() bool {
	return true
}

// ClashRSupport determines whether the node is supported in clashr
func (t Trojan) ClashRSupport() bool {
	return t.ClashSupport()
}

// MarshalLogObject provides a method to marshal zap object
func (t *Trojan) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", t.Type)
	enc.AddString("server", t.Server)
	enc.AddString("name", t.Name)
	enc.AddInt("port", t.Port)
	enc.AddString("password", t.Password)
	enc.AddString("sni", t.Sni)
	enc.AddBool("skip-cert-verify", t.SkipCertVerify)
	return nil
}
