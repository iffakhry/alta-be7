package user

import (
	_entities "be7/layered/entities"
	_userRepository "be7/layered/repository/user"
)

type UserUseCase struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserUseCase(userRepo _userRepository.UserRepositoryInterface) UserUseCaseInterface {
	return &UserUseCase{
		userRepository: userRepo,
	}
}

func (uuc *UserUseCase) GetAll() ([]_entities.User, error) {
	users, err := uuc.userRepository.GetAll()
	return users, err
}
