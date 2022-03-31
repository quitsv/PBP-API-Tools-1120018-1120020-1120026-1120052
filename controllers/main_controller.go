package controllers

import (
	"log"
	"net/smtp"
)

func SendEmail(emailPenerima string) {
	// Configuration
	from := "stevianianggila60@gmail.com"
	password := "NakNik919"
	to := []string{emailPenerima}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("My super secret message.")

	// Create authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatal(err)
	}
}
