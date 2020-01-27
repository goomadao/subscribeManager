package data

import "time"

//Rule - rules
type Rule struct {
	Name        string    `yaml:"name"`
	URLs        []string  `yaml:"urls"`
	CustomRules []string  `yaml:"customRules"`
	Rules       []string  `yaml:"rules"`
	LastUpdate  time.Time `yaml:"lastUpdate"`
}
