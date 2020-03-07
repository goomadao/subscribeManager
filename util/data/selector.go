package data

//ProxySelector - using regex to select specific proxies in group
type ProxySelector struct {
	GroupName string `yaml,json:"groupName" json:"groupName"`
	Regex     string `yaml,json:"regex" json:"regex"`
}

//ClashProxyGroupSelector - determine how to select proxies from groups for clash
type ClashProxyGroupSelector struct {
	Name           string          `yaml:"name" json:"name"`
	Type           string          `yaml:"type" json:"type"`
	URL            string          `yaml:"url,omitempty" json:"url"`
	Interval       int             `yaml:"interval,omitempty" json:"interval"`
	ProxyGroups    []string        `yaml:"proxyGroups" json:"proxyGroups"`
	ProxySelectors []ProxySelector `yaml:"proxySelectors" json:"proxySelectors"`
	Proxies        []Node          `yaml:"proxies" json:"proxies,omitempty"`
}

//RawClashProxyGroupSelector - determine how to select proxies from groups for clash
type RawClashProxyGroupSelector struct {
	Name           string          `yaml:"name"`
	Type           string          `yaml:"type"`
	URL            string          `yaml:"url,omitempty"`
	Interval       int             `yaml:"interval,omitempty"`
	ProxyGroups    []string        `yaml:"proxyGroups"`
	ProxySelectors []ProxySelector `yaml:"proxySelectors"`
	Proxies        []RawNode       `yaml:"proxies"`
}
