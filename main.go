package main

import (
	"time"

	"github.com/PrakPBP/PBP-API-Tools-1120018-1120020-1120026-1120052/controllers"
	"github.com/go-co-op/gocron"
)

func main() {
	// message := []byte("Testing masuk 123")
	// controllers.SendEmail("if-20020@students.ithb.ac.id", message)
	// println("Email sent")

	s := gocron.NewScheduler(time.UTC)

	// s.Every(1).MonthLastDay().Do(controllers.SendEmail("if-20020@students.ithb.ac.id", message))
	s.Every(5).Second().Do(func() { controllers.SendEmail("if-20020@students.ithb.ac.id", []byte("spam")) })
	// s.Every(2).Second().Do(func() { fmt.Print("oooooooooo") })

	s.StartBlocking()
}
