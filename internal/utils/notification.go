package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// SendToTelegram send message to telegram
func SendToTelegram(message string) {
	telegramToken := os.Getenv("TELEGRAM_CHATBOT_TOKEN")

	userID := os.Getenv("TELEGRAM_USER_ID")

	if len(telegramToken) == 0 || len(userID) == 0 {
		fmt.Println("TELEGRAM_CHATBOT_TOKEN not set")
		fmt.Println("Message wanna to send is:", message)
		return
	}

	if len(userID) == 0 {
		fmt.Println("TELEGRAM_USER_ID not set")
		fmt.Println("Message wanna to send is:", message)
		return
	}

	_, err := http.PostForm(`https://api.telegram.org/bot`+telegramToken+`/sendMessage`,
		url.Values{"chat_id": {userID}, "text": {message}})

	if err != nil {
		fmt.Println("Cannot notify user", err)
		fmt.Println("Message wanna to send is:", message)
		return
	}

	fmt.Println("Message sent to Telegram:", message)
}
