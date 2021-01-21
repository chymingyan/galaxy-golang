package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func userRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users page")
	})

	users.GET("/del", func(c *gin.Context) {
		c.JSON(http.StatusOK, "delete users page")
	})

	users.GET("/byid", func(c *gin.Context) {
		c.JSON(http.StatusOK, "by user id search users page")
	})
}

