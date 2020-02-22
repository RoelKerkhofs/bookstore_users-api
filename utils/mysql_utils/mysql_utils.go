package mysql_utils

import (
	"bookstore/bookstore_users-api/utils/errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"strings"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError("No record found")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("Invalid data"))
	}

	return errors.NewInternalServerError("Error processing request")

}
