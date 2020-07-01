package Controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint
	username string
	password string
}

var user = User{
	ID:       1,
	username: "admin",
	password: "admin",
}

func Login(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json!!!")
		return
	}
	fmt.Println(user.username, u.username)
	if user.username != "admin" && user.password != "admin" {
		c.JSON(http.StatusNotFound, "Invalid credentials!!!")
		return
	}

	token, err := CreateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Can't generate Token")
		return
	}
	c.JSON(http.StatusOK, token)
}

func CreateToken(userid uint) (string, error) {
	var err error
	os.Setenv("ACCESS_TOKEN", "JSHAKdssfsdf")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["expiration"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_TOKEN")))
	if err != nil {
		return "", err
	}
	return token, err

}
