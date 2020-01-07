package config

import (
	"encoding/base64"
	"strings"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

func decodeSSD(bts []byte) (nodes []data.Node, err error) {
	ssd, err := base64.RawURLEncoding.DecodeString(string(bts))
	if err != nil {
		logger.Logger.Warn("Decode ssd subscribe link fail",
			zap.Error(err))
		return nil, err
	}
	var ssdStruct data.SSD
	err = jsonIterator.Unmarshal(ssd, &ssdStruct)
	if err != nil {
		logger.Logger.Warn("Unmarshal ssd struct fail",
			zap.Error(err))
		return nil, err
	}
	var plugin string
	var pluginOpts data.Plugin
	if ssdStruct.Plugin == "simple-obfs" {
		plugin = "obfs"
		pluginOptions := strings.Split(ssdStruct.PluginOptions, ";")
		for _, val := range pluginOptions {
			option := strings.Split(val, "=")
			if option[0] == "obfs" {
				pluginOpts.Obfs = option[1]
			} else if option[0] == "obfs-host" {
				pluginOpts.ObfsHost = option[1]
			}
		}
	}
	for _, ss := range ssdStruct.Servers {
		node := data.Node{
			Type:     "ss",
			Cipher:   ss.Cipher,
			Password: ss.Password,
			Name:     ss.Name,
			Server:   ss.Server,
			Port:     ss.Port,
			SS:       ss,
		}
		if ss.Plugin == "simple-obfs" {
			node.Plugin = "obfs"
			pluginOptions := strings.Split(ss.PluginOptions, ";")
			for _, val := range pluginOptions {
				option := strings.Split(val, "=")
				if option[0] == "obfs" {
					node.PluginOpts.Obfs = option[1]
				} else if option[0] == "obfs-host" {
					node.PluginOpts.ObfsHost = option[1]
				}
			}
		}
		if node.Port == 0 {
			node.Port = ssdStruct.Port
		}
		if len(node.Cipher) == 0 {
			node.Cipher = ssdStruct.Cipher
		}
		if len(node.Password) == 0 {
			node.Password = ssdStruct.Password
		}
		if len(node.Plugin) == 0 {
			node.Plugin = plugin
			node.PluginOpts = pluginOpts
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}
