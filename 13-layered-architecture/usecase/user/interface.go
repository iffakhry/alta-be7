package user

import (
	_entities "be7/layered/entities"
)

type UserUseCaseInterface interface {
	GetAll() ([]_entities.User, error)
}
