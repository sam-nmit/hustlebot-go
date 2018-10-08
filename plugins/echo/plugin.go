package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type EchoPlugin struct{}

var Plugin EchoPlugin

const (
	CommandTrigger = "!echo "
)

func (p *EchoPlugin) Init(s *discordgo.Session) {
	s.AddHandler(p.OnMessage)
}

func (p *EchoPlugin) OnMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, CommandTrigger) {
		echo := m.Content[len(CommandTrigger):]
		s.ChannelMessageSend(m.ChannelID, echo)
	}
}
