package routes

import (
	"nocpos/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.GET("/health",
		func(c *gin.Context) {

			c.JSON(200, gin.H{
				"message": "API Running",
			})

		})

	r.POST(
		"/login",
		controllers.Login,
	)

}
