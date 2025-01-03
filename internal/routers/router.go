package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	var r = gin.Default()
	var v1 = r.Group("/v1")
	{
		v1.GET("/ping", Pong)
	}
	return r
}

func Pong(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  }