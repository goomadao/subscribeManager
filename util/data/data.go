package data

import (
	"time"
)

//SS struct
type SS struct {
	Server        string `yaml:"-" json:"server"`
	ID            int    `yaml:"-" json:"id,omitempty"`
	Ratio         int    `yaml:"-" json:"ratio,omitempty"`
	Name          string `yaml:"-" json:"remarks,omitempty"`
	Port          int    `yaml:"-" json:"port,omitempty"`
	Cipher        string `yaml:"-" json:"encryption,omitempty"`
	Password      string `yaml:"-" json:"password,omitempty"`
	Plugin        string `yaml:"-" json:"plugin,omitempty"`
	PluginOptions string `yaml:"-" json:"plugin_options,omitempty"`
}

//SSR struct
type SSR struct {
	Server        string
	Port          int
	Cipher        string
	Password      string
	Name          string
	Protocol      string
	ProtocolParam string
	Obfs          string
	ObfsParam     string
	Group         string
}

//Vmess link json format
type Vmess struct {
	Host    string `yaml:"-" json:"host,omitempty"`
	Path    string `yaml:"-" json:"path,omitempty"`
	TLS     string `yaml:"-" json:"tls,omitempty"`
	Server  string `yaml:"-" json:"add,omitempty"`
	Port    int    `yaml:"-" json:"port,omitempty"`
	AlterID int    `yaml:"-" json:"aid,omitempty"`
	Network string `yaml:"-" json:"net,omitempty"`
	Type    string `yaml:"-" json:"type,omitempty"`
	V       string `yaml:"-" json:"v,omitempty"`
	Name    string `yaml:"-" json:"ps,omitempty"`
	UUID    string `yaml:"-" json:"id,omitempty"`
	Class   int    `yaml:"-" json:"class,omitempty"`
}

//SSD - ssd subscribe
type SSD struct {
	Airport       string  `yaml:"-" json:"airport"`
	Port          int     `yaml:"-" json:"port"`
	Cipher        string  `yaml:"-" json:"encryption"`
	Password      string  `yaml:"-" json:"password"`
	Servers       []SS    `yaml:"-" json:"servers"`
	Plugin        string  `yaml:"-" json:"plugin,omitempty"`
	PluginOptions string  `yaml:"-" json:"plugin_options,omitempty"`
	TrafficUsed   float64 `yaml:"-" json:"traffic_used,omitempty"`
	TrafficTotal  float64 `yaml:"-" json:"traffic_total,omitempty"`
	Expiry        string  `yaml:"-" json:"expiry,omitempty"`
	URL           string  `yaml:"-" json:"url,omitempty"`
}

//Plugin - ss plugin in clash
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

	SS    SS    `yaml:"-"`
	SSR   SSR   `yaml:"-"`
	Vmess Vmess `yaml:"-"`
}

//ProxyGroup - clash proxy group option
type ProxyGroup struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	URL      string   `yaml:"url,omitempty"`
	Interval int      `yaml:"interval,omitempty"`
	Proxies  []string `yaml:"proxies"`
}

//Clash - clash config
type Clash struct {
	Port               int          `yaml:"port"`
	SocksPort          int          `yaml:"socks-port"`
	AllowLan           bool         `yaml:"allow-lan"`
	Mode               string       `yaml:"rule"`
	LogLevel           string       `yaml:"log-level"`
	ExternalController string       `yaml:"external-controller"`
	Proxy              []Node       `yaml:"Proxy"`
	ProxyGroup         []ProxyGroup `yaml:"Proxy Group"`
	Rule               []string     `yaml:"Rule"`
}

//Group - a group of ss, ssr or v2ray
type Group struct {
	Name       string    `yaml:"name" json:"name"`
	URL        string    `yaml:"url" json:"url"`
	Nodes      []Node    `yaml:"nodes"`
	LastUpdate time.Time `yaml:"lastUpdate"`
}

//ProxySelector - using regex to select specific proxies in group
type ProxySelector struct {
	GroupName string `yaml,json:"groupName"`
	Regex     string `yaml,json:"regex"`
}

//ClashProxyGroupSelector - determine how to select proxies from groups for clash
type ClashProxyGroupSelector struct {
	Name          string          `yaml:"name"`
	Type          string          `yaml:"type"`
	URL           string          `yaml:"url,omitempty"`
	Interval      int             `yaml:"interval,omitempty"`
	ProxyGroups   []string        `yaml:"proxyGroups"`
	ProxySelector []ProxySelector `yaml:"proxySelector"`
	Proxies       []Node          `yaml:"proxies"`
}

//Rule - rules
type Rule struct {
	Name        string    `yaml:"name"`
	URLs        []string  `yaml:"urls"`
	CustomRules []string  `yaml:"customRules"`
	Rules       []string  `yaml:"rules"`
	LastUpdate  time.Time `yaml:"lastUpdate"`
}

//NameChanger - change name, such as adding emojis
type NameChanger struct {
	Emoji string `yaml:"emoji"`
	Regex string `yaml:"regex"`
}

//Config - config file for subscribe manager
type Config struct {
	Groups    []Group                   `yaml:"groups"`
	Selectors []ClashProxyGroupSelector `yaml:"selectors"`
	Rules     []Rule                    `yaml:"rules"`
	Changers  []NameChanger             `yaml:"changers"`
}
