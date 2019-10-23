package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngdlong91/mq-runner/middleware/stats"
	"net/http"
)

func main() {
	server := gin.Default()

	server.GET("/index", stats.NewStatsMiddleware().Gin() ,func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":"success",
		})
	})

	server.POST("/stats", stats.NewStatsMiddleware().Gin() ,func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":"success",
		})
	})

	if err := server.Run(":3000"); err != nil {
		panic("cannot start server")
	}

}
