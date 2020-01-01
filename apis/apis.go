package apis

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

var Router *gin.Engine

func InitGin(port int) {
	Router := gin.Default()
	Router.GET("/api/add", add)
	Router.Run(":" + strconv.Itoa(port))
}

func add(c *gin.Context) {
	class := c.Query("class")
	switch class {
	case "node":
		addNode(c)
	case "group":
		addGroup(c)
	}
}

func addNode(c *gin.Context) {
	types := c.Query("type")
	cipher := c.Query("cipher")
	password := c.Query("password")
	name := c.Query("name")
	server := c.Query("server")
	port, err := strconv.Atoi(c.Query("port"))
	if err != nil {
		logger.Logger.Error("Parse port to int fail.")
	}
	node := data.Node{
		Type:     types,
		Cipher:   cipher,
		Password: password,
		Name:     name,
		Server:   server,
		Port:     port,
	}
	err = config.AddNode(node)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.String(http.StatusOK, "Add node success.")
}

func addGroup(c *gin.Context) {
	name := c.Query("name")
	groupUrl := c.Query("url")
	if len(name) == 0 {
		temp, err := url.Parse(groupUrl)
		if err != nil {
			logger.Logger.Panic("Parse url fail.",
				zap.Error(err))
		}
		name = temp.Host
	}
	group := data.Group{
		Name: name,
		Url:  groupUrl,
	}
	err := config.AddGroup(group)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.String(http.StatusOK, "Add group success.")
}