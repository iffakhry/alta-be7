package user

import (
	_userUseCase "be7/layered/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(userUseCase _userUseCase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (uh *UserHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status":  "failed",
				"message": "failed to fetch data",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "success",
			"message": "success get all data",
			"data":    users,
		})
	}
}
