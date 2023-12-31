package main

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore messages from ourself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if len(m.Content) == 0 || m.Content[0] != '!' {
		return
	}

	parts := strings.Split(m.Content, " ")
	command := parts[0]

	switch command {
	case "!board":
		board(s, m)
	case "!rankadd":
		if len(parts) < 2 {
			// TODO spit this out at the user
			log.Println("need to supply username")
			return
		}
		rankAdd(s, m, parts[1])
	}
}
