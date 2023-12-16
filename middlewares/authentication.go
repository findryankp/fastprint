package middlewares

	import (
		"net/http"
		"os"
		"fastprint/helpers"
		"github.com/labstack/echo/v4"
	)
	
	func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, tokenString := helpers.HeaderToken(c)
	
			if token := c.Request().Header.Get(os.Getenv("HEADERAUTH")); tokenString == "" || tokenString == token {
				return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("request does not contain an access token"))
			}
	
			if err := helpers.ValidateToken(tokenString); err != nil {
				return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("request does not contain an valid token"))
			}
			return next(c)
		}
	}
	
	