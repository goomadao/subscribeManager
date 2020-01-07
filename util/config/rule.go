package config

import (
	"bytes"
	"errors"
	"io/ioutil"
	"time"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

//AddRule adds new rule
func AddRule(rule data.Rule) error {
	loadConfig()
	err := ruleDuplicate(rule)
	if err != nil {
		return err
	}
	config.Rules = append(config.Rules, rule)
	err = writeToFile()
	if err != nil {
		logger.Logger.Panic("Rule write to file fail",
			zap.Error(err))
	}
	logger.Logger.Info("Rule write to file success")
	go updateRule(rule.URL)
	return nil
}

func ruleDuplicate(rule data.Rule) error {
	for _, val := range config.Rules {
		if val.URL == rule.URL {
			logger.Logger.Warn("Rule duplicates")
			return errors.New("Rule duplicates")
		}
	}
	return nil
}

func updateRule(url string) error {
	loadConfig()
	index := -1
	for i, val := range config.Rules {
		if val.URL == url {
			index = i
			break
		}
	}
	if index == -1 {
		logger.Logger.Warn("No such rule")
		return errors.New("No such rule")
	}
	resp, err := client.Get(config.Rules[index].URL)
	if err != nil {
		logger.Logger.Warn("HTTP request for "+config.Rules[index].URL+" fail",
			zap.Error(err))
		return err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Warn("Read from resp fail",
			zap.Error(err))
		return err
	}
	rules := bytes.Split(s, []byte("\n"))
	for _, val := range rules {
		if bytes.Index(val, []byte("DOMAIN")) == 0 ||
			bytes.Index(val, []byte("SOURCE-IP-CIDR")) == 0 ||
			bytes.Index(val, []byte("SRC-IP-CIDR")) == 0 ||
			bytes.Index(val, []byte("DST-PORT")) == 0 ||
			bytes.Index(val, []byte("SRC-PORT")) == 0 ||
			bytes.Index(val, []byte("FINAL")) == 0 ||
			bytes.Index(val, []byte("MATCH")) == 0 {

			var buffer bytes.Buffer
			buffer.Write(val)
			buffer.WriteString("," + config.Rules[index].Name)
			config.Rules[index].Rules = append(config.Rules[index].Rules, buffer.String())
		}
		if bytes.Index(val, []byte("IP-CIDR")) == 0 || bytes.Index(val, []byte("GEOIP")) == 0 {
			var buffer bytes.Buffer
			pos := bytes.Index(val, []byte(",no-resolve"))
			if pos != -1 {
				buffer.Write(val[:pos])
				buffer.WriteString("," + config.Rules[index].Name)
				buffer.Write(val[pos:])
			} else {
				buffer.Write(val)
				buffer.WriteString("," + config.Rules[index].Name)
			}
			config.Rules[index].Rules = append(config.Rules[index].Rules, buffer.String())
		}
	}
	config.Rules[index].LastUpdate = time.Now()
	writeToFile()
	logger.Logger.Info("Update rule success")
	return nil
}
