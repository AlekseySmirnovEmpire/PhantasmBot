package config

import (
	"fmt"
	"log"
	"os"
)

var (
	BotPrefix string
	Token     string
	Admin     string
)

type failLoad struct {
	val string
}

func (f failLoad) Error() string {
	return fmt.Sprintf("ERROR: there is no %s in .env!", f.val)
}

func (f failLoad) Unwrap() error {
	return f
}

// IsAdmin Проверяет является ли пользователь админом по его Id
func IsAdmin(ID *string) bool {
	return Admin == *ID
}

func ReadConfig() error {
	log.Println("Reading from config.json ....")

	var exist bool
	Token, exist = os.LookupEnv("Discord_Token")
	if !exist {
		return failLoad{val: "TOKEN"}
	}
	BotPrefix, exist = os.LookupEnv("Discord_Prefix")
	if !exist {
		return failLoad{val: "PREFIX"}
	}
	Admin, exist = os.LookupEnv("Discord_AdminID")
	if !exist {
		return failLoad{val: "ADMIN ID"}
	}
	log.Println("Success reading config.json!")

	return nil
}
