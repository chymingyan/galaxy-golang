package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	router = gin.Default()
)

// Run will start the server
func Run() {
	defaultRoutes()
	router.Run(":5000")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func defaultRoutes() {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "login page")
	})

	v1 := router.Group("/users")
	userRoutes(v1)

	v2 := router.Group("/home")
	homeRoutes(v2)
}
