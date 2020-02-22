package users

import (
	"bookstore/bookstore_users-api/datasources/mysql/users_db"
	"bookstore/bookstore_users-api/utils/date_utils"
	"bookstore/bookstore_users-api/utils/errors"
	"fmt"
	"strings"
)

var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email_address, date_created) VALUES (?, ?, ?, ?);"
	indexUniqueEmail = "email_address_UNIQUE"
)

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.ID]

	if result == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user id %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	fmt.Println("payload", user.FirstName, user.LastName, user.Email, user.DateCreated)

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}
	user.ID = userId
	return nil
}
