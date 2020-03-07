package config

import (
	"errors"
	"time"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
)

//AddNode adds single nodes to group specified by name
func AddNode(groupName string, node data.Node) (data.Group, error) {
	index := getGroupIndex(groupName)
	if index == -1 {
		logger.Logger.Warn("No such group")
		return data.Group{}, errors.New("No such group")
	}
	err := nodeDuplicate(index, node)
	if err != nil {
		return data.Group{}, err
	}
	AddEmoji(node)
	config.Groups[index].Nodes = append(config.Groups[index].Nodes, node)
	config.Groups[index].LastUpdate = time.Now()
	return config.Groups[index], nil
}

//EditNode replaces node
func EditNode(nodeName string, groupName string, node data.Node) (data.Group, error) {
	groupIndex := getGroupIndex(groupName)
	if groupIndex == -1 {
		logger.Logger.Warn("No such group")
		return data.Group{}, errors.New("No such group")
	}
	nodeIndex := getNodeIndex(groupIndex, nodeName)
	if nodeIndex == -1 {
		logger.Logger.Warn("No such node")
		return data.Group{}, errors.New("No such node")
	}
	config.Groups[groupIndex].Nodes[nodeIndex] = node
	return config.Groups[groupIndex], nil
}

//DeleteNode deltes node
func DeleteNode(node string, group string) (data.Group, error) {
	groupIndex := getGroupIndex(group)
	if groupIndex == -1 {
		logger.Logger.Warn("No such group")
		return data.Group{}, errors.New("No such group")
	}
	nodeIndex := getNodeIndex(groupIndex, node)
	if nodeIndex == -1 {
		logger.Logger.Warn("No such node")
		return data.Group{}, errors.New("No such node")
	}
	config.Groups[groupIndex].Nodes = append(config.Groups[groupIndex].Nodes[:nodeIndex], config.Groups[groupIndex].Nodes[nodeIndex+1:]...)
	return config.Groups[groupIndex], nil
}

func nodeDuplicate(index int, node data.Node) error {
	for _, val := range config.Groups[index].Nodes {
		if val.GetName() == node.GetName() {
			logger.Logger.Warn("Node duplicates")
			return errors.New("Node duplicates")
		}
	}
	return nil
}

func getNodeIndex(group int, name string) int {
	for i, val := range config.Groups[group].Nodes {
		if val.GetName() == name {
			return i
		}
	}
	return -1
}
