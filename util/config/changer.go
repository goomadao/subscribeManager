package config

import (
	"regexp"
	"unicode/utf8"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

//AddEmoji adds emoji before each name
func AddEmoji(node *data.Node) {
	if node == nil {
		return
	}
	for _, val := range node.Name {
		//in utf8, emoji is 4B
		if _, size := utf8.DecodeRuneInString(string(val)); size > 3 {
			return
		}
	}
	for _, changer := range config.Changers {
		match, err := regexp.MatchString(changer.Regex, node.Name)
		if err != nil {
			logger.Logger.Warn("Regex error",
				zap.Error(err))
			return
		}
		if match {
			node.Name = changer.Emoji + node.Name
			if node.Type == "ss" {
				node.SS.Name = node.Name
			} else if node.Type == "ssr" {
				node.SSR.Name = node.Name
			} else if node.Type == "vmess" {
				node.Vmess.Name = node.Name
			}
			return
		}
	}
}
