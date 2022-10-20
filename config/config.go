package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	BotPrefix string
	Token     string
	Admin     string

	config *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
	Admin     string `json:"admin"`
}

// IsAdmin Проверяет является ли пользователь админом по его Id
func (c *configStruct) IsAdmin(ID *string) bool {
	str := *ID
	return str == c.Admin
}

func ReadConfig() error {
	fmt.Println("Reading from config.json ....")

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	config = new(configStruct)

	err = json.Unmarshal(file, config)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	Token = config.Token
	BotPrefix = config.BotPrefix
	Admin = config.Admin
	fmt.Println("Success reading config.json!")

	return nil
}
