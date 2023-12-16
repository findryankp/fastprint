package routes

import (
	"fastprint/controllers"
	"github.com/labstack/echo/v4"
)

func ProdukRouter(e *echo.Echo) {
	g := e.Group("/produk")
	g.GET("", controllers.Produk)
	g.GET("/:id", controllers.ProdukById)
	g.POST("", controllers.ProdukCreate)
	g.PUT("/:id", controllers.ProdukUpdate)
	g.DELETE("/:id", controllers.ProdukDelete)
}
	