package Controllers

import (
	"Go_OOP/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (user *UserController) NewUserController() {}

// @Summary Get All User Data
// @Description Get All User Data
// @ID GetAllUserData
// @Tags User Data
// @Success 200 {object} []Models.User "Success"
// @Failure 400 {string} string "Error"
// @response 401 {string} string "Unauthorized"
// @Router /user-api/user [GET]
// @security ApiKeyAuth
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// @Summary Register User
// @Description Register New User
// @ID RegisterUser
// @Tags User
// @Param EnterDetails body Models.UserRegister true "Register"
// @Accept json
// @Success 200 {object} Models.User "Success"
// @Failure 400 {string} string "Error"
// @Router /auth-api/register [POST]
func CreateUser(c *gin.Context) {
	var user Models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Models.CreateUser(&user)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUserById(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")

	err := Models.GetUserById(&user, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")

	err := Models.GetUserById(&user, id)

	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}

	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Models.DeleteUser(&user, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}

}
