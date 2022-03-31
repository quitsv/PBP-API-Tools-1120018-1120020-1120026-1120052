package main

import (
	"github.com/PBPPrak/apiTools/controllers"
)

func main() {
	message := []byte("Testing masuk 123")
	controllers.SendEmail("if-20018@students.ithb.ac.id", message)
	println("Email sent")
}
