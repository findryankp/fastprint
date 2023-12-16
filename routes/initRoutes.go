package routes

import (
	"github.com/labstack/echo/v4"
	"fastprint/helpers"
)

func InitRouter(e *echo.Echo){
	ProdukRouter(e)
	KategoriRouter(e)
	StatusRouter(e)
	UsersRouter(e)
	IndexRouter(e)
}

func IndexRouter(e *echo.Echo) {
	e.GET("/", index)
}

func index(c echo.Context) error {
	var data = map[string]interface{}{
		"message":       "Welcome to mvcGo",
		"developmen_by": "Findryankp",
	}

	return c.JSON(200, helpers.ResponseSuccess("-", data))
}
	
