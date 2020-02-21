package users

import (
	"bookstore/bookstore_users-api/datasources/mysql/users_db"
	"bookstore/bookstore_users-api/utils/date_utils"
	"bookstore/bookstore_users-api/utils/errors"
	"fmt"
)

var (
	usersDB = make(map[int64]*User)
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
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email address %s is already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User id %d already exists", user.ID))
	}

	user.DateCreated = date_utils.GetNowString()

	usersDB[user.ID] = user
	return nil
}
