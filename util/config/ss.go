package config

import (
	"bytes"
	"encoding/base64"
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

func decodeSS(bts []byte) (nodes []data.Node, err error) {
	ssLinks := bytes.Split(bts, []byte("\n"))
	for _, v := range ssLinks {
		if bytes.Index(v, []byte("ss://")) == -1 {
			continue
		}
		node, err := decodeSSLink(v[5:])
		if err != nil {
			// return nil, err
		} else {
			nodes = append(nodes, &node)
		}
	}
	return nodes, nil
}

func decodeSSLink(bts []byte) (node data.SS, err error) {
	var cipher, password, server, cipherAndPasswordBytes, serverAndPortBytes []byte
	var port int
	if (bytes.IndexByte(bts, ':')) == -1 { //server and port in base64
		pos1 := bytes.IndexAny(bts, "/?#")
		if pos1 == -1 {
			pos1 = len(bts)
		}
		allBytes, err := base64.RawURLEncoding.DecodeString(string(bts[:pos1]))
		if err != nil {
			logger.Logger.Warn("Decode ss cipher, password, server and port fail.",
				zap.Error(err))
			return data.SS{}, err
		}
		all := bytes.Split(allBytes, []byte("@"))
		//get cipher and password
		cipherAndPasswordBytes = all[0]

		//get server and port
		serverAndPortBytes = all[1]
	} else {
		//get cipher and password
		pos1 := bytes.IndexByte(bts, '@')
		cipherAndPasswordBytes, err = base64.RawURLEncoding.DecodeString(string(bts[:pos1]))
		if err != nil {
			logger.Logger.Warn("Decode ss cipher and password fail",
				zap.Error(err))
			return data.SS{}, errors.New("Decode ss cipher and password fail")
		}

		pos2 := bytes.IndexAny(bts, "/?#")
		if pos2 == -1 {
			pos2 = len(bts)
		}
		serverAndPortBytes = bts[pos1+1 : pos2]
	}

	cipherAndPassword := bytes.Split(cipherAndPasswordBytes, []byte(":"))
	cipher = cipherAndPassword[0]
	password = cipherAndPassword[1]

	// //not working for ipv6 address
	// serverAndPort := bytes.Split(serverAndPortBytes, []byte(":"))
	// server = serverAndPort[0]
	// port, err = strconv.Atoi(string(serverAndPort[1]))
	// if err != nil {
	// 	logger.Logger.Warn("Decode ss port fail.",
	// 		zap.Error(err))
	// 	return data.Node{}, err
	// }

	pos3 := bytes.LastIndexByte(serverAndPortBytes, ':')
	server = serverAndPortBytes[:pos3]
	port, err = strconv.Atoi(string(serverAndPortBytes[pos3+1:]))
	if err != nil {
		logger.Logger.Warn("Decode ss port",
			zap.Error(err))
		return data.SS{}, err
	}

	//get plugin
	pos4 := bytes.Index(bts, []byte("?plugin="))
	pos5 := bytes.IndexByte(bts, '#')
	var plugin, pluginStr, pluginOptions string
	var pluginOpts data.Plugin
	if pos4 != -1 {
		if pos5 == -1 {
			pos5 = len(bts)
		}
		pluginStr, err = url.QueryUnescape(string(bts[pos4+8 : pos5]))
		if err != nil {
			logger.Logger.Warn(err.Error())
			return data.SS{}, err
		}
		pluginParams := strings.Split(pluginStr, ";")
		pluginOptions = pluginStr[strings.Index(pluginStr, ";")+1:]
		for i, v := range pluginParams {
			if i == 0 {
				plugin = v
				if plugin == "obfs-local" {
					plugin = "obfs"
				} else if plugin == "simple-obfs" {
					plugin = "obfs"
				}
			} else {
				opts := strings.Split(v, "=")
				if opts[0] == "obfs" {
					pluginOpts.Obfs = opts[1]
				} else if opts[0] == "obfs-host" {
					pluginOpts.ObfsHost = opts[1]
				}
			}
		}
	}

	//get tag
	var tag string
	pos6 := bytes.IndexByte(bts, '#')
	if pos6 != -1 {
		tag, err = url.QueryUnescape(string(bts[pos6+1:]))
		if err != nil {
			logger.Logger.Warn("Decode ss tag fail",
				zap.Error(err))
			return data.SS{}, errors.New("Decode ss tag fail")
		}
	}

	node = data.SS{
		Type:          "ss",
		Server:        string(server),
		Name:          tag,
		Port:          port,
		Cipher:        string(cipher),
		Password:      string(password),
		Plugin:        plugin,
		PluginOptions: pluginOptions,
	}
	SS2Node(&node)

	return node, nil
}

//Node2SS adds SS field to Node strcut
func Node2SS(node *data.RawNode) {
	node.SS = data.SS{
		Type:     "ss",
		Server:   node.Server,
		Name:     node.Name,
		Port:     node.Port,
		Cipher:   node.Cipher,
		Password: node.Password,
	}
	if node.Plugin == "obfs" {
		node.SS.Plugin = "obfs"
		node.SS.PluginOptions = "obfs=" + node.PluginOpts.Obfs +
			";obfs-host=" + node.PluginOpts.ObfsHost
		node.SS.PluginOpts = node.PluginOpts
	}
}

//SS2Node constructs Node struct with SS
func SS2Node(node *data.SS) {
	if node.Plugin == "obfs" {
		pluginOptions := strings.Split(node.PluginOptions, ";")
		for _, val := range pluginOptions {
			option := strings.Split(val, "=")
			if option[0] == "obfs" {
				node.PluginOpts.Obfs = option[1]
			} else if option[0] == "obfs-host" {
				node.PluginOpts.ObfsHost = option[1]
			}
		}
	}
}
