package data

import (
	"time"
)

type Plugin struct {
	Obfs     string `yaml:"mode"`
	ObfsHost string `yaml:"host"`
}

type Node struct {
	Type          string `yaml:"type"`
	Cipher        string `yaml:"cipher"`
	Password      string `yaml:"password"`
	Name          string `yaml:"name"`
	Server        string `yaml:"server"`
	Port          int    `yaml:"port"`
	Plugin        string `yaml:"plugin,omitempty"`
	PluginOpts    Plugin `yaml:"plugin-opts,omitempty"`
	Protocol      string `yaml:"protocol,omitempty"`
	ProtocolParam string `yaml:"protocolparam,omitempty"`
	Obfs          string `yaml:"obfs,omitempty"`
	ObfsParam     string `yaml:"obfsparam,omitempty"`
	Group         string `yaml:"group,omitempty"`
}

type Group struct {
	Name       string    `yaml:"name"`
	Url        string    `yaml:"url"`
	Nodes      []Node    `yaml:"nodes,flow"`
	LastUpdate time.Time `yaml:"lastUpdate"`
}

type Config struct {
	Groups []Group `yaml:"Groups"`
}

type Clash struct {
	Port               int    `yaml:"port"`
	SocksPort          int    `yaml:"socks-port"`
	AllowLan           bool   `yaml:"allow-lan"`
	Mode               string `yaml:"rule"`
	LogLevel           string `yaml:"log-level"`
	ExternalController string `yaml:"external-controller"`
	Proxy              []Node `yaml:"Proxy"`
}
