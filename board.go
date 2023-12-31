package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func board(s *discordgo.Session) {
	fmt.Println("board called")

	userManager, err := NewFileBasedUserManager("slp-ranked-bot.db")
	if err != nil {
		fmt.Println("failed to create new FileBasedUserManager:", err)
		return
	}

	users, err := userManager.GetUserList()
	if err != nil {
		fmt.Println("failed to get user list:", err)
		return
	}
	for _, userId := range users {
		fmt.Println(userId)
	}
}
