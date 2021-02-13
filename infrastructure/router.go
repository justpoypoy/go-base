package infrastructure

import (
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/justpoypoy/go-base/interfaces"
	"github.com/justpoypoy/go-base/usecases"
)

// Dispatch is handle routing
func Dispatch(logger usecases.Logger, sqlHandler interfaces.SQLHandler) {
	userController := interfaces.NewUserController(sqlHandler, logger)

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	route := gin.New()
	route.Use(
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	users(route, userController)

	route.NoRoute(noRoute)
	route.Run(":8088")
}

func users(route *gin.Engine, controller *interfaces.UserController) *gin.Engine {
	users := route.Group("/")
	{
		users.GET("users", controller.Index)
		users.POST("users", controller.Store)
		users.GET("users/:id", controller.Show)
		users.PUT("users/:id", controller.Update)
		users.DELETE("users/:id", controller.Destroy)
	}
	return route
}

func noRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to go rest api",
	})
}
