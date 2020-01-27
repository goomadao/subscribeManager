package data

//Vmess link json format
type Vmess struct {
	ClashType    string    `yaml:"type" json:"-"`
	WSHeaders    WSHeaders `yaml:"ws-headers,omitempty" json:"-"`
	ClashNetwork string    `yaml:"network,omitempty" json:"-"`
	ClashTLS     bool      `yaml:"tls" json:"-"`
	Cipher       string    `yaml:"cipher" json:"-"`
	Host         string    `yaml:"-" json:"host,omitempty"`
	Path         string    `yaml:"ws-path,omitempty" json:"path,omitempty"`
	TLS          string    `yaml:"-" json:"tls,omitempty"`
	Server       string    `yaml:"server" json:"add,omitempty"`
	Port         int       `yaml:"port" json:"port,omitempty"`
	AlterID      int       `yaml:"alterId" json:"aid,omitempty"`
	Network      string    `yaml:"-" json:"net,omitempty"`
	Type         string    `yaml:"-" json:"type,omitempty"`
	V            string    `yaml:"-" json:"v,omitempty"`
	Name         string    `yaml:"name" json:"ps,omitempty"`
	UUID         string    `yaml:"uuid" json:"id,omitempty"`
	Class        int       `yaml:"-" json:"class,omitempty"`
}

//WSHeaders - clash config option for vmess
type WSHeaders struct {
	Host string `yaml:"Host,omitempty"`
}

//GetType returns node's type
func (vmess Vmess) GetType() string {
	return "vmess"
}

//GetName returns node's name
func (vmess Vmess) GetName() string {
	return vmess.Name
}

//SetName sets node's name
func (vmess *Vmess) SetName(name string) {
	vmess.Name = name
}

//ClashSupport determines whether the node is supported in clash
func (vmess Vmess) ClashSupport() bool {
	if len(vmess.ClashNetwork) > 0 && vmess.ClashNetwork != "ws" {
		return false
	}
	return true
}

//ClashRSupport determines whether the node is supported in clashr
func (vmess Vmess) ClashRSupport() bool {
	return vmess.ClashSupport()
}
