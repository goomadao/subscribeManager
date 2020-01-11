package config

import (
	"encoding/base64"

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
	for _, ss := range ssdStruct.Servers {
		node := data.Node{
			SS: ss,
		}
		if node.SS.Port == 0 {
			node.SS.Port = ssdStruct.Port
		}
		if len(node.SS.Cipher) == 0 {
			node.SS.Cipher = ssdStruct.Cipher
		}
		if len(node.SS.Password) == 0 {
			node.SS.Password = ssdStruct.Password
		}
		if len(node.SS.Plugin) == 0 {
			node.SS.Plugin = ssdStruct.Plugin
			node.SS.PluginOptions = ssdStruct.PluginOptions
		}
		SS2Node(&node)
		nodes = append(nodes, node)
	}
	return nodes, nil
}
