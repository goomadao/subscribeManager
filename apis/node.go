package apis

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
)

func node(c *gin.Context) {
	config.LoadConfig()
	action := c.Query("action")
	switch action {
	case "add":
		addNode(c)
	case "edit":
		edieNode(c)
	case "delete":
		deleteNode(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}

func addNode(c *gin.Context) {
	types := c.Query("type")
	groupName := c.Query("group")
	var node data.RawNode
	var err error
	var group data.Group
	switch types {
	case "ss":
		c.BindJSON(&node.SS)
		config.SS2Node(&node.SS)
		group, err = config.AddNode(groupName, &node.SS)
	case "ssr":
		c.BindJSON(&node.SSR)
		// config.SSR2Node(&node.SSR)
		group, err = config.AddNode(groupName, &node.SSR)
	case "vmess":
		c.BindJSON(&node.Vmess)
		config.Vmess2Node(&node.Vmess)
		group, err = config.AddNode(groupName, &node.Vmess)
	default:
		err = errors.New("No such node type")
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   group,
	})
}

func edieNode(c *gin.Context) {
	nodeName := c.Query("node")
	groupName := c.Query("group")
	types := c.Query("type")
	var node data.RawNode
	var err error
	var group data.Group
	switch types {
	case "ss":
		c.BindJSON(&node.SS)
		config.SS2Node(&node.SS)
		group, err = config.EditNode(nodeName, groupName, &node.SS)
	case "ssr":
		c.BindJSON(&node.SSR)
		// config.SSR2Node(&node.SSR)
		group, err = config.EditNode(nodeName, groupName, &node.SSR)
	case "vmess":
		c.BindJSON(&node.Vmess)
		config.Vmess2Node(&node.Vmess)
		group, err = config.EditNode(nodeName, groupName, &node.Vmess)
	default:
		err = errors.New("No such node type")
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   group,
	})
}

func deleteNode(c *gin.Context) {
	nodeName := c.Query("node")
	groupName := c.Query("group")
	group, err := config.DeleteNode(nodeName, groupName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   group,
	})
}
