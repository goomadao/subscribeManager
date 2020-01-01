package config

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"

	"gopkg.in/yaml.v2"
)

var (
	config   *data.Config
	CfgFile  string
	cfgMutex *sync.RWMutex
)

func InitConfig() {
	cfgMutex = new(sync.RWMutex)
	loadConfig()
}

func writeToFile() error {
	if len(config.Groups) == 1 {
		return nil
	}
	bts, err := yaml.Marshal(config.Groups[1])
	// bts, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	cfgMutex.Lock()
	err = ioutil.WriteFile(CfgFile, bts, 0644)
	cfgMutex.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func loadConfig() {
	cfgMutex.RLock()
	buffer, err := ioutil.ReadFile(CfgFile)
	cfgMutex.RUnlock()
	if err != nil {
		if os.IsNotExist(err) {
			// node1 := &data.Node{
			// 	Type:     "ss",
			// 	Cipher:   "aes-256-cfb",
			// 	Password: "123456m",
			// 	Name:     "whatever",
			// 	Server:   "127.0.0.1",
			// 	Port:     1024,
			// }
			// node2 := &data.Node{
			// 	Type:     "ss",
			// 	Cipher:   "aes-256-cfb",
			// 	Password: "123456m",
			// 	Name:     "whatever2",
			// 	Server:   "127.0.0.1",
			// 	Port:     10242,
			// }
			group := data.Group{
				Name: "Default",
				Url:  "",
				// Nodes: []*data.Node{},
				// LastUpdate: time.Now(),
			}
			config = &data.Config{
				Groups: []data.Group{group},
			}
			err = writeToFile()
			if err != nil {
				logger.Logger.Panic(err.Error())
			}
			logger.Logger.Info("Write to file success")
			return
		} else {
			logger.Logger.Panic("Read config file fail.",
				zap.Error(err))
			return
		}
	}
	config = &data.Config{}
	err = yaml.Unmarshal(buffer, config)
	if err != nil {
		logger.Logger.Panic("Unmarshal cnofig file fail.",
			zap.Error(err))
	}
	logger.Logger.Info("Unmarshal from config file success.")
}

//Add single nodes to default group
func AddNode(node data.Node) error {
	loadConfig()
	err := nodeDuplicate(node)
	if err != nil {
		return err
	}
	config.Groups[0].Nodes = append(config.Groups[0].Nodes, node)
	config.Groups[0].LastUpdate = time.Now()
	err = writeToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success.")
	return nil
}

func nodeDuplicate(node data.Node) error {
	for _, val := range config.Groups[0].Nodes {
		if val.Name == node.Name {
			return errors.New("Node duplicates.")
		}
	}
	return nil
}

func AddGroup(group data.Group) error {
	loadConfig()
	err := groupDuplicate(group)
	if err != nil {
		return err
	}
	config.Groups = append(config.Groups, group)
	err = writeToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success.")
	go updateGroup(group.Name)
	if err != nil {
		return err
	}
	return nil
}

func groupDuplicate(group data.Group) error {
	for _, val := range config.Groups {
		if val.Name == group.Name {
			return errors.New("Group duplicates.")
		}
	}
	return nil
}

func updateGroup(name string) error {
	index := -1
	for i, val := range config.Groups {
		if val.Name == name {
			index = i
			break
		}
	}
	if index == -1 {
		return errors.New("No such group.")
	}
	resp, err := http.Get(config.Groups[index].Url)
	fmt.Println(config.Groups[index].Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	nodes, err := decode(s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	config.Groups[index].Nodes = nodes
	config.Groups[index].LastUpdate = time.Now()
	writeToFile()
	logger.Logger.Info("Update group success.")
	return nil
}

func decode(bts []byte) (nodes []data.Node, err error) {
	//try ss/ssr/vmess
	decodeBytes, err := base64.RawURLEncoding.DecodeString(string(bts))
	if err == nil {
		if strings.Index(string(decodeBytes), "vmess") == 0 {
			// fmt.Println("vmess")
			// fmt.Println(string(decodeBytes))
		} else if strings.Index(string(decodeBytes), "ssr") == 0 {
			nodes, err = decodeSSR(decodeBytes)
			if err != nil {
				return nil, errors.New("Decode ssr subscribe fail.")
			}
			return nodes, nil
		} else if strings.Index(string(decodeBytes), "ss") == 0 {
			// fmt.Println("ss")
			// fmt.Println(string(decodeBytes))
			nodes, err = decodeSS(decodeBytes)
			if err != nil {
				return nil, err
			}
			return nodes, nil
		}
	}

	nodes, err = decodeClash(bts)
	if err == nil {
		fmt.Println(nodes)
		return nodes, nil
	}
	return nil, err
}

func decodeClash(bts []byte) (nodes []data.Node, err error) {
	fmt.Println("aslkdjkasjdkajkdjsakljdklsj", string(bts))
	clash := data.Clash{}
	err = yaml.Unmarshal(bts, &clash)
	if err != nil {
		logger.Logger.Warn(err.Error())
		return nil, err
	}
	return clash.Proxy, nil
}

func decodeSSR(bts []byte) (nodes []data.Node, err error) {
	ssrLinks := bytes.Split(bts, []byte("\n"))
	for _, v := range ssrLinks {
		if bytes.Index(v, []byte("ssr://")) == -1 {
			continue
		}
		node, err := decodeSSRLink(v[6:])
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func decodeSSRLink(bts []byte) (node data.Node, err error) {
	fmt.Println(string(bts))
	ssr, err := base64.RawURLEncoding.DecodeString(string(bts))
	if err != nil {
		return data.Node{}, errors.New("Decode ssr link fail.")
	}
	pos := bytes.Index(ssr, []byte("/?"))
	//get info before '/?'
	first := bytes.Split(ssr[:pos], []byte(":"))
	password, err := base64.RawURLEncoding.DecodeString(string(first[5]))
	if err != nil {
		return data.Node{}, errors.New("Decode ssr password fail.")
	}
	port, err := strconv.Atoi(string(first[1]))
	if err != nil {
		return data.Node{}, errors.New("Decode ssr port fail.")
	}
	//get info after '/?'
	tempUrl, err := url.Parse("https://get.param" + string(ssr[pos:]))
	if err != nil {
		return data.Node{}, errors.New("Parse second ssr link to url fail.")
	}
	query := tempUrl.Query()
	name, err := base64.RawURLEncoding.DecodeString(query["remarks"][0])
	if err != nil {
		return data.Node{}, errors.New("Decode ssr name fail.")
	}
	protocolParam, err := base64.RawURLEncoding.DecodeString(query["protoparam"][0])
	if err != nil {
		return data.Node{}, errors.New("Decode ssr protocol param fail.")
	}
	obfsParam, err := base64.RawURLEncoding.DecodeString(query["obfsparam"][0])
	if err != nil {
		return data.Node{}, errors.New("Decode ssr obfs param fail.")
	}
	group, err := base64.RawURLEncoding.DecodeString(query["group"][0])
	if err != nil {
		return data.Node{}, errors.New("Decode ssr group fail.")
	}

	node = data.Node{
		Type:          "ssr",
		Cipher:        string(first[3]),
		Password:      string(password),
		Name:          string(name),
		Server:        string(first[0]),
		Port:          port,
		Protocol:      string(first[2]),
		ProtocolParam: string(protocolParam),
		Obfs:          string(first[4]),
		ObfsParam:     string(obfsParam),
		Group:         string(group),
	}

	return node, nil
}

func decodeSS(bts []byte) (nodes []data.Node, err error) {
	ssLinks := bytes.Split(bts, []byte("\n"))
	for _, v := range ssLinks {
		if bytes.Index(v, []byte("ss://")) == -1 {
			continue
		}
		node, err := decodeSSLink(v[5:])
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func decodeSSLink(bts []byte) (node data.Node, err error) {
	var cipher, password, server, cipherAndPasswordBytes, serverAndPortBytes []byte
	var port int
	if (bytes.IndexByte(bts, ':')) == -1 { //server and port in base64
		pos1 := bytes.IndexAny(bts, "/?#")
		if pos1 == -1 {
			pos1 = len(bts)
		}
		allBytes, err := base64.RawURLEncoding.DecodeString(string(bts[:pos1]))
		if err != nil {
			logger.Logger.Warn("Decode ss cipher, password, server and port fail.")
			return data.Node{}, err
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
			return data.Node{}, errors.New("Decode ss cipher and password fail.")
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

	serverAndPort := bytes.Split(serverAndPortBytes, []byte(":"))
	server = serverAndPort[0]
	port, err = strconv.Atoi(string(serverAndPort[1]))
	if err != nil {
		logger.Logger.Warn("Decode ss port fail.",
			zap.Error(err))
		return data.Node{}, err
	}

	//get plugin
	pos4 := bytes.Index(bts, []byte("?plugin="))
	pos5 := bytes.IndexByte(bts, '#')
	var plugin string
	var pluginOpts data.Plugin
	if pos4 != -1 {
		if pos5 == -1 {
			pos5 = len(bts)
		}
		pluginStr, err := url.QueryUnescape(string(bts[pos4+8 : pos5]))
		if err != nil {
			return data.Node{}, err
		}
		pluginParams := strings.Split(pluginStr, ";")
		for i, v := range pluginParams {
			if i == 0 {
				plugin = v
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
			return data.Node{}, errors.New("Decode ss tag fail.")
		}
	}

	node = data.Node{
		Type:       "ss",
		Cipher:     string(cipher),
		Password:   string(password),
		Name:       tag,
		Server:     string(server),
		Port:       port,
		Plugin:     plugin,
		PluginOpts: pluginOpts,
	}
	fmt.Println(node)
	return node, nil
}