package main

import (
	"apiTools/controllers"
)

func main() {

	//goMail
	message := []byte("Testing masuk 123")
	to := controllers.GetRecipient(1).Email
	controllers.SendEmail(to, message)

	//goRoutine
	controllers.SendAsync()
}
