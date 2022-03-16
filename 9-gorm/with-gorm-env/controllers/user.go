package controllers

import (
	_entities "be7/withgorm/entities"
	"fmt"

	"gorm.io/gorm"
)

func LihatUsers(db *gorm.DB) []_entities.User {
	var users []_entities.User
	tx := db.Find(&users)
	if tx.Error != nil {
		// panic(tx.Error)
		fmt.Println("error ", tx.Error)
	}
	return users
}
