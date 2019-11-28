package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vearne/tinyurl/handler"
	"net/http"
	"time"
)

//NewServer NewServer
func NewServer() *gin.Engine {
	r := gin.Default()

	r.POST("/api/tinyurl", handler.UrlChange)
	r.GET("/:sid", handler.UrlGet)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, time.Now().String())
	})
	return r
}
