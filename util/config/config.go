package config

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"sync"

	"go.uber.org/zap"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"

	"gopkg.in/yaml.v2"
)

var (
	//CfgFile - location of config file
	CfgFile      string
	config       *data.Config
	cfgMutex     *sync.RWMutex
	client       *http.Client
	jsonIterator jsoniter.API
)

//InitConfig init rwMutex
func InitConfig() {
	cfgMutex = new(sync.RWMutex)
	LoadConfig()

	proxy, _ := url.Parse("http://127.0.0.1:7890")
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	client = &http.Client{
		Transport: tr,
	}

	extra.RegisterFuzzyDecoders()
	jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
}

//WriteToFile writes config to file
func WriteToFile() error {
	bts, err := yaml.Marshal(config)
	if err != nil {
		logger.Logger.Warn(err.Error())
		return err
	}
	cfgMutex.Lock()
	err = ioutil.WriteFile(CfgFile, bts, 0644)
	cfgMutex.Unlock()
	if err != nil {
		logger.Logger.Warn(err.Error())
		return err
	}
	return nil
}

//LoadConfig loads config from file
func LoadConfig() {
	cfgMutex.RLock()
	buffer, err := ioutil.ReadFile(CfgFile)
	cfgMutex.RUnlock()
	if err != nil {
		if os.IsNotExist(err) {
			config = &data.Config{}
			return
		}
		logger.Logger.Panic("Read config file fail.",
			zap.Error(err))
		return
	}
	rawConfig := &data.RawConfig{}
	err = yaml.Unmarshal(buffer, rawConfig)
	if err != nil {
		logger.Logger.Panic("Unmarshal cnofig file fail.",
			zap.Error(err))
	}
	var groups []data.Group
	for _, rawGroup := range rawConfig.Groups {
		group := data.Group{
			Name:       rawGroup.Name,
			URL:        rawGroup.URL,
			Nodes:      decodeClashProxies(rawGroup.Nodes),
			LastUpdate: rawGroup.LastUpdate,
		}
		groups = append(groups, group)
	}
	var selectors []data.ClashProxyGroupSelector
	for _, rawSelector := range rawConfig.Selectors {
		selector := data.ClashProxyGroupSelector{
			Name:           rawSelector.Name,
			Type:           rawSelector.Type,
			URL:            rawSelector.URL,
			Interval:       rawSelector.Interval,
			ProxyGroups:    rawSelector.ProxyGroups,
			ProxySelectors: rawSelector.ProxySelectors,
			Proxies:        decodeClashProxies(rawSelector.Proxies),
		}
		selectors = append(selectors, selector)
	}
	config = &data.Config{
		Groups:    groups,
		Selectors: selectors,
		Rules:     rawConfig.Rules,
		Changers:  rawConfig.Changers,
	}
	logger.Logger.Info("Unmarshal from config file success.")
}

//UpdateAll updates group, selector and rule
func UpdateAll() error {
	errMsg := ""
	_, err := UpdateAllGroups()
	if err != nil {
		errMsg += err.Error()
	}
	_, err = UpdateAllSelectorProxies()
	if err != nil {
		errMsg += "\n" + err.Error()
	}
	_, err = UpdateAllRules()
	if err != nil {
		errMsg += "\n" + err.Error()
	}
	if errMsg == "" {
		return nil
	}
	return errors.New(errMsg)
}
