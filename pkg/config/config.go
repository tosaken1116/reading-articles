package config

import "os"

type Config struct {
	DiscordToken string
}

func NewConfig() *Config {
	DiscordToken := os.Getenv("DISCORD_BOT_TOKEN")
	return &Config{
		DiscordToken: DiscordToken,
	}
}
