package users

import (
	"fmt"
	"strings"

	usersdb "github.com/shayegh/bookstore_users-api/datasources/mysql/users_db"
	"github.com/shayegh/bookstore_users-api/utils/dateutils"
	"github.com/shayegh/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email,date_created) VALUES (?,?,?,?); "
)

func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = dateutils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE"){
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exist",user.Email))
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
	if err := usersdb.Client.Ping(); err != nil {
		panic(err)
	}
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user with id %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
