package config

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
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
	go UpdateRule(rule.Name)
	return nil
}

func ruleDuplicate(rule data.Rule) error {
	for _, val := range config.Rules {
		if val.Name == rule.Name {
			logger.Logger.Warn("Rule duplicates")
			return errors.New("Rule duplicates")
		}
	}
	return nil
}

//UpdateAllRules updates all rules
func UpdateAllRules() error {
	errorMsg := ""
	for _, rule := range config.Rules {
		if len(rule.URLs) > 0 {
			err := UpdateRule(rule.Name)
			if err != nil {
				errorMsg += err.Error() + "\n"
			}
		}
	}
	if len(errorMsg) > 0 {
		return errors.New(errorMsg)
	}
	return nil
}

//UpdateRule updates rule specified by rule name
func UpdateRule(name string) error {
	loadConfig()
	index := -1
	for i, val := range config.Rules {
		if val.Name == name {
			index = i
			break
		}
	}
	if index == -1 {
		logger.Logger.Warn("No such rule")
		return errors.New("No such rule")
	}
	var tempRules []string
	for _, url := range config.Rules[index].URLs {
		resp, err := client.Get(url)
		if err != nil {
			logger.Logger.Warn("HTTP request for "+url+" fail",
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
			if res, err := addProxyGroupNameAfterRule(string(val), config.Rules[index].Name); err == nil {
				tempRules = append(tempRules, res)
			}
		}
	}
	config.Rules[index].Rules = tempRules
	config.Rules[index].LastUpdate = time.Now()
	writeToFile()
	logger.Logger.Info("Update rule success")
	return nil
}

func addProxyGroupNameAfterRule(rule string, name string) (res string, err error) {
	if strings.Index(rule, "DOMAIN") == 0 ||
		strings.Index(rule, "SOURCE-IP-CIDR") == 0 ||
		strings.Index(rule, "SRC-IP-CIDR") == 0 ||
		strings.Index(rule, "DST-PORT") == 0 ||
		strings.Index(rule, "SRC-PORT") == 0 ||
		strings.Index(rule, "FINAL") == 0 ||
		strings.Index(rule, "MATCH") == 0 {

		var buffer bytes.Buffer
		buffer.WriteString(rule)
		buffer.WriteString("," + name)
		return buffer.String(), nil
	}
	if strings.Index(rule, "IP-CIDR") == 0 || strings.Index(rule, "GEOIP") == 0 {
		var buffer bytes.Buffer
		pos := strings.Index(rule, ",no-resolve")
		if pos != -1 {
			buffer.WriteString(rule[:pos])
			buffer.WriteString("," + name)
			buffer.WriteString(rule[pos:])
		} else {
			buffer.WriteString(rule)
			buffer.WriteString("," + name)
		}
		return buffer.String(), nil
	}
	return "", errors.New("The rule type is not supported by clash")
}
