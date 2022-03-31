package controllers

import (
	"log"
	"net/smtp"

	_ "github.com/go-sql-driver/mysql"
)

func getSender() User {
	db := Connect()
	defer db.Close()

	var sender User
	err := db.QueryRow("SELECT email, passwd FROM Users where id = 1").Scan(&sender.Email, &sender.Passwd)
	if err != nil {
		log.Fatal(err)
	}

	return sender
}

func SendEmail(emailPenerima string, message []byte) {
	// Configuration
	Sender := getSender()
	to := []string{emailPenerima}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Confirmation
	println("Sending email to: " + emailPenerima)
	println("sender email: " + Sender.Email)

	// Create authentication
	auth := smtp.PlainAuth("", Sender.Email, Sender.Passwd, smtpHost)

	// Send actual message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, Sender.Email, to, message)
	if err != nil {
		log.Fatal(err)
	}
}
