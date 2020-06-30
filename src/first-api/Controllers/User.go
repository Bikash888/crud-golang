package Controllers

import (
	"first/src/first-api/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user Model.User
	c.BindJSON(&user)
	err := Model.SaveUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetAllUsers(c *gin.Context) {
	var user []Model.User
	err := Model.FindAllUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}

}

func GetUserById(c *gin.Context) {

	id := c.Params.ByName("id")
	var user Model.User
	err := Model.FindUserById(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Model.User
	err := Model.FindUserById(&user, id)
	if err != nil {
		println("no data found")
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&user)
	err = Model.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUserById(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Model.User
	err := Model.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

}
