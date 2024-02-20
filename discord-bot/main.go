package main

import (
	"discord-bot/bot"
	"log"
	"os"
)

func main() {
	// load the bot token from .env
	botToken, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Bot Token was not found in the .env file")
	}

	AppID, ok := os.LookupEnv("APP_ID")
	if !ok {
		log.Fatal("Bot Token was not found in the .env file")
	}
	// start the bot
	bot.BotToken = botToken
	bot.AppID = AppID
	bot.Run()
}
