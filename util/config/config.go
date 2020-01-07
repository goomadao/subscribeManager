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

	"gopkg.in/yaml.v2"
)

var (
	config *data.Config
	//CfgFile - location of config file
	CfgFile  string
	cfgMutex *sync.RWMutex
	client   *http.Client
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
			// node1 := &data.Node{
			// 	Type:     "ss",
			// 	Cipher:   "aes-256-cfb",
			// 	Password: "123456m",
			// 	Name:     "whatever",
			// 	Server:   "127.0.0.1",
			// 	Port:     1024,
			// }
			// node2 := &data.Node{
			// 	Type:     "ss",
			// 	Cipher:   "aes-256-cfb",
			// 	Password: "123456m",
			// 	Name:     "whatever2",
			// 	Server:   "127.0.0.1",
			// 	Port:     10242,
			// }
			// group := data.Group{
			// 	Name: "Default",
			// 	URL:  "",
			// 	// Nodes: []*data.Node{},
			// 	// LastUpdate: time.Now(),
			// }
			config = &data.Config{
				// Groups: []data.Group{group},
			}
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
	for _, group := range config.Groups {
		if len(group.URL) > 0 {
			updateGroup(group.Name)
		}
	}
	for _, selector := range config.Selectors {
		updateSelectorProxies(selector.Name, selector.Type)
	}
	for _, rule := range config.Rules {
		if len(rule.URL) > 0 {
			updateRule(rule.URL)
		}
	}
}
