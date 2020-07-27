package config

import (
	"errors"
	"regexp"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

//GetSelectors returns all selectors
func GetSelectors() []data.ClashProxyGroupSelector {
	return config.Selectors
}

//AddSelector adds new selector
func AddSelector(selector data.ClashProxyGroupSelector) ([]data.ClashProxyGroupSelector, error) {
	err := selectorDuplicate(selector)
	if err != nil {
		return nil, err
	}
	config.Selectors = append(config.Selectors, selector)
	go UpdateSelectorProxies(selector.Name, selector.Type)
	return config.Selectors, nil
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
func UpdateAllSelectorProxies() ([]data.ClashProxyGroupSelector, error) {
	errorMsg := ""
	for _, selector := range config.Selectors {
		_, err := UpdateSelectorProxies(selector.Name, selector.Type)
		if err != nil {
			errorMsg += err.Error() + "\n"
		}
	}
	if len(errorMsg) > 0 {
		return config.Selectors, errors.New(errorMsg)
	}
	return config.Selectors, nil
}

//UpdateSelectorProxies updates selector's proxies specified by selector name and type
func UpdateSelectorProxies(name string, selectType string) (data.ClashProxyGroupSelector, error) {
	index := -1
	for i, val := range config.Selectors {
		if val.Name == name && val.Type == selectType {
			index = i
			break
		}
	}
	if index == -1 {
		logger.Logger.Warn("No such selector")
		return data.ClashProxyGroupSelector{}, errors.New("No such selector")
	}
	var proxies []data.Node
	for _, group := range config.Groups {
		haveMatch := false
		for _, selector := range config.Selectors[index].ProxySelectors {
			if group.Name == selector.GroupName {
				for _, node := range group.Nodes {
					match, err := regexp.MatchString(selector.Include, node.GetName())
					if err != nil {
						logger.Logger.Warn("Include regex error",
							zap.Error(err))
						continue
					}
					if !match {
						continue
					}
					if len(selector.Exclude) > 0 {
						match, err = regexp.MatchString(selector.Exclude, node.GetName())
						if err != nil {
							logger.Logger.Warn("Exclude regex error",
								zap.Error(err))
							continue
						}
						if match {
							continue
						}
					}
					proxies = append(proxies, node)
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
	return config.Selectors[index], nil
}

//EditSelector replace selector specified by name
func EditSelector(selectorName string, selector data.ClashProxyGroupSelector) ([]data.ClashProxyGroupSelector, error) {
	index := getSelectorIndex(selectorName)
	if index == -1 {
		logger.Logger.Warn("No such selector")
		return nil, errors.New("No such selector")
	}
	config.Selectors[index] = selector
	editProxyGroup(selectorName, selector.Name)
	return config.Selectors, nil
}

//DeleteSelector delete selector specified by name
func DeleteSelector(name string) ([]data.ClashProxyGroupSelector, error) {
	index := getSelectorIndex(name)
	if index == -1 {
		logger.Logger.Warn("No such selector")
		return nil, errors.New("No such selector")
	}
	config.Selectors = append(config.Selectors[:index], config.Selectors[index+1:]...)
	deleteProxyGroup(name)
	return config.Selectors, nil
}

func getSelectorIndex(name string) int {
	for i, val := range config.Selectors {
		if val.Name == name {
			return i
		}
	}
	return -1
}

func editProxyGroup(name string, newName string) {
	if name == newName {
		return
	}
	for i := 0; i < len(config.Selectors); i++ {
		index := getProxyGroupIndex(name, i)
		if index != -1 {
			config.Selectors[i].ProxyGroups[index] = newName
		}
	}
}

func deleteProxyGroup(name string) {
	for i := 0; i < len(config.Selectors); i++ {
		index := getProxyGroupIndex(name, i)
		if index != -1 {
			config.Selectors[i].ProxyGroups = append(config.Selectors[i].ProxyGroups[:index], config.Selectors[i].ProxyGroups[index+1:]...)
		}
	}
}

func getProxyGroupIndex(name string, selectorIndex int) int {
	for i, val := range config.Selectors[selectorIndex].ProxyGroups {
		if val == name {
			return i
		}
	}
	return -1
}
