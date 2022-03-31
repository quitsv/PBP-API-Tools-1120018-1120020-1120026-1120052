package controllers

func SendAsync() {
	message1 := []byte("Testing masuk 123")
	message2 := []byte("Testing async")

	to := GetRecipient().Email
	go SendEmail(to, message1)
	SendEmail(to, message2)
}
