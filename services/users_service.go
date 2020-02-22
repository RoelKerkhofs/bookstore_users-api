package services

import (
	"bookstore/bookstore_users-api/domain/users"
	"bookstore/bookstore_users-api/utils/errors"
	"fmt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {

	result := &users.User{ID: userId}
	fmt.Println("service", string(userId))
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.ID)

	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil

	return nil, nil
}
