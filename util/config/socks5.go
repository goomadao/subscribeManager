package config

import "github.com/goomadao/subscribeManager/util/data"

// Node2Socks5 constructs Node struct with Socks5
func Node2Socks5(node *data.RawNode) {
	node.Socks5 = data.Socks5{
		Type:           "socks5",
		Server:         node.Server,
		Name:           node.Name,
		Port:           node.Port,
		Username:       node.Username,
		Password:       node.Password,
		ClashTLS:       node.TLS,
		SkipCertVerify: node.SkipCertVerify,
	}
	if node.TLS {
		node.Socks5.TLS = "true"
	} else {
		node.Socks5.TLS = "false"
	}
}
