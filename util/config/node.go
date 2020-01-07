package config

import (
	"errors"
	"time"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

//AddNode adds single nodes to default group
func AddNode(node data.Node) error {
	loadConfig()
	err := nodeDuplicate(node)
	if err != nil {
		return err
	}
	config.Groups[0].Nodes = append(config.Groups[0].Nodes, node)
	config.Groups[0].LastUpdate = time.Now()
	err = writeToFile()
	if err != nil {
		logger.Logger.Panic("Node write to file fail",
			zap.Error(err))
	}
	logger.Logger.Info("Node write to file success.")
	return nil
}

func nodeDuplicate(node data.Node) error {
	index := -1
	for i, group := range config.Groups {
		if group.Name == "Default" {
			index = i
		}
	}
	if index == -1 {
		index = len(config.Groups)
		config.Groups = append(config.Groups, data.Group{
			Name: "Default",
		})
	}
	writeToFile()
	for _, val := range config.Groups[0].Nodes {
		if val.Server == node.Server && val.Port == node.Port {
			logger.Logger.Warn("Node duplicates")
			return errors.New("Node duplicates")
		}
	}
	return nil
}
