package data

func streamCipherSupported(cipher string) bool {
	switch cipher {
	case "rc4-md5":
		fallthrough
	case "aes-128-ctr":
		fallthrough
	case "aes-192-ctr":
		fallthrough
	case "aes-256-ctr":
		fallthrough
	case "aes-128-cfb":
		fallthrough
	case "aes-192-cfb":
		fallthrough
	case "aes-256-cfb":
		fallthrough
	case "chacha20-ietf":
		fallthrough
	case "xchacha20":
		return true
	}
	return false
}

func aeadCipherSupported(cipher string) bool {
	switch cipher {
	case "aes-128-gcm":
		fallthrough
	case "aes-192-gcm":
		fallthrough
	case "aes-256-gcm":
		fallthrough
	case "chacha20-ietf-poly1305":
		fallthrough
	case "xchacha20-ietf-poly1305":
		return true
	}
	return false
}

func ssCipherSupported(cipher string) bool {
	if cipher == "none" || streamCipherSupported(cipher) || aeadCipherSupported(cipher) {
		return true
	}
	return false
}

func ssPluginSupported(plugin string) bool {
	if plugin == "obfs" {
		return true
	}
	return false
}

func ssrObfsSupported(obfs string) bool {
	switch obfs {
	case "plain":
		fallthrough
	case "http_simple":
		fallthrough
	case "http_post":
		fallthrough
	case "random_head":
		fallthrough
	case "tls1.2_ticket_auth":
		fallthrough
	case "tls1.2_ticket_fastauth":
		return true
	}
	return false
}

func ssrProtocolSupported(protocol string) bool {
	switch protocol {
	case "origin":
		fallthrough
	case "auth_sha1_v4":
		fallthrough
	case "auth_aes128_md5":
		fallthrough
	case "auth_aes128_sha1":
		fallthrough
	case "auth_chain_a":
		fallthrough
	case "auth_chain_b":
		return true
	}
	return false
}
