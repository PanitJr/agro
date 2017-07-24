package user

import (
	"github.com/go-ozzo/ozzo-validation"
)

//ValidateUser validate user will check wether user struct if correct follow it's structure
func (user User) ValidateUser() error {
	return validation.ValidateStruct(&user,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&user.Username, validation.Required, validation.Length(5, 128)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&user.Email, validation.Required, validation.Length(5, 128)),
	)
}
