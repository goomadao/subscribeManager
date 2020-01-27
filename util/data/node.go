package data

//Node - ss,ssr or vmess in clash
type Node interface {
	GetType() string
	GetName() string
	SetName(string)
	ClashSupport() bool
	ClashRSupport() bool
}

//RawNode - used to read from config file
type RawNode struct {
	//General
	Type     string `json:"-" yaml:"type"`
	Cipher   string `json:"-" yaml:"cipher"`
	Password string `json:"-" yaml:"password"`
	Name     string `json:"-" yaml:"name"`
	Server   string `json:"-" yaml:"server"`
	Port     int    `json:"-" yaml:"port"`
	//SS
	Plugin     string `json:"-" yaml:"plugin,omitempty"`
	PluginOpts Plugin `json:"-" yaml:"plugin-opts,omitempty"`
	//SSR
	Protocol      string `json:"-" yaml:"protocol,omitempty"`
	ProtocolParam string `json:"-" yaml:"protocolparam,omitempty"`
	Obfs          string `json:"-" yaml:"obfs,omitempty"`
	ObfsParam     string `json:"-" yaml:"obfsparam,omitempty"`
	Group         string `json:"-" yaml:"group,omitempty"`
	//Vmess
	UUID      string    `json:"-" yaml:"uuid,omitempty"`
	AlterID   int       `json:"-" yaml:"alterId,omitempty"`
	TLS       bool      `json:"-" yaml:"tls,omitempty"`
	Network   string    `json:"-" yaml:"network,omitempty"`
	WSPath    string    `json:"-" yaml:"ws-path,omitempty"`
	WSHeaders WSHeaders `json:"-" yaml:"ws-headers,omitempty"`

	SS    SS    `yaml:"-"`
	SSR   SSR   `yaml:"-"`
	Vmess Vmess `yaml:"-"`
}
