package user

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository *userRepositoryImpl
}

func newUserServiceImpl(userRero *userRepositoryImpl) *userServiceImpl {
	return &userServiceImpl{userRepository: userRero}
}

func (userService *userServiceImpl) Create(user *User) (*User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.encrypedPassword = string(bytes)
	return userService.userRepository.create(user)
}

func (userService *userServiceImpl) Update(user *User) (*User, error) {
	return userService.userRepository.update(user)
}

func (userService *userServiceImpl) Delete(user *User) (*User, error) {
	return userService.userRepository.delete(user)
}

func (userService *userServiceImpl) Get(user *User, id string) (interface{}, error) {
	user.Id, _ = strconv.Atoi(id)
	return userService.userRepository.get(user)
}

func (userService *userServiceImpl) List() ([]*User, error) {
	return userService.userRepository.list()
}
