package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
)

func addSelector(c *gin.Context) {
	var selector data.ClashProxyGroupSelector
	c.BindJSON(&selector)
	err := config.AddSelector(selector)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "Add selector success",
	})
}

func updateSelector(c *gin.Context) {
	types := c.Query("type")
	if types == "all" {
		err := config.UpdateAllSelectorProxies()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"msg":    err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
		return
	}
	var selector data.ClashProxyGroupSelector
	c.BindJSON(&selector)
	err := config.UpdateSelectorProxies(selector.Name, selector.Type)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
