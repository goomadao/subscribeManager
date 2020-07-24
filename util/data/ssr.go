package data

import "go.uber.org/zap/zapcore"

//SSR struct
type SSR struct {
	Type          string `yaml:"type" json:"nodeType"`
	Server        string `yaml:"server" json:"server"`
	Port          int    `yaml:"port" json:"port"`
	Cipher        string `yaml:"cipher" json:"encryption"`
	Password      string `yaml:"password" json:"password"`
	Name          string `yaml:"name" json:"remarks"`
	Protocol      string `yaml:"protocol" json:"protocol"`
	ProtocolParam string `yaml:"protocol-param,omitempty" json:"protocol_param"`
	Obfs          string `yaml:"obfs" json:"obfs"`
	ObfsParam     string `yaml:"obfs-param,omitempty" json:"obfs_param"`
	Group         string `yaml:"-" json:"group"`
}

//GetType returns node's type
func (ssr SSR) GetType() string {
	return "ssr"
}

//GetName returns node's name
func (ssr SSR) GetName() string {
	return ssr.Name
}

//SetName sets node's name
func (ssr *SSR) SetName(name string) {
	ssr.Name = name
}

//ClashSupport determines whether the node is supported in clash
func (ssr SSR) ClashSupport() bool {
	if streamCipherSupported(ssr.Cipher) && ssrObfsSupported(ssr.Obfs) && ssrProtocolSupported(ssr.Protocol) {
		return true
	}
	return false
}

//MarshalLogObject provides a method to marshal zap object
func (ssr *SSR) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", ssr.Type)
	enc.AddString("server", ssr.Server)
	enc.AddInt("port", ssr.Port)
	enc.AddString("cipher", ssr.Cipher)
	enc.AddString("password", ssr.Password)
	enc.AddString("name", ssr.Name)
	enc.AddString("protocol", ssr.Protocol)
	enc.AddString("protocolParam", ssr.ProtocolParam)
	enc.AddString("obfs", ssr.Obfs)
	enc.AddString("obfsParam", ssr.ObfsParam)
	enc.AddString("group", ssr.Group)
	return nil
}
