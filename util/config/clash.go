package config

import (
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
		Node2SS(node)
	} else if node.Type == "vmess" {
		Node2Vmess(node)
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
		UpdateSelectorProxies(selector.Name, selector.Type)
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
		clash.Rule = append(clash.Rule, rule.Rules...)
		for _, val := range rule.CustomRules {
			if res, err := addProxyGroupNameAfterRule(val, rule.Name); err == nil {
				clash.Rule = append(clash.Rule, res)
			}
		}
	}
	clashFile, err := yaml.Marshal(clash)
	if err != nil {
		logger.Logger.Warn("Marshal clash config fail",
			zap.Error(err))
		return nil
	}
	return clashFile
}

func clashSupport(node data.Node) bool {
	if node.Cipher == "chacha20" {
		return false
	}
	if len(node.Network) > 0 && node.Network != "ws" {
		return false
	}
	if node.Type == "ss" || node.Type == "vmess" {
		return true
	}
	return false
}
