package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func getAuthToken() (string, error) {
	dat, err := os.ReadFile("./token.txt")
	token := string(dat)
	token = strings.ReplaceAll(token, "\n", "")

	return token, err
}

func main() {
	authToken, err := getAuthToken()
	if err != nil {
		log.Fatalf("failed to get auth token: %s", err)
	}

	bot, err := discordgo.New("Bot " + authToken)
	if err != nil {
		log.Fatalf("failed to create bot: %s", err)
	}

	bot.Identify.Intents = discordgo.IntentsGuildMessages

	bot.AddHandler(messageHandler)

	err = bot.Open()
	if err != nil {
		log.Println("error opening connection:", err)
		return
	}

	log.Println("Bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	bot.Close()
}
