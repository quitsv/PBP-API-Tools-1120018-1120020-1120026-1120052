package controllers

import (
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendEmail(emailPenerima string) {
	var htmlBody = `
	<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
	<title>Hello, World</title>
	</head>
	<body>
	<p>This is an email using Go</p>
	</body>
	`
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = "if-20026@students.ithb.ac.id"
	server.Password = "g"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From Me <if-20026@students.ithb.ac.id>")
	email.AddTo("michaelleonardo69@gmail.com")
	// email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)
	// email.AddAttachment("super_cool_file.png")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}

}
