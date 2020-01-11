package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goomadao/subscribeManager/util/config"
)

func generateClash(c *gin.Context) {
	clashFile := config.GenerateClashConfig()
	c.String(http.StatusOK, string(clashFile))
}
