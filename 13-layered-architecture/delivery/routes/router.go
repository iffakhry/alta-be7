package routes

import (
	_userHandler "be7/layered/delivery/handler/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.GET("/users", uh.GetAllHandler())
	e.GET("/users/:id", uh.GetByIdHandler())
}
