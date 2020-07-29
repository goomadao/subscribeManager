package data

// ProxyGroups - clash proxy group option
type ProxyGroups struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	URL      string   `yaml:"url,omitempty"`
	Interval int      `yaml:"interval,omitempty"`
	Proxies  []string `yaml:"proxies"`
}

//Clash - clash config
type Clash struct {
	Port               int           `yaml:"port"`
	SocksPort          int           `yaml:"socks-port"`
	AllowLan           bool          `yaml:"allow-lan"`
	Mode               string        `yaml:"rule"`
	LogLevel           string        `yaml:"log-level"`
	ExternalController string        `yaml:"external-controller"`
	Proxies            []Node        `yaml:"proxies"`
	ProxyGroups        []ProxyGroups `yaml:"proxy-groups"`
	Rules              []string      `yaml:"rules"`
}

//RawClash - used to load config from clash file
type RawClash struct {
	Port               int           `yaml:"port"`
	SocksPort          int           `yaml:"socks-port"`
	AllowLan           bool          `yaml:"allow-lan"`
	Mode               string        `yaml:"rule"`
	LogLevel           string        `yaml:"log-level"`
	ExternalController string        `yaml:"external-controller"`
	Proxies            []RawNode     `yaml:"proxies"`
	ProxyGroups        []ProxyGroups `yaml:"proxy-groups"`
	Rules              []string      `yaml:"rules"`
}
