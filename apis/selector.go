package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
)

func getSelectors(c *gin.Context) {
	config.LoadConfig()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   config.GetSelectors(),
	})
}

func selector(c *gin.Context) {
	config.LoadConfig()
	action := c.Query("action")
	switch action {
	case "add":
		addSelector(c)
	case "edit":
		editSelector(c)
	case "update":
		updateSelector(c)
	case "delete":
		deleteSelector(c)
	case "updateall":
		updateAllSelectors(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}

func addSelector(c *gin.Context) {
	var selector data.ClashProxyGroupSelector
	c.BindJSON(&selector)
	selectors, err := config.AddSelector(selector)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   selectors,
	})
}

func updateSelector(c *gin.Context) {
	var selector data.ClashProxyGroupSelector
	c.BindJSON(&selector)
	newSelector, err := config.UpdateSelectorProxies(selector.Name, selector.Type)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newSelector,
	})
}

func updateAllSelectors(c *gin.Context) {
	selectors, err := config.UpdateAllSelectorProxies()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": selectors,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": selectors,
	})
}

func editSelector(c *gin.Context) {
	selectorName := c.Query("selector")
	var selector data.ClashProxyGroupSelector
	c.BindJSON(&selector)
	selectors, err := config.EditSelector(selectorName, selector)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   selectors,
	})
}

func deleteSelector(c *gin.Context) {
	selector := c.Query("selector")
	selectors, err := config.DeleteSelector(selector)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   selectors,
	})
}
