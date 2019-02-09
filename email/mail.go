package email

import (
	"crypto/tls"
	"emailservice/utils"
	"github.com/go-gomail/gomail"
)

type Sender struct {
	UserName string
	Password string
}

const (
	GOOGLE_SMTP_SERVER = "smtp.gmail.com"
	GOOLGE_SMTP_PORT   = 587
	AWS_SMTP_SERVER    = "email-smtp.us-east-1.amazonaws.com"
	AWS_AMTP_PORT      = 587
)

func NewSender(Username, Password string) Sender {
	return Sender{Username, Password}
}

func (sender Sender) SendMail(Subject, bodyMessage, From, To string) {

	utils.Info.Println("Trying sending message")
	m := gomail.NewMessage()
	m.SetHeader("From", From)
	m.SetHeader("To", To)
	m.SetHeader("Subject", Subject)
	m.SetBody("text/html", bodyMessage)

	d := gomail.NewDialer(AWS_SMTP_SERVER, AWS_AMTP_PORT, sender.UserName, sender.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email
	if err := d.DialAndSend(m); err != nil {
		utils.Info.Println(err.Error())
	} else {
		utils.Info.Printf("Email to %s successful", To)
	}
}
