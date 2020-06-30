package Controllers

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/gin-gonic/gin"
)
var auth smtp.Auth

func MailSender(c *gin.Context){
	auth =smtp.PlainAuth("","89cc9543aa58eb","5b2e666923e86d","smtp.mailtrap.io")
	 
	indexTemplateData := struct{
		Name string
        URL string

	}{
		Name: "Bikash",
		URL: "https://www.youtube.com",
	}
	
	m:= NewRequestToSendMail([]string{"mail@golang.com"}, "Hello golang!", "Hello, World!")
    err := m.ParseTemplate("index.html", indexTemplateData)
    if err != nil {
		fmt.Println(err)
        ok, _ := m.SendEmail()
        fmt.Println(ok)
    }
}

type RequestToSendMail struct{
	from string
	to [] string
	subject string
	message string
}
func NewRequestToSendMail(to []string,subject string,message string) *RequestToSendMail{
	return &RequestToSendMail{
		to:to,
		subject: subject,
		message: message,
	}

}

func (r *RequestToSendMail) ParseTemplate(templateName string,data interface{})error{
	temp,err :=template.ParseFiles(templateName) //convert data value as a plain data
	if err !=nil{
		return err
	}
	buffer := new(bytes.Buffer)
	if err=temp.Execute(buffer,data); err !=nil{
		return err
	}
	r.message=buffer.String()
	return nil
}
func (r *RequestToSendMail)SendEmail()(bool,error){
	mime :="MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject :="Subject: "+r.subject+"\n"
	msg :=[]byte(subject + mime + "\n" + r.message)
	address :="smtp.mailtrap.io:25"
	if err := smtp.SendMail(address, auth, "dhanush@geektrust.in", r.to, msg); err != nil {
	   fmt.Println(err)
		return false, err
    }
    return true, nil

}