package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func rankAdd(s *discordgo.Session, m *discordgo.MessageCreate, userId string) {
	fmt.Println("rankadd called for:", userId)

	userManager, err := NewFileBasedUserManager("slp-ranked-bot.db")
	if err != nil {
		fmt.Println("Failed to create new FileBasedUserManager:", err)
		return
	}

	connectCode, err := sanitizeConnectCode(userId)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, err.Error())
		if err != nil {
			fmt.Println("failed to send reply message:", err)
		}
		return
	}

	err = userManager.AddUser(connectCode)
	if err != nil {
		fmt.Println("Failed to add new user:", err)
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Added user %s", connectCode))
	if err != nil {
		fmt.Println("failed to send reply message:", err)
	}
}
