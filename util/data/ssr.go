package data

//SSR struct
type SSR struct {
	Type          string `yaml:"type" json:"nodeType"`
	Server        string `yaml:"server" json:"server"`
	Port          int    `yaml:"port" json:"port"`
	Cipher        string `yaml:"cipher" json:"encryption"`
	Password      string `yaml:"password" json:"password"`
	Name          string `yaml:"name" json:"remarks"`
	Protocol      string `yaml:"protocol" json:"protocol"`
	ProtocolParam string `yaml:"protocolparam" json:"protocol_param"`
	Obfs          string `yaml:"obfs" json:"obfs"`
	ObfsParam     string `yaml:"obfsparam" json:"obfs_param"`
	Group         string `yaml:"-" json:"group"`
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
