package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
)

func addNode(c *gin.Context) {
	types := c.Query("type")
	var node data.RawNode
	var err error
	switch types {
	case "ss":
		c.BindJSON(node.SS)
		config.SS2Node(&node.SS)
		err = config.AddNode(&node.SS)
	case "ssr":
		c.BindJSON(node.SSR)
		// config.SSR2Node(&node.SSR)
		err = config.AddNode(&node.SSR)
	case "vmess":
		c.BindJSON(node.Vmess)
		config.Vmess2Node(&node.Vmess)
		err = config.AddNode(&node.Vmess)
	}
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
