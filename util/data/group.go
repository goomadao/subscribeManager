package data

import (
	"time"
)

//Group - a group of ss, ssr or v2ray
type Group struct {
	Name       string    `yaml:"name" json:"name"`
	URL        string    `yaml:"url" json:"url"`
	Nodes      []Node    `yaml:"nodes" json:"nodes"`
	LastUpdate time.Time `yaml:"lastUpdate" json:"lastUpdate"`
}

//RawGroup - used in RawConfig with RawNode
type RawGroup struct {
	Name       string    `yaml:"name" json:"name"`
	URL        string    `yaml:"url" json:"url"`
	Nodes      []RawNode `yaml:"nodes"`
	LastUpdate time.Time `yaml:"lastUpdate"`
}
