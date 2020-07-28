package config

import (
	"encoding/base64"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

func decodeSSD(bts []byte) (nodes []data.Node, err error) {
	ssd, err := base64.RawStdEncoding.DecodeString(string(bts))
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
	for idx := range ssdStruct.Servers {
		ss := ssdStruct.Servers[idx]
		ss.Type = "ss"
		if ss.Port == 0 {
			ss.Port = ssdStruct.Port
		}
		if len(ss.Cipher) == 0 {
			ss.Cipher = ssdStruct.Cipher
		}
		if len(ss.Password) == 0 {
			ss.Password = ssdStruct.Password
		}
		if len(ss.Plugin) == 0 {
			ss.Plugin = ssdStruct.Plugin
			ss.PluginOptions = ssdStruct.PluginOptions
		}
		if ss.Plugin == "simple-obfs" {
			ss.Plugin = "obfs"
		}
		SS2Node(&ss)
		nodes = append(nodes, &ss)
	}
	return nodes, nil
}
