package config

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"

	"github.com/goomadao/subscribeManager/util/data"
)

func decodeTrojan(bts []byte) (nodes []data.Node, err error) {
	trojanLinks := bytes.Split(bts, []byte("\n"))
	for _, v := range trojanLinks {
		if bytes.Index(bytes.TrimSpace(v), []byte("trojan://")) != 0 {
			continue
		}
		node, err := decodeTrojanLink(v[9:])
		if err == nil {
			nodes = append(nodes, &node)
		}
	}
	return nodes, nil
}

func decodeTrojanLink(bts []byte) (node data.Trojan, err error) {
	var password, server, name, additional string
	var port int64
	port = 443
	sni := ""
	skipCertVerify := false

	// get trojan name
	pos := bytes.LastIndexByte(bts, '#')
	if pos != -1 {
		var err error
		name, err = url.QueryUnescape(string(bts[pos+1:]))
		if err != nil {
			name = "URL decode node's name error"
		}
		bts = bts[:pos]
	}

	// get additional params
	pos = bytes.LastIndexByte(bts, '?')
	if pos != -1 {
		additional = string(bts[pos+1:])
		bts = bts[:pos]
		fmt.Println(additional)
		// Todo: get params
		p, err := url.ParseQuery(additional)
		if err != nil {
			logger.Logger.Warn("Decode trojan param fail",
				zap.Error(err))
			return data.Trojan{}, errors.New("Decode trojan param fail")
		}

		if len(p["sni"]) > 0 {
			sni = p["sni"][0]
		}
		if len(p["allowInsecure"]) > 0 && p["allowInsecure"][0] == "true" {
			skipCertVerify = true
		}
	}

	// get trojan port
	pos = bytes.LastIndexByte(bts, ':')
	if pos != -1 && pos+6 > len(bts) {
		var err error
		port, err = strconv.ParseInt(string(bts[pos+1:]), 10, 32)
		if err != nil {
			logger.Logger.Warn("Decode trojan port fail",
				zap.Error(err))
			return data.Trojan{}, errors.New("Decode trojan port fail")
		}
		bts = bts[:pos]
	}

	// get trojan server
	pos = bytes.LastIndexByte(bts, '@')
	if pos == -1 {
		logger.Logger.Warn("Decode trojan server fail")
		return data.Trojan{}, errors.New("Decode trojan server fail")
	}
	server = string(bts[pos+1:])
	bts = bts[:pos]

	// get trojan pwd
	password = string(bts)

	node = data.Trojan{
		Type:           "trojan",
		Server:         server,
		Name:           name,
		Port:           int(port),
		Password:       password,
		Sni:            sni,
		SkipCertVerify: skipCertVerify,
	}
	return node, nil
}

// Node2Trojan adds Trojan field to Node struct
func Node2Trojan(node *data.RawNode) {
	node.Trojan = data.Trojan{
		Type:           "trojan",
		Server:         node.Server,
		Name:           node.Name,
		Port:           node.Port,
		Password:       node.Password,
		Sni:            node.Sni,
		SkipCertVerify: node.SkipCertVerify,
	}
}
