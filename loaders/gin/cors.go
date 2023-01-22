package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin"},
		MaxAge:       20 * time.Minute,
	})
}
