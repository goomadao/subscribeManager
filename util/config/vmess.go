package config

import (
	"bytes"
	"encoding/base64"

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
			nodes = append(nodes, &node)
		}
	}
	return nodes, nil
}

func decodeVmessLink(bts []byte) (node data.Vmess, err error) {
	vmess, err := base64.StdEncoding.DecodeString(string(bts))
	if err != nil {
		logger.Logger.Warn("Decode vmess link fail", zap.Error(err))
		return data.Vmess{}, err
	}
	var vmessStruct data.Vmess
	if err := jsonIterator.Unmarshal(vmess, &vmessStruct); err != nil {
		logger.Logger.Warn("Turn vmess json to struct fail",
			zap.Error(err))
		return data.Vmess{}, err
	}
	node = vmessStruct
	Vmess2Node(&node)

	return node, nil
}

//Node2Vmess adds Vmess field to Node strcut
func Node2Vmess(node *data.RawNode) {
	node.Vmess = data.Vmess{
		ClashType:    "vmess",
		WSHeaders:    node.WSHeaders,
		ClashTLS:     node.TLS,
		ClashNetwork: node.Network,
		Cipher:       node.Cipher,
		Host:         node.WSHeaders.Host,
		Path:         node.WSPath,
		Server:       node.Server,
		Port:         node.Port,
		AlterID:      node.AlterID,
		Network:      node.Network,
		Type:         "none",
		V:            "2",
		Name:         node.Name,
		UUID:         node.UUID,
		Class:        1,
	}
	if node.TLS {
		node.Vmess.TLS = "tls"
	}
	if len(node.Network) == 0 {
		node.Vmess.Network = "tcp"
	}
}

//Vmess2Node constructs Node struct with Vmess
func Vmess2Node(node *data.Vmess) {
	node.ClashType = "vmess"
	node.WSHeaders.Host = node.Host
	if node.TLS == "tls" {
		node.ClashTLS = true
	}
	if node.Network == "ws" {
		node.ClashNetwork = "ws"
	}
	node.Cipher = "auto"
}
