package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/home")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "home page")
	})
}
