package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shayegh/bookstore_users-api/domain/users"
	"github.com/shayegh/bookstore_users-api/services"
	"github.com/shayegh/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {

	var user users.User
	// fmt.Println(user)

	// bytes , err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// if err:= json.Unmarshal(bytes,&user); err != nil {
	// 	fmt.Sprintln(err.Error())
	// 	return
	// }
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("invalid json input")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "implement me!")

}
