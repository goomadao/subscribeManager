package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
)

func getRules(c *gin.Context) {
	config.LoadConfig()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   config.GetRules(),
	})
}

func rule(c *gin.Context) {
	config.LoadConfig()
	action := c.Query("action")
	switch action {
	case "add":
		addRule(c)
	case "edit":
		editRule(c)
	case "update":
		updateRule(c)
	case "delete":
		deleteRule(c)
	case "updateall":
		updateAllRules(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}

func addRule(c *gin.Context) {
	var rule data.Rule
	c.BindJSON(&rule)
	fmt.Println(rule)
	rules, err := config.AddRule(rule)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   rules,
	})
}

func updateRule(c *gin.Context) {
	var rule data.Rule
	c.BindJSON(&rule)
	newRule, err := config.UpdateRule(rule.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   newRule,
	})
}

func updateAllRules(c *gin.Context) {
	rules, err := config.UpdateAllRules()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": rules,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": rules,
	})
}

func editRule(c *gin.Context) {
	ruleName := c.Query("rule")
	var rule data.Rule
	c.BindJSON(&rule)
	rules, err := config.EditRule(ruleName, rule)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   rules,
	})
}

func deleteRule(c *gin.Context) {
	rule := c.Query("rule")
	rules, err := config.DeleteRule(rule)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   rules,
	})
}
