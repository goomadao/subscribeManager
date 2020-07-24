package config

import (
	"errors"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

func decodeClash(bts []byte) (nodes []data.Node, err error) {
	clash := data.RawClash{}
	err = yaml.Unmarshal(bts, &clash)
	if err != nil {
		logger.Logger.Warn("Decode clash config file fail",
			zap.Error(err))
		return nil, err
	}
	nodes = decodeClashProxies(clash.Proxies)
	return nodes, nil
}

func decodeClashProxies(proxies []data.RawNode) (nodes []data.Node) {
	for _, proxy := range proxies {
		node, err := decodeClashProxy(proxy)
		if err != nil {
			continue
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func decodeClashProxy(proxy data.RawNode) (data.Node, error) {
	if proxy.Type == "ss" {
		Node2SS(&proxy)
		return &proxy.SS, nil
	} else if proxy.Type == "vmess" {
		Node2Vmess(&proxy)
		return &proxy.Vmess, nil
	} else if proxy.Type == "ssr" {
		Node2SSR(&proxy)
		return &proxy.SSR, nil
	}
	logger.Logger.Warn("Unsupported type")
	return nil, errors.New("Unsupported type")
}

//GenerateClashConfig generates clash config file
func GenerateClashConfig() []byte {
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
			if node.ClashSupport() {
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
				clash.Proxies = append(clash.Proxies, node)
			}
		}
	}
	for idx := range config.Selectors {
		selector := &config.Selectors[idx]
		UpdateSelectorProxies(selector.Name, selector.Type)
		var proxies []string
		for _, val := range selector.Proxies {
			if val.ClashSupport() {
				proxies = append(proxies, val.GetName())
			}
		}
		if len(selector.ProxyGroups) == 0 && len(proxies) == 0 {
			proxies = []string{"DIRECT"}
		}
		clash.ProxyGroups = append(clash.ProxyGroups, data.ProxyGroups{
			Name:     selector.Name,
			Type:     selector.Type,
			URL:      selector.URL,
			Interval: selector.Interval,
			Proxies:  append(selector.ProxyGroups, proxies...),
		})
	}
	for _, rule := range config.Rules {
		clash.Rules = append(clash.Rules, rule.Rules...)
		for _, val := range rule.CustomRules {
			if res, err := addProxyGroupNameAfterRule(val, rule.ProxyGroup); err == nil {
				clash.Rules = append(clash.Rules, res)
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
