package main

import (
	"bytes"
	"fmt"
	"sort"
	"sync"

	"github.com/bwmarrin/discordgo"
)

func board(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	c := make(chan UserData)
	userDataList := []UserData{}
	var wg sync.WaitGroup
	for _, userId := range users {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			userData, err := fetchUserData(id)
			if err != nil {
				fmt.Printf("error fetching user data for %s: %s\n", id, err)
				return
			}
			c <- *userData
		}(userId)
	}

	// close the channel after all the requests are complete
	go func() {
		wg.Wait()
		close(c)
	}()

	for userData := range c {
		userDataList = append(userDataList, userData)
	}

	// sort by rating ordinal in descending order
	sort.Slice(userDataList, func(i, j int) bool {
		return userDataList[i].RatingOrdinal > userDataList[j].RatingOrdinal
	})

	var buffer bytes.Buffer
	for idx, userData := range userDataList {
		place := idx + 1
		userDisplayString := userData.String()
		buffer.WriteString(fmt.Sprintf("%d: %s\n", place, userDisplayString))
	}

	boardMessage := buffer.String()
	if boardMessage == "" {
		boardMessage = "board is empty - add a user with !rankadd"
	}

	_, err = s.ChannelMessageSend(m.ChannelID, boardMessage)
	if err != nil {
		fmt.Println("failed to send reply message:", err)
	}
}
