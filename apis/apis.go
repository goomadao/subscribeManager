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

var router *gin.Engine

//InitGin init gin router
func InitGin(port int) {
	router := gin.Default()
	router.GET("/api/add", add)
	router.POST("/api/add", add)
	router.GET("/api/clash", generateClash)
	router.POST("/api/update", update)
	router.Run(":" + strconv.Itoa(port))
}

func add(c *gin.Context) {
	class := c.Query("class")
	switch class {
	case "node":
		addNode(c)
	case "group":
		addGroup(c)
	case "selector":
		addSelector(c)
	case "rule":
		addRule(c)
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
		logger.Logger.Error("Parse port to int fail")
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"error":  "Parse port to int fail",
		})
		return
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

func addGroup(c *gin.Context) {
	name := c.Query("name")
	groupURL := c.Query("url")
	if len(name) == 0 {
		temp, err := url.Parse(groupURL)
		if err != nil {
			logger.Logger.Panic("Parse url fail",
				zap.Error(err))
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
				"error":  "Parse url fail",
			})
			return
		}
		name = temp.Host
	}
	group := data.Group{
		Name: name,
		URL:  groupURL,
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

func addRule(c *gin.Context) {
	var rule data.Rule
	c.BindJSON(&rule)
	// name := c.Query("name")
	// url := c.Query("url")
	// rule := data.Rule{
	// 	Name: name,
	// 	URL:  url,
	// }
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

func generateClash(c *gin.Context) {
	clashFile := config.GenerateClashConfig()
	// c.String()
	c.String(http.StatusOK, string(clashFile))
}

func update(c *gin.Context) {
	class := c.Query("class")
	switch class {
	case "all":
		config.UpdateAll()
	}
}
