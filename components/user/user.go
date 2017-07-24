package user

import "time"

//User Structure
type User struct {
	Id               int    `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	password         string `json:"password"`
	encrypedPassword string
	DateOfBirth      time.Time `json:"dateOfBirth"`
}
