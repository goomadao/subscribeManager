package config

import "github.com/goomadao/subscribeManager/util/data"

// Node2HTTP constructs Node struct with HTTP
func Node2HTTP(node *data.RawNode) {
	node.HTTP = data.HTTP{
		Type:           "http",
		Server:         node.Server,
		Name:           node.Name,
		Port:           node.Port,
		Username:       node.Username,
		Password:       node.Password,
		ClashTLS:       node.TLS,
		SkipCertVerify: node.SkipCertVerify,
	}
	if node.TLS {
		node.HTTP.TLS = "true"
	} else {
		node.HTTP.TLS = "false"
	}
}
