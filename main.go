package main

import (
	"apiTools/controllers"
  "time"
	"github.com/go-co-op/gocron"
)

func main() {
  
  s := gocron.NewScheduler(time.UTC)
	//goMail
	message := []byte("Testing masuk 123")
	to := controllers.GetRecipient(1).Email
	controllers.SendEmail(to, message)

	//goRoutine
  s.Every(5).Second().Do(func () {controllers.SendAsync()})
  s.StartBlocking()
}
