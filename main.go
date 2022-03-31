package main

import (
	"apiTools/controllers"
)

func main() {

	//goRedis

	//goMail
	message := []byte("Testing masuk 123")
	to := controllers.GetRecipient().Email
	controllers.SendEmail(to, message)

	//goRoutine
	controllers.SendAsync()
}
