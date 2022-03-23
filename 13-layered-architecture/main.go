package main

import (
	"be7/layered/configs"
	_userHandler "be7/layered/delivery/handler/user"
	_userRepository "be7/layered/repository/user"
	_userUseCase "be7/layered/usecase/user"
	"fmt"
	"log"

	_routes "be7/layered/delivery/routes"
	_utils "be7/layered/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	e := echo.New()
	_routes.RegisterPath(e, userHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
