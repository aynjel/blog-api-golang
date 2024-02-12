package services

import (
	"anggi.blog/domain/users"
	"anggi.blog/utils/errors"
	"anggi.blog/utils/logs"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		logs.Error.Println(err)
		return nil, err
	}

	pwSlice, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.NewBadRequestError("Error when trying to hash user password", "Service")
	}
	user.Password = string(pwSlice[:])

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user users.User) (*users.User, *errors.RestErr) {
	result := &users.User{Email: user.Email}

	if err := result.GetByEmailAndUsername(); err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password)); err != nil {
		return nil, errors.NewBadRequestError("Invalid user credentials", "Service")
	}

	resultWp := &users.User{ID: result.ID, Username: result.Username, Email: result.Email}
	return resultWp, nil
}

func GetUserByID(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userId}

	if err := result.GetByID(); err != nil {
		return nil, err
	}
	resultWp := &users.User{ID: result.ID, Username: result.Username, Email: result.Email}
	return resultWp, nil
}
