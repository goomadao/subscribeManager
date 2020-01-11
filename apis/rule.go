package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
)

func addRule(c *gin.Context) {
	var rule data.Rule
	c.BindJSON(&rule)
	err := config.AddRule(rule)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "Add rule success",
	})
}

func updateRule(c *gin.Context) {
	types := c.Query("type")
	if types == "all" {
		err := config.UpdateAllRules()
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
	var rule data.Rule
	c.BindJSON(&rule)
	err := config.UpdateRule(rule.Name)
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
