package config

import (
	"errors"
	"regexp"
	"unicode/utf8"

	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

//GetChangers returns all changers
func GetChangers() []data.NameChanger {
	return config.Changers
}

//AddChanger adds new changer
func AddChanger(changer data.NameChanger) ([]data.NameChanger, error) {
	err := changerDuplicate(changer)
	if err != nil {
		return nil, err
	}
	config.Changers = append(config.Changers, changer)
	// if err != nil {
	// 	return err
	// }
	return config.Changers, nil
}

func changerDuplicate(changer data.NameChanger) error {
	for _, val := range config.Changers {
		if val.Emoji == changer.Emoji {
			logger.Logger.Warn("Changer duplicates")
			return errors.New("Changer duplicates")
		}
	}
	return nil
}

//EditChanger replace changer specified by emoji
func EditChanger(changerEmoji string, changer data.NameChanger) ([]data.NameChanger, error) {
	index := getChangerIndex(changerEmoji)
	if index == -1 {
		logger.Logger.Warn("No such changer")
		return nil, errors.New("No such changer")
	}
	config.Changers[index] = changer
	return config.Changers, nil
}

//DeleteChanger delete changer specified by emojii
func DeleteChanger(emoji string) ([]data.NameChanger, error) {
	index := getChangerIndex(emoji)
	if index == -1 {
		logger.Logger.Warn("No such changer")
		return nil, errors.New("No such changer")
	}
	config.Changers = append(config.Changers[:index], config.Changers[index+1:]...)
	return config.Changers, nil
}

func getChangerIndex(emoji string) int {
	for i, val := range config.Changers {
		if val.Emoji == emoji {
			return i
		}
	}
	return -1
}

//AddEmoji adds emoji before each name
func AddEmoji(node data.Node) {
	if node == nil {
		return
	}
	for _, val := range node.GetName() {
		//in utf8, emoji is 4B
		if _, size := utf8.DecodeRuneInString(string(val)); size > 3 {
			return
		}
	}
	for _, changer := range config.Changers {
		match, err := regexp.MatchString(changer.Regex, node.GetName())
		if err != nil {
			logger.Logger.Warn("Regex error",
				zap.Error(err))
			return
		}
		if match {
			node.SetName(changer.Emoji + " " + node.GetName())
			return
		}
	}
}
