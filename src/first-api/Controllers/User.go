package Controllers

import (
	models "first/src/first-api/Model"
	utils "first/src/first-api/Utils"
	"first/src/first-api/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// func CreateUser(c *gin.Context) {
// 	var user Model.User
// 	c.BindJSON(&user)
// 	err := Model.SaveUser(&user)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

func GetAllUsers(c *gin.Context) {

	pagination := utils.GeneratePaginationFromRequest(c)
	var user models.User
	userLists, err := repo.GetAllUsers(&user, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": userLists,
	})

}

// func GetUserById(c *gin.Context) {

// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.FindUserById(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func UpdateUserById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.FindUserById(&user, id)
// 	if err != nil {
// 		println("no data found")
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}
// 	c.BindJSON(&user)
// 	err = Model.UpdateUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, user)
// 	}
// }

// func DeleteUserById(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var user Model.User
// 	err := Model.DeleteUser(&user, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}

// }
