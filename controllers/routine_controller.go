package controllers

func SendAsync() {
	message1 := []byte("Testing masuk 123")
	message2 := []byte("Testing async")

	to := getRecipient().Email
	go SendEmail(to, message1)
	go SendEmail(to, message2)
}
