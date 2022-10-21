package main

import (
	"PhantasmBot/bot"
	"PhantasmBot/config"
	"PhantasmBot/db"
	"fmt"
)

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
