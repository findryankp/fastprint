package routes

import (
	"fastprint/controllers"
	"github.com/labstack/echo/v4"
)

func StatusRouter(e *echo.Echo) {
	g := e.Group("/status")
	g.GET("", controllers.Status)
	g.GET("/:id", controllers.StatusById)
	g.POST("", controllers.StatusCreate)
	g.PUT("/:id", controllers.StatusUpdate)
	g.DELETE("/:id", controllers.StatusDelete)
}
	