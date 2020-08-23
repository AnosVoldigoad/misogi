package main

import (
	"github.com/bwmarrin/discordgo"
)

func AvatarCommands(dc *discordgo.Session, member *discordgo.Member, message *discordgo.Message, args []string) error {
	var embed discordgo.MessageEmbed

	embed.Image = &discordgo.MessageEmbedImage{
		URL: member.User.AvatarURL("2048"),
	}

	_, err := dc.ChannelMessageSendEmbed(message.ChannelID, &embed)
	if err != nil {
		return err
	}

	return nil
}
