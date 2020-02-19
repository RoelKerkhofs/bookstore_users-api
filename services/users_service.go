package services

import (
	"bookstore/bookstore_users-api/domain/users"
	"bookstore/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
