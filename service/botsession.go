package service

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/itsXanadium/GoSec/tools"

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
		// The !scan command
		if strings.HasPrefix(m.Content, "!scan") {
			args := strings.Split(m.Content, " ")
			if len(args) < 2 {
				s.ChannelMessageSend(m.ChannelID, "To use command !scan: !scan <host>")
				return
			}
			host := args[1]
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Scanning %s.....", host))
			go func() {
				result := tools.Portscanner(host)
				s.ChannelMessageSend(m.ChannelID, result)
			}()
		}
		//!apkoi
		if strings.HasPrefix(m.Content, "!apkoi") {
			s.ChannelMessageSend(m.ChannelID, "http://apkoi.web.id/")
		}
		session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

		//!Developer
		if m.Content == "!dev" {
			s.ChannelMessageSend(m.ChannelID, "use !developer")
			return
		}
		if strings.HasPrefix(m.Content, "!developer") {
			s.ChannelMessageSend(m.ChannelID, "https://xanadium.my.id\nhttps://github.com/itsXanadium")
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
