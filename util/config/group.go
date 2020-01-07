package config

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

//AddGroup adds new group
func AddGroup(group data.Group) error {
	loadConfig()
	err := groupDuplicate(group)
	if err != nil {
		return err
	}
	config.Groups = append(config.Groups, group)
	err = writeToFile()
	if err != nil {
		logger.Logger.Panic("Group write to file fail",
			zap.Error(err))
	}
	logger.Logger.Info("Group write to file success.")
	go updateGroup(group.Name)
	// if err != nil {
	// 	return err
	// }
	return nil
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

func updateGroup(name string) error {
	loadConfig()
	index := -1
	for i, val := range config.Groups {
		if val.Name == name {
			index = i
			break
		}
	}
	if index == -1 {
		logger.Logger.Warn("No such group")
		return errors.New("No such group")
	}
	resp, err := http.Get(config.Groups[index].URL)
	fmt.Println(config.Groups[index].URL)
	if err != nil {
		logger.Logger.Warn("HTTP request for "+config.Groups[index].URL+" fail",
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
	nodes, err := decode(s)
	if err != nil {
		return err
	}
	config.Groups[index].Nodes = nodes
	config.Groups[index].LastUpdate = time.Now()
	writeToFile()
	logger.Logger.Info("Update group success")
	return nil
}

func decode(bts []byte) (nodes []data.Node, err error) {
	//try ss/ssr/vmess
	decodeBytes, err := base64.RawURLEncoding.DecodeString(string(bts))
	if err == nil {
		if strings.Index(string(decodeBytes), "vmess") == 0 {
			nodes, _ = decodeVmess(decodeBytes)
			return nodes, nil
		} else if strings.Index(string(decodeBytes), "ssr") == 0 {
			nodes, _ = decodeSSR(decodeBytes)
			// if err != nil {
			// 	return nil, err
			// }
			return nodes, nil
		} else if strings.Index(string(decodeBytes), "ss") == 0 {
			// fmt.Println("ss")
			// fmt.Println(string(decodeBytes))
			nodes, _ = decodeSS(decodeBytes)
			// if err != nil {
			// 	return nil, err
			// }
			return nodes, nil
		}
	}

	nodes, err = decodeClash(bts)
	if err == nil {
		fmt.Println(nodes)
		return nodes, nil
	}
	return nil, nil
	// return nil, err
}
