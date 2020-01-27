package data

//SSR struct
type SSR struct {
	Type          string `yaml:"type" json:"-"`
	Server        string `yaml:"server"`
	Port          int    `yaml:"port"`
	Cipher        string `yaml:"cipher"`
	Password      string `yaml:"password"`
	Name          string `yaml:"name"`
	Protocol      string `yaml:"protocol"`
	ProtocolParam string `yaml:"protocolparam"`
	Obfs          string `yaml:"obfs"`
	ObfsParam     string `yaml:"obfsparam"`
	Group         string `yaml:"-"`
}

//GetType returns node's type
func (ssr SSR) GetType() string {
	return "ssr"
}

//GetName returns node's name
func (ssr SSR) GetName() string {
	return ssr.Name
}

//SetName sets node's name
func (ssr *SSR) SetName(name string) {
	ssr.Name = name
}

//ClashSupport determines whether the node is supported in clash
func (ssr SSR) ClashSupport() bool {
	return false
}

//ClashRSupport determines whether the node is supported in clashr
func (ssr SSR) ClashRSupport() bool {
	return true
}
