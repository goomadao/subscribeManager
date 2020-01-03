package data

import (
	"time"
)

//Plugin - ss plugin
type Plugin struct {
	Obfs     string `yaml:"mode"`
	ObfsHost string `yaml:"host"`
}

//WSHeaders - clash config option for vmess
type WSHeaders struct {
	Host string `yaml:"Host,omitempty"`
}

//Node - ss,ssr or vmess in clash
type Node struct {
	//Clash
	Type          string    `json:"-" yaml:"type"`
	Cipher        string    `json:"-" yaml:"cipher"`
	Password      string    `json:"-" yaml:"password"`
	Name          string    `json:"-" yaml:"name"`
	Server        string    `json:"-" yaml:"server"`
	Port          int       `json:"-" yaml:"port"`
	Plugin        string    `json:"-" yaml:"plugin,omitempty"`
	PluginOpts    Plugin    `json:"-" yaml:"plugin-opts,omitempty"`
	Protocol      string    `json:"-" yaml:"protocol,omitempty"`
	ProtocolParam string    `json:"-" yaml:"protocolparam,omitempty"`
	Obfs          string    `json:"-" yaml:"obfs,omitempty"`
	ObfsParam     string    `json:"-" yaml:"obfsparam,omitempty"`
	Group         string    `json:"-" yaml:"group,omitempty"`
	UUID          string    `json:"-" yaml:"uuid,omitempty"`
	AlterID       int       `json:"-" yaml:"alterId,omitempty"`
	TLS           bool      `json:"-" yaml:"tls,omitempty"`
	Network       string    `json:"-" yaml:"network,omitempty"`
	WSPath        string    `json:"-" yaml:"ws-path,omitempty"`
	WSHeaders     WSHeaders `json:"-" yaml:"ws-headers,omitempty"`
	//Vmess
	VmessHost  string `yaml:"-" json:"host,omitempty"`
	VmessPath  string `yaml:"-" json:"path,omitempty"`
	VmessTLS   string `yaml:"-" json:"tls,omitempty"`
	VmessAdd   string `yaml:"-" json:"add,omitempty"`
	VmessPort  int    `yaml:"-" json:"port,omitempty"`
	VmessAid   int    `yaml:"-" json:"aid,omitempty"`
	VmessNet   string `yaml:"-" json:"net,omitempty"`
	VmessType  string `yaml:"-" json:"type,omitempty"`
	VmessV     string `yaml:"-" json:"v,omitempty"`
	VmessPs    string `yaml:"-" json:"ps,omitempty"`
	VmessID    string `yaml:"-" json:"id,omitempty"`
	VmessClass int    `yaml:"-" json:"class,omitempty"`
}

//Clash - clash config
type Clash struct {
	Port               int    `yaml:"port"`
	SocksPort          int    `yaml:"socks-port"`
	AllowLan           bool   `yaml:"allow-lan"`
	Mode               string `yaml:"rule"`
	LogLevel           string `yaml:"log-level"`
	ExternalController string `yaml:"external-controller"`
	Proxy              []Node `yaml:"Proxy"`
}

//Group - a group of ss, ssr or v2ray
type Group struct {
	Name       string    `yaml:"name"`
	URL        string    `yaml:"url"`
	Nodes      []Node    `yaml:"nodes"`
	LastUpdate time.Time `yaml:"lastUpdate"`
}

//Config - config file for subscribe manager
type Config struct {
	Groups []Group `yaml:"Groups"`
}
