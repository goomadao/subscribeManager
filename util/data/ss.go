package data

import "go.uber.org/zap/zapcore"

//SS struct
type SS struct {
	Type          string `yaml:"type" json:"nodeType"`
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
	if !ssCipherSupported(ss.Cipher) {
		return false
	}
	if len(ss.Plugin) > 0 && !ssPluginSupported(ss.Plugin) {
		return false
	}
	return true
}

// //MarshalLogObject provides a method to marshal zap object
// func (p *Plugin) ObjectMarshaler(enc zapcore.ObjectEncoder) error {
// 	enc.AddString("obfs", p.Obfs)
// 	enc.AddString("obfsHost", p.ObfsHost)
// 	return nil
// }

//MarshalLogObject provides a method to marshal zap object
func (ss *SS) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("type", ss.Type)
	enc.AddString("server", ss.Server)
	enc.AddString("name", ss.Name)
	enc.AddInt("port", ss.Port)
	enc.AddString("cipher", ss.Cipher)
	enc.AddString("password", ss.Password)
	enc.AddString("plugin", ss.Plugin)
	enc.AddString("pluginOptions", ss.PluginOptions)
	enc.AddString("obfs", ss.PluginOpts.Obfs)
	enc.AddString("obfsHost", ss.PluginOpts.ObfsHost)
	return nil
}
