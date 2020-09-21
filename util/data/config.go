package data

//Config - config file for subscribe manager
type Config struct {
	IPv6      bool                      `yaml:"ipv6"`
	Groups    []Group                   `yaml:"groups"`
	Selectors []ClashProxyGroupSelector `yaml:"selectors"`
	Rules     []Rule                    `yaml:"rules"`
	Changers  []NameChanger             `yaml:"changers"`
}

//RawConfig - used to load config from file
type RawConfig struct {
	IPv6      bool                         `yaml:"ipv6"`
	Groups    []RawGroup                   `yaml:"groups"`
	Selectors []RawClashProxyGroupSelector `yaml:"selectors"`
	Rules     []Rule                       `yaml:"rules"`
	Changers  []NameChanger                `yaml:"changers"`
}
