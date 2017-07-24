package user

type userRepositoryImpl struct {
	dao *userDAO
}

func newUserRepositoryImpl(userDao *userDAO) *userRepositoryImpl {
	return &userRepositoryImpl{dao: userDao}
}

func (userRepository *userRepositoryImpl) create(user *User) (*User, error) {
	return userRepository.dao.create(user)
}
func (userRepository *userRepositoryImpl) update(user *User) (*User, error) {
	return userRepository.dao.update(user)
}
func (userRepository *userRepositoryImpl) delete(user *User) (*User, error) {
	return userRepository.dao.delete(user)
}
func (userRepository *userRepositoryImpl) get(user *User) (interface{}, error) {
	return userRepository.dao.get(user)
}
func (userRepository *userRepositoryImpl) list() ([]*User, error) {
	return userRepository.dao.list()
}
