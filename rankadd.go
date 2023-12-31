package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func rankAdd(s *discordgo.Session, userId string) {
	fmt.Println("rankadd called for:", userId)
}
