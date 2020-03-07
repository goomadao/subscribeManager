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

func getGroups(c *gin.Context) {
	config.LoadConfig()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   config.GetGroups(),
	})
}

func group(c *gin.Context) {
	config.LoadConfig()
	action := c.Query("action")
	switch action {
	case "add":
		addGroup(c)
	case "edit":
		editGroup(c)
	case "update":
		updateGroup(c)
	case "delete":
		deleteGroup(c)
	case "updateall":
		updateAllGroups(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}

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
				"msg":    "Parse url fail",
			})
			return
		}
		group.Name = temp.Host
	}
	groups, err := config.AddGroup(group)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   groups,
	})
}

func updateGroup(c *gin.Context) {
	var group data.Group
	c.BindJSON(&group)
	fmt.Println(group)
	group, err := config.UpdateGroup(group.Name)
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

func updateAllGroups(c *gin.Context) {
	groups, err := config.UpdateAllGroups()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": config.GetGroups(),
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": groups,
	})
}

func editGroup(c *gin.Context) {
	groupName := c.Query("group")
	var group data.Group
	c.BindJSON(&group)
	groups, err := config.EditGroup(groupName, group)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   groups,
	})
}

func deleteGroup(c *gin.Context) {
	groupName := c.Query("group")
	groups, err := config.DeleteGroup(groupName)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   groups,
	})
}
