package apis

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
	"github.com/goomadao/subscribeManager/util/logger"
	"github.com/rakyll/statik/fs"
	"go.uber.org/zap"

	//web dist directory
	_ "github.com/goomadao/subscribeManager/statik"
)

var router *gin.Engine

//InitGin init gin router
func InitGin(port int) {
	router := gin.Default()
	router.GET("/api/groups", getGroups)
	router.POST("/api/group", group)

	router.POST("/api/node", node)

	router.GET("/api/rules", getRules)
	router.POST("/api/rule", rule)

	router.GET("/api/selectors", getSelectors)
	router.POST("/api/selector", selector)

	router.GET("/api/changers", getChangers)
	router.POST("/api/changer", changer)

	router.POST("/api/updateall", updateAll)

	router.GET("/api/sub", subscribe)

	statikFS, err := fs.New()
	if err != nil {
		logger.Logger.Fatal("Create file systemd fail", zap.Error(err))
	}
	router.StaticFS("/dist", statikFS)
	// router.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "")
	// })

	router.Run(":" + strconv.Itoa(port))
}

func subscribe(c *gin.Context) {
	config.LoadConfig()
	class := c.Query("class")
	switch class {
	case "clash":
		generateClash(c)
	case "clashr":
		generateClashR(c)
	}
	err := config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}

func updateAll(c *gin.Context) {
	config.LoadConfig()
	err := config.UpdateAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	err = config.WriteToFile()
	if err != nil {
		logger.Logger.Panic(err.Error())
	}
	logger.Logger.Info("Write to file success")
}
