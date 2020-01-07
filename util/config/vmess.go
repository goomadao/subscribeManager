package config

import (
	"bytes"
	"encoding/base64"
	"encoding/json"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

func decodeVmess(bts []byte) (nodes []data.Node, err error) {
	vmessLinks := bytes.Split(bts, []byte("\n"))
	for _, v := range vmessLinks {
		if bytes.Index(v, []byte("vmess://")) == -1 {
			continue
		}
		node, err := decodeVmessLink(v[8:])
		if err != nil {

		} else {
			nodes = append(nodes, node)
		}
	}
	return nodes, nil
}

func decodeVmessLink(bts []byte) (node data.Node, err error) {
	vmess, err := base64.StdEncoding.DecodeString(string(bts))
	if err != nil {
		logger.Logger.Warn("Decode vmess link fail", zap.Error(err))
		return data.Node{}, err
	}
	var vmessStruct data.Vmess
	if err := json.Unmarshal(vmess, &vmessStruct); err != nil {
		logger.Logger.Warn("Turn vmess json to struct fail",
			zap.Error(err))
		return data.Node{}, err
	}
	node = data.Node{
		Type:     "vmess",
		Cipher:   "auto",
		Password: "",
		Name:     vmessStruct.Name,
		Server:   vmessStruct.Server,
		Port:     vmessStruct.Port,
		UUID:     vmessStruct.UUID,
		AlterID:  vmessStruct.AlterID,
		TLS:      false,
		Network:  vmessStruct.Network,
		WSPath:   vmessStruct.Path,
		WSHeaders: data.WSHeaders{
			Host: vmessStruct.Host,
		},
		Vmess: vmessStruct,
	}
	if vmessStruct.TLS == "tls" {
		node.TLS = true
	}

	return node, nil
}
