package data

//SS struct
type SS struct {
	Type          string `yaml:"type" json:"-"`
	Server        string `yaml:"server" json:"server"`
	ID            int    `yaml:"-" json:"id,omitempty"`
	Ratio         int    `yaml:"-" json:"ratio,omitempty"`
	Name          string `yaml:"name" json:"remarks,omitempty"`
	Port          int    `yaml:"port" json:"port,omitempty"`
	Cipher        string `yaml:"cipher" json:"encryption,omitempty"`
	Password      string `yaml:"password" json:"password,omitempty"`
	Plugin        string `yaml:"plugin,omitempty" json:"plugin,omitempty"`
	PluginOptions string `yaml:"-" json:"plugin_options,omitempty"`
	PluginOpts    Plugin `yaml:"plugin-opts,omitempty" json:"-"`
}

//Plugin - ss plugin in clash
type Plugin struct {
	Obfs     string `yaml:"mode"`
	ObfsHost string `yaml:"host"`
}

//GetType returns node's type
func (ss SS) GetType() string {
	return "ss"
}

//GetName returns node's name
func (ss SS) GetName() string {
	return ss.Name
}

//SetName sets node's name
func (ss *SS) SetName(name string) {
	ss.Name = name
}

//ClashSupport determines whether the node is supported in clash
func (ss SS) ClashSupport() bool {
	if ss.Cipher == "chacha20" {
		return false
	}
	if len(ss.Plugin) > 0 && ss.Plugin != "obfs" {
		return false
	}
	return true
}

//ClashRSupport determines whether the node is supported in clashr
func (ss SS) ClashRSupport() bool {
	return ss.ClashSupport()
}
