package main

import (
	"github.com/bwmarrin/discordgo"
)

func onMessageCreate(dc *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
}
