package data

//Config - config file for subscribe manager
type Config struct {
	Groups    []Group                   `yaml:"groups"`
	Selectors []ClashProxyGroupSelector `yaml:"selectors"`
	Rules     []Rule                    `yaml:"rules"`
	Changers  []NameChanger             `yaml:"changers"`
}

//RawConfig - used to load config from file
type RawConfig struct {
	Groups    []RawGroup                   `yaml:"groups"`
	Selectors []RawClashProxyGroupSelector `yaml:"selectors"`
	Rules     []Rule                       `yaml:"rules"`
	Changers  []NameChanger                `yaml:"changers"`
}
