package utils

import (
	"encoding/json"
	"os"
)

//Config infomation for the application
type Config struct {
	Discord struct {
		// Discord Api Key
		APIKey string `json:"apikey"`
	} `json:"discord"`
}

//ReadConfig json config file to Config struct
func ReadConfig(filename string) (Config, error) {
	var c Config //out

	f, err := os.Open(filename)
	if err != nil {
		return c, err
	}
	defer f.Close()

	j := json.NewDecoder(f)
	err = j.Decode(&c)
	return c, err
}
