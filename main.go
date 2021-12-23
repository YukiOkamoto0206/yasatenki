package main

import (
	"github.com/YukiOkamoto0206/yasatenki/weather"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		// .env読めなかった場合の処理
		log.Fatal(err)
	}
	bot, _ := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	result := weather.GetWeather()
	message := linebot.NewTextMessage(result)
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
}
