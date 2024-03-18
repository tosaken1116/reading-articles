package bot

import (
	"fmt"
	"os"
	"os/signal"
	"reading-articles/pkg/config"
	"reading-articles/pkg/handler"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type IBot interface {
	InitializeBot()
}
type Bot struct {
	dbt      string
	session  *discordgo.Session
	handlers handler.IHandler
}

func (b Bot) InitializeBot() {
	var err error
	b.session, err = discordgo.New(fmt.Sprintf("Bot %s", b.dbt))
	if err != nil {
		fmt.Printf("Error creating Discord session: %v", err)
		return
	}

	err = b.session.Open()
	b.session.AddHandler(b.handlers.OnAnyoneJoinVoiceChannel)
	if err != nil {
		fmt.Printf("Error opening connection to Discord: %v", err)
		return
	}
	stopBot := make(chan os.Signal, 1)

	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-stopBot

	err = b.session.Close()

	if err != nil {
		fmt.Printf("Error closing connection to Discord: %v", err)
		return
	}
}

func NewBot(config config.Config, handlers *handler.IHandler) IBot {
	return &Bot{
		dbt:      config.DiscordToken,
		session:  nil,
		handlers: *handlers,
	}
}
