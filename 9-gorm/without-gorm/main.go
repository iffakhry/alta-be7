//using RAW SQL
package main

//https://github.com/go-sql-driver/mysql
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
	Phone    string
	Address  string
}

func main() {
	//<username>:<passoword>@tcp(<host>:<port_DB>)/<db_name>
	db, err := sql.Open("mysql", "root:qwerty123@tcp(localhost:3306)/altagormdb")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Masukkan pilihan anda? (1=Lihat)/(2=tambah)")
	var pilihan string
	fmt.Scanln(&pilihan)

	switch pilihan {
	case "1":
		// select id, name, email, password, phone, address from users;
		results, err := db.Query("select id, name, email, password, phone, address from users")
		if err != nil {
			panic(err.Error())
		}

		var users []User
		for results.Next() {
			var user User
			err := results.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Phone, &user.Address)
			if err != nil {
				panic(err.Error())
			}
			users = append(users, user)
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

		var query = fmt.Sprintf("INSERT INTO users (name, email, password, phone, address) VALUES('%s','%s','%s','%s','%s' )", newUser.Name, newUser.Email, newUser.Password, newUser.Phone, newUser.Address)
		_, err := db.Exec(query)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Insert successfully")

	}
}
