package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/data"
	"github.com/goomadao/subscribeManager/util/logger"
)

func getChangers(c *gin.Context) {
	config.LoadConfig()
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   config.GetChangers(),
	})
}

func changer(c *gin.Context) {
	config.LoadConfig()
	action := c.Query("action")
	switch action {
	case "add":
		addChanger(c)
	case "edit":
		editChanger(c)
	case "delete":
		deleteChanger(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}

func addChanger(c *gin.Context) {
	var changer data.NameChanger
	c.BindJSON(&changer)
	changers, err := config.AddChanger(changer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   changers,
	})
}

func editChanger(c *gin.Context) {
	changerEmoji := c.Query("changer")
	var changer data.NameChanger
	c.BindJSON(&changer)
	changers, err := config.EditChanger(changerEmoji, changer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   changers,
	})
}

func deleteChanger(c *gin.Context) {
	changerEmoji := c.Query("changer")
	changers, err := config.DeleteChanger(changerEmoji)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   changers,
	})
}
