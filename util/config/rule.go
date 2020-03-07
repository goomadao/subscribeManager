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

//GetRules returns all rules
func GetRules() []data.Rule {
	return config.Rules
}

//AddRule adds new rule
func AddRule(rule data.Rule) ([]data.Rule, error) {
	err := ruleDuplicate(rule)
	if err != nil {
		return nil, err
	}
	config.Rules = append(config.Rules, rule)
	go UpdateRule(rule.Name)
	return config.Rules, nil
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
func UpdateAllRules() ([]data.Rule, error) {
	errorMsg := ""
	for _, rule := range config.Rules {
		if len(rule.URL) > 0 {
			_, err := UpdateRule(rule.Name)
			if err != nil {
				errorMsg += err.Error() + "\n"
			}
		}
	}
	if len(errorMsg) > 0 {
		return config.Rules, errors.New(errorMsg)
	}
	return config.Rules, nil
}

//UpdateRule updates rule specified by rule name
func UpdateRule(name string) (data.Rule, error) {
	index := getRuleIndex(name)
	if index == -1 {
		logger.Logger.Warn("No such rule")
		return data.Rule{}, errors.New("No such rule")
	}
	var tempRules []string
	url := config.Rules[index].URL
	resp, err := client.Get(url)
	if err != nil {
		logger.Logger.Warn("HTTP request for "+url+" fail",
			zap.Error(err))
		return data.Rule{}, err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Warn("Read from resp fail",
			zap.Error(err))
		return data.Rule{}, err
	}
	rules := bytes.Split(s, []byte("\n"))
	for _, val := range rules {
		if res, err := addProxyGroupNameAfterRule(string(val), config.Rules[index].ProxyGroup); err == nil {
			tempRules = append(tempRules, res)
		}
	}

	config.Rules[index].Rules = tempRules
	config.Rules[index].LastUpdate = time.Now()
	return config.Rules[index], nil
}

//EditRule replaces rule
func EditRule(ruleName string, rule data.Rule) ([]data.Rule, error) {
	index := getRuleIndex(ruleName)
	if index == -1 {
		logger.Logger.Warn("No such rule")
		return nil, errors.New("No such rule")
	}
	config.Rules[index] = rule
	return config.Rules, nil
}

//DeleteRule deletes rule
func DeleteRule(name string) ([]data.Rule, error) {
	index := getRuleIndex(name)
	if index == -1 {
		logger.Logger.Warn("No such rule")
		return nil, errors.New("No such rule")
	}
	config.Rules = append(config.Rules[:index], config.Rules[index+1:]...)
	return config.Rules, nil
}

func getRuleIndex(name string) int {
	for i, val := range config.Rules {
		if val.Name == name {
			return i
		}
	}
	return -1
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
