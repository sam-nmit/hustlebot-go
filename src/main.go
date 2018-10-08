package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/bwmarrin/discordgo"

	"./plugin"
	"./utils"
)

var binder plugin.Binder

const (
	PluginDirectory = "plugin"
)

func main() {
	config, err := utils.ReadConfig("config.json")
	if err != nil {
		log.Fatalln("Failed to load config file.", err)
	}

	discord, err := discordgo.New("Bot " + config.Discord.APIKey)
	if err != nil {
		log.Fatalln("error creating Discord session.", err)
	}

	if f, err := ioutil.ReadDir(PluginDirectory); err != nil {
		for _, file := range f {
			if file.IsDir() {
				continue
			}
			log.Println("Loading", file.Name())
			binder.Add(filepath.Join(PluginDirectory, file.Name()))
		}
	} else {
		log.Fatalln("Failed to open plugins directory \"", PluginDirectory, "\". ", err)
	}

	binder.Init(discord)

}
