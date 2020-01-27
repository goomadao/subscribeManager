package config

import (
	"errors"
	"regexp"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

//AddSelector adds new selector
func AddSelector(selector data.ClashProxyGroupSelector) error {
	err := selectorDuplicate(selector)
	if err != nil {
		return err
	}
	config.Selectors = append(config.Selectors, selector)
	go UpdateSelectorProxies(selector.Name, selector.Type)
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

//UpdateAllSelectorProxies updates all selectors' proxies
func UpdateAllSelectorProxies() error {
	errorMsg := ""
	for _, selector := range config.Selectors {
		err := UpdateSelectorProxies(selector.Name, selector.Type)
		if err != nil {
			errorMsg += err.Error() + "\n"
		}
	}
	if len(errorMsg) > 0 {
		return errors.New(errorMsg)
	}
	return nil
}

//UpdateSelectorProxies updates selector's proxies specified by selector name and type
func UpdateSelectorProxies(name string, selectType string) error {
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
	var proxies []data.Node
	for _, group := range config.Groups {
		haveMatch := false
		for _, selector := range config.Selectors[index].ProxySelector {
			if group.Name == selector.GroupName {
				for _, node := range group.Nodes {
					match, err := regexp.MatchString(selector.Regex, node.GetName())
					if err != nil {
						logger.Logger.Warn("Regex error",
							zap.Error(err))
						continue
					}
					if match {
						proxies = append(proxies, node)
					}
				}
				haveMatch = true
				break
			}
		}
		if !haveMatch {
			// for _, node := range group.Nodes {
			// 	proxies = append(proxies, node)
			// }
		}
	}
	config.Selectors[index].Proxies = proxies
	return nil
}
