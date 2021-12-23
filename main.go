package main

import (
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		// .env読めなかった場合の処理
	}
	bot, _ := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)

	message := linebot.NewTextMessage("hello, yuki")
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
