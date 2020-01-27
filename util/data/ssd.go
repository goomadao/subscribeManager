package data

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
