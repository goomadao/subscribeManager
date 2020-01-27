package data

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

//RawClashProxyGroupSelector - determine how to select proxies from groups for clash
type RawClashProxyGroupSelector struct {
	Name          string          `yaml:"name"`
	Type          string          `yaml:"type"`
	URL           string          `yaml:"url,omitempty"`
	Interval      int             `yaml:"interval,omitempty"`
	ProxyGroups   []string        `yaml:"proxyGroups"`
	ProxySelector []ProxySelector `yaml:"proxySelector"`
	Proxies       []RawNode       `yaml:"proxies"`
}
