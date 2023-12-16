package routes

import (
	"fastprint/controllers"
	"github.com/labstack/echo/v4"
)

func KategoriRouter(e *echo.Echo) {
	g := e.Group("/kategori")
	g.GET("", controllers.Kategori)
	g.GET("/:id", controllers.KategoriById)
	g.POST("", controllers.KategoriCreate)
	g.PUT("/:id", controllers.KategoriUpdate)
	g.DELETE("/:id", controllers.KategoriDelete)
}
	