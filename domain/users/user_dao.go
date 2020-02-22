package users

import (
	"bookstore/bookstore_users-api/datasources/mysql/users_db"
	"bookstore/bookstore_users-api/utils/date_utils"
	"bookstore/bookstore_users-api/utils/errors"
	"fmt"
	"strings"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email_address, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email_address, date_created FROM users WHERE id=? "
	indexUniqueEmail = "email_address_UNIQUE"
	errorNoRow       = "no rows in result set"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRow) {
			return errors.NewNotFoundError(fmt.Sprintf("User with id %d not found", user.ID))
		}
		return errors.NewInternalServerError(fmt.Sprintf("something went ronk trying to read user %d: %s", user.ID, err.Error()))
	}

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
