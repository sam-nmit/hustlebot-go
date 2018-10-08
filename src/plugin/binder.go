package plugin

import (
	"errors"
	"fmt"
	"plugin"

	"github.com/bwmarrin/discordgo"
)

type Binder struct {
	plugins []IPlugin
}

//Add plugin to binder
func (b *Binder) Add(filename string) error {
	plug, err := plugin.Open(filename)
	if err != nil {
		return err
	}

	inst, err := plug.Lookup("Plugin")
	if err != nil {
		return err
	}

	iplug, ok := inst.(IPlugin)
	if !ok {
		return errors.New(fmt.Sprint("\"Plugin\" not exposed in plugin file for", filename))
	}
	b.plugins = append(b.plugins, iplug)
	return nil
}

//Init all plugins.
func (b *Binder) Init(s *discordgo.Session) {
	for _, p := range b.plugins {
		p.Init(s)
	}
}
