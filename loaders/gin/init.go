package gin

import (
	"LMWN-assignment/logs"
	"LMWN-assignment/router"
	"LMWN-assignment/utils/config"
	"github.com/gin-gonic/gin"
)

func Init() {
	gonic := gin.Default()

	// * Log
	logs.Info("Starting server ...")

	// * Trusted Proxies
	gonic.SetTrustedProxies(config.C.Cors)

	// * 404 path
	gonic.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// * Register router
	endpoints := gonic.Group("/")

	// * Router
	router.Router(endpoints)

	err := gonic.Run(":" + config.C.Port)
	if err != nil {
		logs.Error("[Error] failed to start Gin server due to: " + err.Error())
		return
	}
}
