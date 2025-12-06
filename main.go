package main

import (
	"fmt"
	bot "go-discord-bot/service"
)

func main() {
	if err := bot.BotSessionHandler(); err != nil {
		fmt.Printf("bot not started: %v", err)

	}
}
