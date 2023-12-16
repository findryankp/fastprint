
package routes

import (
	"fastprint/controllers"

	"github.com/labstack/echo/v4"
)

func UsersRouter(e *echo.Echo) {
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)
}
	