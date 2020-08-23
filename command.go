package main

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix      string         = "!"
	prefixRegex *regexp.Regexp = regexp.MustCompile(`(?i)^\s*` + regexp.QuoteMeta(prefix) + `\s*`)
)

type UsageData struct {
	usage    string
	desc     string
	examples []string
}

type Command struct {
	name     string
	aliases  []string
	category string
	usages   []UsageData
	exec     func(dc *discordgo.Session, member *discordgo.Member, message *discordgo.Message, args []string) error
}

var Commands = []Command{
	{
		name:     "avatar",
		aliases:  []string{"av"},
		category: "utility",
		usages: []UsageData{
			{usage: "!avatar <member>", desc: "menampilkan avatar member", examples: []string{"!avatar @Misogi"}},
		},
		exec: AvatarCommands,
	},
}

func onMessageCreateCommandHandler(dc *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Message
	if msg == nil || msg.Author == nil || msg.Type != discordgo.MessageTypeDefault || msg.Author.Bot {
		return
	}

	match := prefixRegex.FindString(m.Content)
	if match != "" {
		args := strings.Fields(m.Content[len(match):])
		command := getCmd(strings.ToLower(args[0]))
		if command == nil {
			return
		}

		member, err := dc.GuildMember(m.GuildID, m.Author.ID)
		if err != nil {
			return
		}

		err = command.exec(dc, member, msg, args)
		if err != nil {
			return
		}
	}
}

func getCmd(command string) *Command {
	for _, it := range Commands {
		if command == it.name {
			return &it
		}
	}
	return nil
}
