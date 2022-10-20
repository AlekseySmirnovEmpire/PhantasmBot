package main

import (
	"PhantasmBot/bot"
	"PhantasmBot/config"
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

	<-make(chan struct{})
}
