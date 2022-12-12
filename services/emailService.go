package services

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"

	"github.com/aldisaputra17/dapur-fresh-id/config"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
)

type EmailService interface {
	Create()
}

type email struct {
}

func NewEmailService() EmailService {
	return &email{}
}

func (service *email) Create() {
	msg := helpers.RandomString(10)
	to := []string{"recipient1@gmail.com", "emaillain@gmail.com"}
	subject := "this is your otp code, don't tell anyone"
	cc := []string{"aldisaput17@gmail.com"}
	body := "From: " + config.SenderName() + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		msg
	auth := smtp.PlainAuth("", config.AuthEmail(), config.AuthPassword(), config.SmptHost())
	smtAddress := fmt.Sprintf("%s:%d", config.SmptHost(), config.SmptPort())

	err := smtp.SendMail(smtAddress, auth, config.AuthEmail(), append(to, cc...), []byte(body))
	if err != nil {
		log.Fatal(err.Error())
	}
}
