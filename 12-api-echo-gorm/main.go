package main

import (
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Author    string `json:"author" form:"author"`
	Publisher string `json:"publisher" form:"publisher"`
}

func InitDB() {
	// declare struct config & variable connectionString
	connectionString := os.Getenv("CONNECTION_DB")

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}

func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func init() {
	InitDB()
	InitialMigration()
}

func responseSuccess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

func responseSuccessWithoutData(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": message,
	}
}

func responseFailed(message string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "failed",
		"message": message,
	}
}

// get all users
func GetUsersController(c echo.Context) error {
	var users []User
	tx := DB.Find(&users)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, responseFailed("failed to get users"))
	}
	return c.JSON(http.StatusOK, responseSuccess("success get all users", users))
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responseFailed("failed to bind data. please check your data."))
	}
	tx := DB.Save(&user)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, responseFailed("failed to save data"))
	}
	return c.JSON(http.StatusOK, responseSuccessWithoutData("success create data user"))
}

func main() {
	e := echo.New()
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
