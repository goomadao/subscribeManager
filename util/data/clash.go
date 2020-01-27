package data

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

//RawClash - used to load config from clash file
type RawClash struct {
	Port               int          `yaml:"port"`
	SocksPort          int          `yaml:"socks-port"`
	AllowLan           bool         `yaml:"allow-lan"`
	Mode               string       `yaml:"rule"`
	LogLevel           string       `yaml:"log-level"`
	ExternalController string       `yaml:"external-controller"`
	Proxy              []RawNode    `yaml:"Proxy"`
	ProxyGroup         []ProxyGroup `yaml:"Proxy Group"`
	Rule               []string     `yaml:"Rule"`
}
