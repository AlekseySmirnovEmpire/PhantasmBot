package commands

import (
	"PhantasmBot/config"
	"PhantasmBot/player"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
)

func ListCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	var err error
	switch m.Author.ID {
	case config.Admin:
		str, err = adminCommands()
		if err != nil {
			str = "не могу выполнить команду!"
		}
		break
	default:
		if player.IsInGame(&m.Author.ID) {
			str, err = playerCommands()
		} else {
			str, err = userCommands()
		}
		if err != nil {
			str = "не могу выполнить команду!"
		}
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func adminCommands() (string, error) {
	file, err := ioutil.ReadFile("./commands/text/admin_commands.txt")
	if err != nil {
		fmt.Println("File 'admin_commands.txt' not found!")
		return "", err
	}
	return fmt.Sprintf("доступные команды:\n%s", string(file)), nil
}

func userCommands() (string, error) {
	file, err := ioutil.ReadFile("./commands/text/user_commands.txt")
	if err != nil {
		fmt.Println("File 'user_commands.txt' not found!")
		return "", err
	}
	return fmt.Sprintf("доступные команды:\n%s", string(file)), nil
}

func playerCommands() (string, error) {
	file, err := ioutil.ReadFile("./commands/text/player_commands.txt")
	if err != nil {
		fmt.Println("File 'user_commands.txt' not found!")
		return "", err
	}
	return fmt.Sprintf("доступные команды:\n%s", string(file)), nil
}
