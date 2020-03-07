package data

//NameChanger - change name, such as adding emojis
type NameChanger struct {
	Emoji string `yaml:"emoji" json:"emoji"`
	Regex string `yaml:"regex" json:"regex"`
}
