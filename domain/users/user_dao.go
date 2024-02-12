package users

import (
	"time"

	"anggi.blog/datasources/sqlite/crud_db"
	"anggi.blog/utils/errors"
	"anggi.blog/utils/logs"
)

var (
	queryInsertUser          = "INSERT INTO users(id, username, email, password) VALUES(?, ?, ?, ?)"
	queryGetUserByID         = "SELECT id, username, email, password FROM users WHERE id = ?"
	queryGetUsernameAndEmail = "SELECT id, username, email, password FROM users WHERE username = ? OR email = ?"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := crud_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logs.Error.Println(err)
		return errors.NewInternalServerError("Error when trying to prepare save user statement", "Data Access")
	}
	defer stmt.Close()

	userID := time.Now().Unix()
	insertResult, saveErr := stmt.Exec(userID, user.Username, user.Email, user.Password)
	if saveErr != nil {
		logs.Error.Println(saveErr)
		return errors.NewInternalServerError("Error when trying to save user", "Data Access")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logs.Error.Println(err)
		return errors.NewInternalServerError("Error when trying to get last insert id after creating a new user", "Data Access")
	}
	user.ID = userId
	return nil
}

func (user *User) GetByEmailAndUsername() *errors.RestErr {
	stmt, err := crud_db.Client.Prepare(queryGetUsernameAndEmail)
	if err != nil {
		return errors.NewInternalServerError("Error when trying to prepare get user by username and email statement", "Data Access")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Email)
	if getErr := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password); getErr != nil {
		logs.Error.Println(getErr)
		return errors.NewInternalServerError("User not found", "Data Access")
	}
	return nil
}

func (user *User) GetByID() *errors.RestErr {
	stmt, err := crud_db.Client.Prepare(queryGetUserByID)
	if err != nil {
		return errors.NewInternalServerError("Error when trying to prepare get user by id statement", "Data Access")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.Username, &user.Email, &user.Password); getErr != nil {
		logs.Error.Println(getErr)
		return errors.NewInternalServerError("User not found", "Data Access")
	}
	return nil
}
