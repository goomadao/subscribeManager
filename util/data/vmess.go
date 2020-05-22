package data

import "go.uber.org/zap/zapcore"

//Vmess link json format
type Vmess struct {
	ClashType    string    `yaml:"type" json:"nodeType"`
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

//MarshalLogObject provides a method to marshal zap object
func (vmess *Vmess) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("clashType", vmess.ClashType)
	enc.AddString("WSHeaders.Host", vmess.WSHeaders.Host)
	enc.AddString("clashNetwork", vmess.ClashNetwork)
	enc.AddBool("clashTLS", vmess.ClashTLS)
	enc.AddString("cipher", vmess.Cipher)
	enc.AddString("host", vmess.Host)
	enc.AddString("path", vmess.Path)
	enc.AddString("tls", vmess.TLS)
	enc.AddString("server", vmess.Server)
	enc.AddInt("port", vmess.Port)
	enc.AddInt("alterID", vmess.AlterID)
	enc.AddString("network", vmess.Network)
	enc.AddString("type", vmess.Type)
	enc.AddString("v", vmess.V)
	enc.AddString("name", vmess.Name)
	enc.AddString("uuid", vmess.UUID)
	enc.AddInt("class", vmess.Class)
	return nil
}
