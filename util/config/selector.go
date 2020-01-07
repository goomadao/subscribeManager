package config

import "github.com/goomadao/subscribeManager/util/data"

import "github.com/goomadao/subscribeManager/util/logger"

import "errors"

import "go.uber.org/zap"

import "regexp"

//AddSelector adds new selector
func AddSelector(selector data.ClashProxyGroupSelector) error {
	loadConfig()
	err := selectorDuplicate(selector)
	if err != nil {
		return err
	}
	config.Selectors = append(config.Selectors, selector)
	err = writeToFile()
	if err != nil {
		logger.Logger.Panic("Selector write to file fail",
			zap.Error(err))
	}
	logger.Logger.Info("Selector write to file success")
	go updateSelectorProxies(selector.Name, selector.Type)
	return nil
}

func selectorDuplicate(selector data.ClashProxyGroupSelector) error {
	for _, val := range config.Selectors {
		if val.Name == selector.Name && val.Type == selector.Type {
			logger.Logger.Warn("Selector duplicates")
			return errors.New("Selector duplicates")
		}
	}
	return nil
}

func updateSelectorProxies(name string, selectType string) error {
	loadConfig()
	index := -1
	for i, val := range config.Selectors {
		if val.Name == name && val.Type == selectType {
			index = i
			break
		}
	}
	if index == -1 {
		logger.Logger.Warn("No such selector")
		return errors.New("No such selector")
	}
	for _, group := range config.Groups {
		haveMatch := false
		for _, selector := range config.Selectors[index].ProxySelector {
			if group.Name == selector.GroupName {
				for _, node := range group.Nodes {
					match, err := regexp.MatchString(selector.Regex, node.Name)
					if err != nil {
						logger.Logger.Warn("Regex error",
							zap.Error(err))
						continue
					}
					if match {
						config.Selectors[index].Proxies = append(config.Selectors[index].Proxies, node)
					}
				}
				haveMatch = true
				break
			}
		}
		if !haveMatch {
			// for _, node := range group.Nodes {
			// 	config.Selectors[index].Proxies = append(config.Selectors[index].Proxies, node)
			// }
		}
	}
	writeToFile()
	logger.Logger.Info("Update selector proxies success")
	return nil
}
