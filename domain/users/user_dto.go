package users

import (
	"strings"

	"anggi.blog/utils/errors"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	if user.Username == "" {
		return errors.NewBadRequestError("Invalid username", "Domain")
	}
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email", "Domain")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid password", "Domain")
	}
	return nil
}
