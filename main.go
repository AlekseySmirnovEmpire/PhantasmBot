package main

import (
	"PhantasmBot/bot"
	"PhantasmBot/config"
	"PhantasmBot/db"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file!")
	}
}

func main() {
	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if err = bot.Start(); err != nil {
		return
	}
	defer db.CloseDB()

	<-make(chan struct{})
}
