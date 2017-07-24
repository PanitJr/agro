package user

import "database/sql"

func NewUserAPI(db *sql.DB) *APIUser {

	DAO := newUserDAO(db)
	Repository := newUserRepositoryImpl(DAO)
	Service := newUserServiceImpl(Repository)

	return &APIUser{userService: Service}
}
