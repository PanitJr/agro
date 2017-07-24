package user

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type userDAO struct {
	db *sql.DB
}

func newUserDAO(databaseConnection *sql.DB) *userDAO {
	return &userDAO{db: databaseConnection}
}

func (dao *userDAO) create(user *User) (*User, error) {
	transaction, err := dao.db.Begin()
	if err != nil {
		return nil, err
	}
	statement := fmt.Sprintf(`INSERT INTO users(username, encrytped_password, email, date_of_birth)
	VALUES('%v', '%v', '%v', '%v') RETURNING id`, user.Username, user.encrypedPassword, user.Email, user.DateOfBirth.Format(time.RFC3339))
	_, err = transaction.Prepare(statement)
	if err != nil {
		transaction.Rollback()
		return nil, err
	}
	err = transaction.QueryRow(statement).Scan(&user.Id)
	if err != nil {
		transaction.Rollback()
		return nil, err
	}
	transaction.Commit()
	return user, nil
}
func (dao *userDAO) update(user *User) (*User, error) {
	return &User{}, nil
}
func (dao *userDAO) delete(user *User) (*User, error) {
	return &User{}, nil
}
func (dao *userDAO) get(user *User) (interface{}, error) {
	statement := fmt.Sprintf(`SELECT * FROM users WHERE id = %v LIMIT 1`, user.Id)
	err := dao.db.QueryRow(statement).Scan(&user.Id, &user.Username, &user.encrypedPassword, &user.Email, &user.DateOfBirth)
	if err != nil {
		return err.Error(), err
	}
	return user, nil
}
func (dao *userDAO) list() ([]*User, error) {
	users := []*User{}
	statement := fmt.Sprintf(`SELECT * FROM users`)
	rows, err := dao.db.Query(statement)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.Id, &user.Username, &user.encrypedPassword, &user.Email, &user.DateOfBirth)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
