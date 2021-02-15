package Controllers

import (
	"crypto/tls"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

//SendMailUsingMailer sends email
func SendMailUsingMailer(c *gin.Context) {
	message := gomail.NewMessage()
	message.SetHeader("From", "noreply@testing.com")
	message.SetHeader("To", "bikashdulal150@gmail.com")
	message.SetHeader("Subject", "Hello!")
	message.SetBody("text/html", "Hello <h1>Golang</h1>")
	dail := gomail.NewDialer("smtp.gmail.com", 465, "your email is here", "email password is here")
	dail.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := dail.DialAndSend(message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "SucessFully sent"})

}
