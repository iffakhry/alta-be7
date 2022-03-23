package user

import (
	_entities "be7/layered/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (ur *UserRepository) GetAll() ([]_entities.User, error) {
	var users []_entities.User
	tx := ur.database.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}
