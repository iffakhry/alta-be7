package user

import (
	_entities "be7/layered/entities"
)

type UserRepositoryInterface interface {
	GetAll() ([]_entities.User, error)
}
