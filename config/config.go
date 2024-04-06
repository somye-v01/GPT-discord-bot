package config

import (
	"encoding/json"
	"fmt"
	"os"
	
)

var (
	Token     string
	BotPrefix string
	AuthToken string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"Token`
	BotPrefix string `json:"BotPrefix"`
	AuthToken string `json:"AuthToken"`
}

func ReadConfig() error {
	fmt.Println("Reading config...")

	file, err := os.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	AuthToken = config.AuthToken

	return nil
}
