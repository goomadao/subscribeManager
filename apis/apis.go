package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
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
}

func update(c *gin.Context) {
	class := c.Query("class")
	switch class {
	case "all":
		config.UpdateAll()
	case "group":
		updateGroup(c)
	case "rule":
		updateRule(c)
	}
}

func subscribe(c *gin.Context) {
	class := c.Query("class")
	switch class {
	case "clash":
		generateClash(c)
	}
}
