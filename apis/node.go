package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
)

func addNode(c *gin.Context) {
	types := c.Query("type")
	var node data.Node
	switch types {
	case "ss":
		c.BindJSON(node.SS)
		config.SS2Node(&node)
	case "ssr":
		c.BindJSON(node.SSR)
		config.SSR2Node(&node)
	case "vmess":
		c.BindJSON(node.Vmess)
		config.Vmess2Node(&node)
	}
	err := config.AddNode(node)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "fail",
		"msg":    "Add node success",
	})
}
