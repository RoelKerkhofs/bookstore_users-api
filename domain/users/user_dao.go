package users

import (
	"bookstore/bookstore_users-api/datasources/mysql/users_db"
	"bookstore/bookstore_users-api/logger"
	"bookstore/bookstore_users-api/utils/errors"
	"fmt"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email_address, date_created, password, status) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser                = "SELECT id, first_name, last_name, email_address, date_created, status FROM users WHERE id=?; "
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email_address=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?"
	queryFindByStatus           = "SELECT id, first_name, last_name, email_address, date_created, status FROM users WHERE status=?"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email_address, date_created, status FROM users WHERE email_address=? AND password=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error when trying to get user by id", getErr)
		return errors.NewNotFoundError("ID niet gevonden")
	}

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		logger.Error("Error when trying to save user", saveErr)
		return errors.NewInternalServerError("Database error")
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to insert user", err)
		return errors.NewInternalServerError("Database error")
	}
	user.ID = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)

	if err != nil {
		logger.Error("Error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)

	if err != nil {
		logger.Error("Error when trying to prepare update user ", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error("Error when trying to  delete user", err)
		return errors.NewInternalServerError("Database error")
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("Error when trying to prepare find by status statement", err)
		return nil, errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("Error when trying to prepare get user by status", err)
			return nil, errors.NewInternalServerError("Database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("No users with status %s", status))
	}

	return results, nil
}

func (user *User) FindByEmailAndPAssword() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("Error when trying to prepare get user by email and password statement", err)
		return errors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("Error when trying to get user by email and password", getErr)
		return errors.NewNotFoundError("ID niet gevonden")
	}

	return nil
}
