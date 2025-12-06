package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func BotSessionHandler() error {
	// BotSession
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("unable to load enviroment: %v", err)
	}
	token := os.Getenv("DISCORD_BOT_TOKEN")
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}
		if m.Content == "hello" {
			s.ChannelMessageSend(m.ChannelID, "Halo")
		}
		session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	})
	err = session.Open()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	defer session.Close()
	fmt.Printf("Bot Online")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	return nil
}
