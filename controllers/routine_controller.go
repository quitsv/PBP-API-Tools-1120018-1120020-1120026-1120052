package controllers

func SendAsync(message []byte) {
	to1 := GetRecipient(1).Email

	go SendEmail(to1, message)
	SendEmail(to1, message)
}
