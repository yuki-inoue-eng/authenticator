package route

import (
	"github.com/labstack/echo/v4"
	"github.com/yuki-inoue-eng/authenticator/internal/handlers"
)

func Routing(e *echo.Echo, handler handlers.Handler) {
	e.POST("/users", handler.RegisterUser())
}