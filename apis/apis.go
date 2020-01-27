package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/logger"
)

var router *gin.Engine

//InitGin init gin router
func InitGin(port int) {
	router := gin.Default()
	router.POST("/api/add", add)
	router.POST("/api/sub", subscribe)
	router.POST("/api/update", update)
	router.Run(":" + strconv.Itoa(port))
}

func add(c *gin.Context) {
	config.LoadConfig()
	class := c.Query("class")
	switch class {
	case "group":
		addGroup(c)
	case "node":
		addNode(c)
	case "rule":
		addRule(c)
	case "selector":
		addSelector(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}

func update(c *gin.Context) {
	config.LoadConfig()
	class := c.Query("class")
	switch class {
	case "all":
		config.UpdateAll()
	case "group":
		updateGroup(c)
	case "rule":
		updateRule(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}

func subscribe(c *gin.Context) {
	config.LoadConfig()
	class := c.Query("class")
	switch class {
	case "clash":
		generateClash(c)
	case "clashr":
		generateClashR(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}
