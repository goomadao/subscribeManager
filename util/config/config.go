package config

import (
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
	config *data.Config
	//CfgFile - location of config file
	CfgFile      string
	cfgMutex     *sync.RWMutex
	client       *http.Client
	jsonIterator jsoniter.API
)

//InitConfig init rwMutex
func InitConfig() {
	cfgMutex = new(sync.RWMutex)
	loadConfig()

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

func writeToFile() error {
	// if len(config.Groups) == 1 {
	// 	return nil
	// }
	// bts, err := yaml.Marshal(config.Groups[1])
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

func loadConfig() {
	cfgMutex.RLock()
	buffer, err := ioutil.ReadFile(CfgFile)
	cfgMutex.RUnlock()
	if err != nil {
		if os.IsNotExist(err) {
			config = &data.Config{}
			err = writeToFile()
			if err != nil {
				logger.Logger.Panic(err.Error())
			}
			logger.Logger.Info("Write to file success")
			return
		}
		logger.Logger.Panic("Read config file fail.",
			zap.Error(err))
		return
	}
	config = &data.Config{}
	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		logger.Logger.Panic("Unmarshal cnofig file fail.",
			zap.Error(err))
	}
	logger.Logger.Info("Unmarshal from config file success.")
}

//UpdateAll updates group, selector and rule
func UpdateAll() {
	UpdateAllGroups()
	UpdateAllSelectorProxies()
	UpdateAllRules()
}
