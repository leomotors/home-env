package services

import (
	"os"
	"strconv"
)

type Secret struct {
	// Password for client to access this app
	PASSWORD string

	DISCORD_TOKEN      string
	DISCORD_CHANNEL_ID uint64

	DATABASE_URL string
}

var secret = Secret{}
var initialized = false

func parseSecret() {
	secret.PASSWORD = os.Getenv("PASSWORD")
	if secret.PASSWORD == "" {
		panic("PASSWORD environment variable not set.")
	}

	secret.DISCORD_TOKEN = os.Getenv("DISCORD_TOKEN")
	if secret.DISCORD_TOKEN == "" {
		panic("DISCORD_TOKEN environment variable not set.")
	}

	channelIDStr := os.Getenv("DISCORD_CHANNEL_ID")
	channelID, _ := strconv.ParseUint(channelIDStr, 10, 64)

	if channelID == 0 {
		panic("DISCORD_CHANNEL_ID environment variable not set or invalid number.")
	}

	secret.DISCORD_CHANNEL_ID = channelID

	secret.DATABASE_URL = os.Getenv("DATABASE_URL")
	if secret.DATABASE_URL == "" {
		panic("DATABASE_URL environment variable not set.")
	}

	initialized = true
}

func GetSecret() Secret {
	if !initialized {
		parseSecret()
		initialized = true
	}

	return secret
}
