package data

import "time"

//Rule - rules
type Rule struct {
	Name        string    `yaml:"name" json:"name"`
	ProxyGroup  string    `yaml:"proxyGroup" json:"proxyGroup"`
	URL         string    `yaml:"url" json:"url"`
	CustomRules []string  `yaml:"customRules" json:"customRules"`
	Rules       []string  `yaml:"rules" json:"rules"`
	LastUpdate  time.Time `yaml:"lastUpdate" json:"lastUpdate"`
}
