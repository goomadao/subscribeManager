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
	if err := jsonIterator.Unmarshal(vmess, &vmessStruct); err != nil {
		logger.Logger.Warn("Turn vmess json to struct fail",
			zap.Error(err))
		return data.Node{}, err
	}
	node = data.Node{Vmess: vmessStruct}
	Vmess2Node(&node)

	return node, nil
}

//Node2Vmess adds Vmess field to Node strcut
func Node2Vmess(node *data.Node) {
	node.Vmess = data.Vmess{
		Host:    node.WSHeaders.Host,
		Path:    node.WSPath,
		Server:  node.Server,
		Port:    node.Port,
		AlterID: node.AlterID,
		Network: node.Network,
		Type:    "none",
		V:       "2",
		Name:    node.Name,
		UUID:    node.UUID,
		Class:   1,
	}
	if node.TLS {
		node.Vmess.TLS = "tls"
	}
}

//Vmess2Node constructs Node struct with Vmess
func Vmess2Node(node *data.Node) {
	*node = data.Node{
		Type:     "vmess",
		Cipher:   "auto",
		Password: "",
		Name:     node.Vmess.Name,
		Server:   node.Vmess.Server,
		Port:     node.Vmess.Port,
		UUID:     node.Vmess.UUID,
		AlterID:  node.Vmess.AlterID,
		TLS:      false,
		Network:  node.Vmess.Network,
		WSPath:   node.Vmess.Path,
		WSHeaders: data.WSHeaders{
			Host: node.Vmess.Host,
		},
		Vmess: node.Vmess,
	}
	if node.Vmess.TLS == "tls" {
		node.TLS = true
	}
}
