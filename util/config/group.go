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

//AddGroup adds new group
func AddGroup(group data.Group) error {
	err := groupDuplicate(group)
	if err != nil {
		return err
	}
	config.Groups = append(config.Groups, group)
	go UpdateGroup(group.Name)
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

//UpdateAllGroups updates all groups
func UpdateAllGroups() error {
	errorMsg := ""
	for _, group := range config.Groups {
		err := UpdateGroup(group.Name)
		if err != nil {
			errorMsg += err.Error() + "\n"
		}
	}
	if len(errorMsg) > 0 {
		return errors.New(errorMsg)
	}
	return nil
}

//UpdateGroup updates group specified by group name
func UpdateGroup(name string) error {
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
	var nodes []data.Node
	if len(config.Groups[index].URL) > 0 {
		resp, err := client.Get(config.Groups[index].URL)
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
		nodes, err = decode(s)
		if err != nil {
			return err
		}
	} else {
		nodes = config.Groups[index].Nodes
	}
	for i := range nodes {
		AddEmoji(nodes[i])
	}
	config.Groups[index].Nodes = nodes
	config.Groups[index].LastUpdate = time.Now()
	return nil
}

func decode(bts []byte) (nodes []data.Node, err error) {
	//try ssd
	if bytes.Index(bts, []byte("ssd://")) == 0 {
		nodes, _ := decodeSSD(bts[6:])
		return nodes, nil
	}

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
			nodes, _ = decodeSS(decodeBytes)
			// if err != nil {
			// 	return nil, err
			// }
			return nodes, nil
		}
	}

	nodes, err = decodeClash(bts)
	if err == nil {
		return nodes, nil
	}
	return nil, nil
}
