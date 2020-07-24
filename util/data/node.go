package data

import "go.uber.org/zap/zapcore"

//Node - ss,ssr or vmess in clash
type Node interface {
	GetType() string
	GetName() string
	SetName(string)
	ClashSupport() bool
	MarshalLogObject(enc zapcore.ObjectEncoder) error
}

//RawNode - used to read from config file
type RawNode struct {
	//General
	Type     string `json:"-" yaml:"type"`
	Cipher   string `json:"-" yaml:"cipher"`
	Password string `json:"-" yaml:"password"`
	Name     string `json:"-" yaml:"name"`
	Server   string `json:"-" yaml:"server"`
	Port     int    `json:"-" yaml:"port"`
	//SS
	Plugin     string `json:"-" yaml:"plugin,omitempty"`
	PluginOpts Plugin `json:"-" yaml:"plugin-opts,omitempty"`
	//SSR
	Protocol      string `json:"-" yaml:"protocol,omitempty"`
	ProtocolParam string `json:"-" yaml:"protocol-param,omitempty"`
	Obfs          string `json:"-" yaml:"obfs,omitempty"`
	ObfsParam     string `json:"-" yaml:"obfs-param,omitempty"`
	Group         string `json:"-" yaml:"group,omitempty"`
	//Vmess
	UUID      string    `json:"-" yaml:"uuid,omitempty"`
	AlterID   int       `json:"-" yaml:"alterId,omitempty"`
	TLS       bool      `json:"-" yaml:"tls,omitempty"`
	Network   string    `json:"-" yaml:"network,omitempty"`
	WSPath    string    `json:"-" yaml:"ws-path,omitempty"`
	WSHeaders WSHeaders `json:"-" yaml:"ws-headers,omitempty"`

	SS    SS    `yaml:"-"`
	SSR   SSR   `yaml:"-"`
	Vmess Vmess `yaml:"-"`
}

//MarshalLogObject provides a method to marshal zap object
func (n *RawNode) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", n.Type)
	enc.AddString("server", n.Server)
	enc.AddString("name", n.Name)
	enc.AddInt("port", n.Port)
	enc.AddString("cipher", n.Cipher)
	enc.AddString("password", n.Password)
	if n.Type == "ss" {
		enc.AddString("plugin", n.Plugin)
		enc.AddString("obfs", n.PluginOpts.Obfs)
		enc.AddString("obfsHost", n.PluginOpts.ObfsHost)
	} else if n.Type == "ssr" {
		enc.AddString("protocol", n.Protocol)
		enc.AddString("protocolParam", n.ProtocolParam)
		enc.AddString("obfs", n.Obfs)
		enc.AddString("obfsParam", n.ObfsParam)
		enc.AddString("group", n.Group)
	} else if n.Type == "vmess" {
		enc.AddString("uuid", n.UUID)
		enc.AddInt("alterID", n.AlterID)
		enc.AddBool("tls", n.TLS)
		enc.AddString("network", n.Network)
		enc.AddString("path", n.WSPath)
		enc.AddString("WSHeaders.Host", n.WSHeaders.Host)
	}
	return nil
}
