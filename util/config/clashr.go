package config

import (
	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

//GenerateClashRConfig generates clash config file
func GenerateClashRConfig() []byte {
	clash := data.Clash{
		Port:               7890,
		SocksPort:          7891,
		AllowLan:           true,
		Mode:               "Rule",
		LogLevel:           "info",
		ExternalController: "0.0.0.0:9090",
	}
	proxies := make(map[string]data.Node)
	originalNames := make(map[data.Node]string)
	for idx := range config.Groups {
		group := config.Groups[idx]
		for _, node := range group.Nodes {
			if node.ClashRSupport() {
				originalName := node.GetName()
				for true {
					if _, ok := proxies[node.GetName()]; !ok {
						break
					}
					node.SetName(node.GetName() + "#")
				}
				if originalName != node.GetName() {
					originalNames[node] = originalName
				}
				proxies[node.GetName()] = node
				clash.Proxy = append(clash.Proxy, node)
			}
		}
	}
	for idx := range config.Selectors {
		selector := &config.Selectors[idx]
		UpdateSelectorProxies(selector.Name, selector.Type)
		var proxies []string
		for _, val := range selector.Proxies {
			if val.ClashRSupport() {
				proxies = append(proxies, val.GetName())
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
	for node, name := range originalNames {
		node.SetName(name)
	}
	WriteToFile()
	return clashFile
}
