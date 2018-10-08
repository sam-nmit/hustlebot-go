package plugin

import (
	"github.com/bwmarrin/discordgo"
)

type IPlugin interface {
	Init(s *discordgo.Session)
}
