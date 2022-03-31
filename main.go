package main

import (
	"apiTools/controllers"
)

func main() {

	//goMail
	message := []byte("Testing masuk 123")
	controllers.SendEmail("if-20018@students.ithb.ac.id", message)

	//goRoutine
	controllers.SendAsync()
}
