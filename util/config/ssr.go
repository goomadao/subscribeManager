package config

import (
	"bytes"
	"encoding/base64"
	"errors"
	"net/url"
	"strconv"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
)

func decodeSSR(bts []byte) (nodes []data.Node, err error) {
	ssrLinks := bytes.Split(bts, []byte("\n"))
	for _, v := range ssrLinks {
		if bytes.Index(v, []byte("ssr://")) == -1 {
			continue
		}
		node, err := decodeSSRLink(v[6:])
		if err != nil {
			// return nil, err
		} else {
			nodes = append(nodes, node)
		}
	}
	return nodes, nil
}

func decodeSSRLink(bts []byte) (node data.Node, err error) {
	ssr, err := base64.RawURLEncoding.DecodeString(string(bts))
	if err != nil {
		logger.Logger.Warn("Decode ssr link fail")
		return data.Node{}, errors.New("Decode ssr link fail")
	}
	pos := bytes.Index(ssr, []byte("/?"))
	//get info before '/?'
	first := bytes.Split(ssr[:pos], []byte(":"))
	// //not working for ipv6 address
	// password, err := base64.RawURLEncoding.DecodeString(string(first[5]))
	// if err != nil {
	// 	logger.Logger.Warn("Decode ssr password fail")
	// 	return data.Node{}, errors.New("Decode ssr password fail")
	// }
	// port, err := strconv.Atoi(string(first[1]))
	// if err != nil {
	// 	logger.Logger.Warn("Decode ssr port fail")
	// 	return data.Node{}, errors.New("Decode ssr port fail")
	// }

	numOfFields := len(first)
	password, err := base64.RawURLEncoding.DecodeString(string(first[numOfFields-1]))
	if err != nil {
		logger.Logger.Warn("Decode ssr password fail")
		return data.Node{}, errors.New("Decode ssr password fail")
	}
	obfs := first[numOfFields-2]
	cipher := first[numOfFields-3]
	protocol := first[numOfFields-4]
	port, err := strconv.Atoi(string(first[numOfFields-5]))
	if err != nil {
		logger.Logger.Warn("Decode ssr port fail")
		return data.Node{}, errors.New("Decode ssr port fail")
	}
	var buffer bytes.Buffer
	buffer.Write(first[numOfFields-5])
	buffer.WriteString(":")
	buffer.Write(protocol)
	pos2 := bytes.Index(ssr, buffer.Bytes()) //position of "port:protocol"
	server := ssr[:pos2-1]

	//get info after '/?'
	tempURL, err := url.Parse("https://get.param" + string(ssr[pos:]))
	if err != nil {
		logger.Logger.Warn("Parse second ssr link to url fail")
		return data.Node{}, errors.New("Parse second ssr link to url fail")
	}
	query := tempURL.Query()
	name, err := base64.RawURLEncoding.DecodeString(query["remarks"][0])
	if err != nil {
		logger.Logger.Warn("Decode ssr name fail")
		return data.Node{}, errors.New("Decode ssr name fail")
	}
	protocolParam, err := base64.RawURLEncoding.DecodeString(query["protoparam"][0])
	if err != nil {
		logger.Logger.Warn("Decode ssr protocol param fail")
		return data.Node{}, errors.New("Decode ssr protocol param fail")
	}
	obfsParam, err := base64.RawURLEncoding.DecodeString(query["obfsparam"][0])
	if err != nil {
		logger.Logger.Warn("Decode ssr obfs param fail")
		return data.Node{}, errors.New("Decode ssr obfs param fail")
	}
	group, err := base64.RawURLEncoding.DecodeString(query["group"][0])
	if err != nil {
		logger.Logger.Warn("Decode ssr group fail")
		return data.Node{}, errors.New("Decode ssr group fail")
	}

	node = data.Node{
		Type:          "ssr",
		Cipher:        string(cipher),
		Password:      string(password),
		Name:          string(name),
		Server:        string(server),
		Port:          port,
		Protocol:      string(protocol),
		ProtocolParam: string(protocolParam),
		Obfs:          string(obfs),
		ObfsParam:     string(obfsParam),
		Group:         string(group),
	}
	node.SSR = data.SSR{
		Server:        node.Server,
		Port:          node.Port,
		Cipher:        node.Cipher,
		Password:      node.Password,
		Name:          node.Name,
		Protocol:      node.Protocol,
		ProtocolParam: node.ProtocolParam,
		Obfs:          node.Obfs,
		ObfsParam:     node.ObfsParam,
		Group:         node.Group,
	}

	return node, nil
}