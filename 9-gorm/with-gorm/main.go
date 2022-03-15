package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `gorm:"unique" json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

type Product struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
}

func InitDB() {
	// declare struct config & variable connectionString
	connectionString := "root:qwerty123@tcp(127.0.0.1:3306)/altagormdb?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Product{})
}

func init() {
	InitDB()
	InitialMigration()
}

func main() {

	fmt.Println("Masukkan pilihan anda? (1=Lihat)/(2=tambah)")
	var pilihan string
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		// select id, name, email, password, phone, address from users;
		// results, err := db.Query("select id, name, email, password, phone, address from users")
		// if err != nil {
		// 	panic(err.Error())
		// }

		var users []User
		tx := DB.Find(&users)
		if tx.Error != nil {
			// panic(tx.Error)
			fmt.Println("error ", tx.Error)
		}
		for _, value := range users {
			fmt.Println(value.ID, "-", value.Name)
		}
		// fmt.Println(users)
	case "2":
		newUser := User{}
		fmt.Println("Masukkan Nama:")
		fmt.Scanln(&newUser.Name)
		fmt.Println("Masukkan Email:")
		fmt.Scanln(&newUser.Email)
		fmt.Println("Masukkan Password:")
		fmt.Scanln(&newUser.Password)
		fmt.Println("Masukkan Phone:")
		fmt.Scanln(&newUser.Phone)
		fmt.Println("Masukkan Address:")
		fmt.Scanln(&newUser.Address)

		// var query = fmt.Sprintf("INSERT INTO users (name, email, password, phone, address) VALUES('%s','%s','%s','%s','%s' )", newUser.Name, newUser.Email, newUser.Password, newUser.Phone, newUser.Address)
		// _, err := db.Exec(query)
		tx := DB.Save(&newUser)
		if tx.Error != nil {
			// panic(tx.Error)
			fmt.Println("error when insert data")
		}
		if tx.RowsAffected == 0 {
			fmt.Println("insert failed")
		}
		fmt.Println("Insert successfully")

	}
}
