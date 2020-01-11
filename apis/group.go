package apis

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
	"go.uber.org/zap"
)

func addGroup(c *gin.Context) {
	var group data.Group
	c.BindJSON(&group)
	if len(group.Name) == 0 {
		temp, err := url.Parse(group.URL)
		if err != nil {
			logger.Logger.Panic("Parse url fail",
				zap.Error(err))
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"error":  "Parse url fail",
			})
			return
		}
		group.Name = temp.Host
	}
	err := config.AddGroup(group)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"msg":    "Add group success",
	})
}

func updateGroup(c *gin.Context) {
	types := c.Query("type")
	if types == "all" {
		err := config.UpdateAllGroups()
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
	var group data.Group
	c.BindJSON(&group)
	fmt.Println(group)
	err := config.UpdateGroup(group.Name)
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
