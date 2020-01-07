package config

import (
	"fmt"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

func decodeClash(bts []byte) (nodes []data.Node, err error) {
	clash := data.Clash{}
	err = yaml.Unmarshal(bts, &clash)
	if err != nil {
		logger.Logger.Warn("Decode clash config file fail",
			zap.Error(err))
		return nil, err
	}
	for _, node := range clash.Proxy {
		translateNode(&node)
	}
	return clash.Proxy, nil
}

func translateNode(node *data.Node) {
	if node.Type == "ss" {
		node.SS = data.SS{
			Server:   node.Server,
			Name:     node.Name,
			Port:     node.Port,
			Cipher:   node.Cipher,
			Password: node.Password,
			Plugin:   node.Plugin,
			PluginOptions: "obfs=" +
				node.PluginOpts.Obfs +
				";obfs-host=" +
				node.PluginOpts.ObfsHost,
		}
	} else if node.Type == "vmess" {
		node.Vmess = data.Vmess{
			Host:    node.WSHeaders.Host,
			Path:    node.WSPath,
			Server:  node.Server,
			Port:    node.Port,
			AlterID: node.AlterID,
			Network: node.Network,
			Type:    "none",
			V:       "2",
			Name:    node.Name,
			UUID:    node.UUID,
			Class:   1,
		}
		if node.TLS {
			node.Vmess.TLS = "tls"
		}
	}
}

//GenerateClashConfig generates clash config file
func GenerateClashConfig() []byte {
	InitConfig()
	clash := data.Clash{
		Port:               7890,
		SocksPort:          7891,
		AllowLan:           true,
		Mode:               "Rule",
		LogLevel:           "info",
		ExternalController: "0.0.0.0:9090",
	}
	for _, group := range config.Groups {
		for _, node := range group.Nodes {
			if clashSupport(node) {
				clash.Proxy = append(clash.Proxy, node)
			}
		}
	}
	for _, selector := range config.Selectors {
		updateSelectorProxies(selector.Name, selector.Type)
		var proxies []string
		for _, val := range selector.Proxies {
			if clashSupport(val) {
				proxies = append(proxies, val.Name)
			}
		}
		clash.ProxyGroup = append(clash.ProxyGroup, data.ProxyGroup{
			Name:     selector.Name,
			Type:     selector.Type,
			URL:      selector.URL,
			Interval: selector.Interval,
			Proxies:  append(selector.ProxyGroups, proxies...),
		})
	}
	for _, rule := range config.Rules {
		// content := rule.Content
		// rules := strings.Split(content, "\n")
		// for _, val := range rules {
		// 	if len(val) > 0 && strings.Index(val, "#") != 0 {
		// 		clash.Rule = append(clash.Rule, val)
		// 	}
		// }
		clash.Rule = append(clash.Rule, rule.Rules...)
	}
	clashFile, err := yaml.Marshal(clash)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println(string(clashFile))
	return clashFile
}

func clashSupport(node data.Node) bool {
	if node.Cipher == "chacha20" {
		return false
	}
	if node.Type == "ss" || node.Type == "vmess" {
		return true
	}
	return false
}