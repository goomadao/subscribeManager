package config

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

//GetGroups returns all groups
func GetGroups() []data.Group {
	return config.Groups
}

//AddGroup adds new group
func AddGroup(group data.Group) ([]data.Group, error) {
	if len(group.Name) == 0 {
		return nil, errors.New("Group name can't be empty")
	}
	err := groupDuplicate(group)
	if err != nil {
		return nil, err
	}
	config.Groups = append(config.Groups, group)
	go UpdateGroup(group.Name)
	// if err != nil {
	// 	return err
	// }
	return config.Groups, nil
}

func groupDuplicate(group data.Group) error {
	for _, val := range config.Groups {
		if val.Name == group.Name {
			logger.Logger.Warn("Group duplicates")
			return errors.New("Group duplicates")
		}
	}
	return nil
}

//UpdateAllGroups updates all groups
func UpdateAllGroups() ([]data.Group, error) {
	errorMsg := ""
	for _, group := range config.Groups {
		_, err := UpdateGroup(group.Name)
		if err != nil {
			errorMsg += err.Error() + "\n"
		}
	}
	if len(errorMsg) > 0 {
		return config.Groups, errors.New(errorMsg)
	}
	return config.Groups, nil
}

//UpdateGroup updates group specified by group name
func UpdateGroup(name string) (data.Group, error) {
	index := getGroupIndex(name)
	if index == -1 {
		logger.Logger.Warn("No such group")
		return data.Group{}, errors.New("No such group")
	}
	var nodes []data.Node
	if len(config.Groups[index].URL) > 0 {
		logger.Logger.Debug("Requesting " + config.Groups[index].URL)
		resp, err := client.Get(config.Groups[index].URL)
		if err != nil {
			logger.Logger.Warn("HTTP request for "+config.Groups[index].URL+" fail",
				zap.Error(err))
			return data.Group{}, err
		}
		defer resp.Body.Close()
		s, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Logger.Warn("Read from resp fail",
				zap.Error(err))
			return data.Group{}, err
		}
		logger.Logger.Debug("Response: " + string(s))
		nodes, err = decode(s)
		if err != nil {
			return data.Group{}, err
		}
	} else {
		nodes = config.Groups[index].Nodes
	}
	for i := range nodes {
		AddEmoji(nodes[i])
	}
	config.Groups[index].Nodes = nodes
	config.Groups[index].LastUpdate = time.Now()
	return config.Groups[index], nil
}

//EditGroup replace group specified by group name
func EditGroup(groupName string, group data.Group) ([]data.Group, error) {
	index := getGroupIndex(groupName)
	if index == -1 {
		logger.Logger.Warn("No such group")
		return nil, errors.New("No such group")
	}
	group.Nodes = config.Groups[index].Nodes
	config.Groups[index] = group
	return config.Groups, nil
}

//DeleteGroup delete group specified by name
func DeleteGroup(name string) ([]data.Group, error) {
	index := getGroupIndex(name)
	if index == -1 {
		logger.Logger.Warn("No such group")
		return nil, errors.New("No such group")
	}
	config.Groups = append(config.Groups[:index], config.Groups[index+1:]...)
	return config.Groups, nil
}

func getGroupIndex(name string) int {
	for i, val := range config.Groups {
		if val.Name == name {
			return i
		}
	}
	return -1
}

func editProxySelector(groupName string, newGroupName string) {
	if groupName == newGroupName {
		return
	}
	for i := 0; i < len(config.Selectors); i++ {
		index := getProxySelectorIndex(groupName, i)
		if index != -1 {
			config.Selectors[i].ProxySelectors[index].GroupName = newGroupName
		}
	}
}

func deleteProxySelector(groupName string) {
	for i := 0; i < len(config.Selectors); i++ {
		index := getProxySelectorIndex(groupName, i)
		if index != -1 {
			config.Selectors[i].ProxySelectors = append(config.Selectors[i].ProxySelectors[:index], config.Selectors[i].ProxySelectors[index+1:]...)
		}
	}
}

func getProxySelectorIndex(name string, selectorIndex int) int {
	for i, val := range config.Selectors[selectorIndex].ProxySelectors {
		if val.GroupName == name {
			return i
		}
	}
	return -1
}

func decode(bts []byte) (nodes []data.Node, err error) {
	//try ssd
	if bytes.Index(bts, []byte("ssd://")) == 0 {
		nodes, _ := decodeSSD(bts[6:])
		return nodes, nil
	}

	//try ss/ssr/vmess
	//Base64 Decode
	decodeBytes, err := base64.RawURLEncoding.DecodeString(string(bts))
	if err != nil {
		logger.Logger.Debug("RawURLEncoding fail.", zap.Error(err))
		decodeBytes, err = base64.URLEncoding.DecodeString(string(bts))
	}
	if err != nil {
		logger.Logger.Debug("URLEncoding fail.", zap.Error(err))
		decodeBytes, err = base64.RawStdEncoding.DecodeString(string(bts))
	}
	if err != nil {
		logger.Logger.Debug("RawStdEncoding fail.", zap.Error(err))
		decodeBytes, err = base64.StdEncoding.DecodeString(string(bts))
	}
	if err == nil {
		if strings.Index(string(decodeBytes), "vmess") == 0 {
			logger.Logger.Debug("Decoded: " + string(decodeBytes))
			nodes, _ = decodeVmess(decodeBytes)
			return nodes, nil
		} else if strings.Index(string(decodeBytes), "ssr") == 0 {
			logger.Logger.Debug("Decoded: " + string(decodeBytes))
			nodes, _ = decodeSSR(decodeBytes)
			// if err != nil {
			// 	return nil, err
			// }
			return nodes, nil
		} else if strings.Index(string(decodeBytes), "ss") == 0 {
			logger.Logger.Debug("Decoded: " + string(decodeBytes))
			nodes, _ = decodeSS(decodeBytes)
			// if err != nil {
			// 	return nil, err
			// }
			return nodes, nil
		} else if strings.Index(string(decodeBytes), "trojan") == 0 {
			logger.Logger.Debug("Decoded: " + string(decodeBytes))
			nodes, _ = decodeTrojan(decodeBytes)
			return nodes, nil
		}
	}
	logger.Logger.Debug("StdEncoding fail.", zap.Error(err))
	nodes, err = decodeClash(bts)
	if err == nil {
		return nodes, nil
	}
	return nil, nil
}
