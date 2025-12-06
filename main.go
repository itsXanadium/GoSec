package main

import (
	"fmt"

	bot "github.com/itsXanadium/GoSec/service"
)

func main() {
	if err := bot.BotSessionHandler(); err != nil {
		fmt.Printf("bot not started: %v", err)

	}
}
