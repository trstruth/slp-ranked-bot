package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func rankAdd(s *discordgo.Session, userId string) {
	fmt.Println("rankadd called for:", userId)

	userManager, err := NewFileBasedUserManager("slp-ranked-bot.db")
	if err != nil {
		fmt.Println("Failed to create new FileBasedUserManager:", err)
		return
	}

	err = userManager.AddUser(userId)
	if err != nil {
		fmt.Println("Failed to add new user:", err)
		return
	}
}
