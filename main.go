package main

import (
	"apiTools/controllers"
	"time"

	"github.com/go-co-op/gocron"
	redis "github.com/go-redis/redis/v8"
)

func main() {

	//goRedis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	controllers.SetRedis(rdb, "diskon", "Beli sekarang diskon 50%", 0)
	controllers.SetRedis(rdb, "gratis", "Beli 2 dapet 2 buku != gratis !!!!!!", 1)

	s := gocron.NewScheduler(time.UTC)
	//goMail
	message := []byte("Testing masuk 123")
	to := controllers.GetRecipient(1).Email
	controllers.SendEmail(to, message)

	//goRoutine
	s.Every(5).Second().Do(func() { controllers.SendAsync([]byte(controllers.GetRedis(rdb, "gratis"))) })
	s.StartBlocking()
}
