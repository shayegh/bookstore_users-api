package users

import (
	"database/sql"
	"fmt"
	"strings"

	usersdb "github.com/shayegh/bookstore_users-api/datasources/mysql/users_db"
	"github.com/shayegh/bookstore_users-api/utils/dateutils"
	"github.com/shayegh/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email,date_created) VALUES (?,?,?,?); "
	queryGetUser    = "SELECT id,first_name, last_name, email,date_created from users WHERE id = ?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id = ?;"
	queryDeleteUser = "DELETE from users WHERE id=?;"
)

func (user *User) Save() *errors.RestErr {
	stmt, restErr := prepStatement(queryInsertUser)
	if restErr != nil {
		return restErr
	}
	defer stmt.Close()

	user.DateCreated = dateutils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to insert user: %s", err.Error()))
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when getting user Id: %s", err.Error()))
	}

	user.Id = userId

	return nil

}

func (user *User) Get() *errors.RestErr {
	stmt, err := prepStatement(queryGetUser)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError(fmt.Sprintf("User with Id %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error while getting user %d : %s", user.Id, err.Error()))
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, restErr := prepStatement(queryUpdateUser)
	if restErr != nil {
		return restErr
	}
	defer stmt.Close()
	_, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, restErr := prepStatement(queryDeleteUser)
	if restErr != nil {
		return restErr
	}
	defer stmt.Close()

	_, err := stmt.Exec(user.Id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func prepStatement(queryString string) (*sql.Stmt, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryString)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return stmt, nil
}
