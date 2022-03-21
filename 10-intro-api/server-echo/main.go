package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// var users []User

func main() {
	e := echo.New()
	e.GET("/hello", HelloController)
	e.GET("/users", GetAllUsersController)
	e.GET("/users/:id", GetUserByIdController)
	e.POST("/users", CreateUserController)
	e.POST("/users/binding", CreateUserBindController)
	e.Start(":8080")
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func GetAllUsersController(c echo.Context) error {
	user := User{Id: 1, Name: "alta", Email: "alta@mail.com", Password: "12345"}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})

}

// localhost:8080/users/10?query=qwerty&limit=5
func GetUserByIdController(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	q1 := c.QueryParam("query")
	limit := c.QueryParam("limit")
	user := User{Id: idInt, Name: "alta", Email: "alta@mail.com", Password: "12345"}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"query":   q1,
		"limit":   limit,
		"data":    user,
	})
}

func CreateUserController(c echo.Context) error {
	// get data from value
	name := c.FormValue("name")
	email := c.FormValue("email")
	id, _ := strconv.Atoi(c.FormValue("id"))
	password := c.FormValue("password")
	var user User
	user.Name = name
	user.Email = email
	user.Id = id
	user.Password = password

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"data":     user,
	})
}

func CreateUserBindController(c echo.Context) error {
	// get data from value
	var user User
	// c.Bind(&user)
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "failed bind data",
		})
	}
	// if user.Email == "admin@mail.com" {

	// }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"data":     user,
	})
}
