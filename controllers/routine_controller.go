package controllers

func SendAsync() {
	message1 := []byte("Testing masuk 123")
	to1 := GetRecipient(1).Email
	message2 := []byte("Testing async")
	to2 := GetRecipient(1).Email

	go SendEmail(to1, message1)
	SendEmail(to2, message2)
}
