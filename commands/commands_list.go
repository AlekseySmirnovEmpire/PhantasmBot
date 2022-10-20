package commands

import (
	"PhantasmBot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
)

func ListCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	switch m.Author.ID {
	case config.Admin:
		str, err := adminCommands()
		if err != nil {
			*str = "не могу выполнить команду!"
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(str))
		break
	default:
		str, err := userCommands()
		if err != nil {
			*str = "не могу выполнить команду!"
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(str))
	}
}

func adminCommands() (*string, error) {
	file, err := ioutil.ReadFile("./commands/text/admin_commands.txt")
	if err != nil {
		fmt.Println("File 'admin_commands.txt' not found!")
		return nil, err
	}
	str := fmt.Sprintf("доступные команды:\n%s", string(file))
	return &str, nil
}

func userCommands() (*string, error) {
	file, err := ioutil.ReadFile("./commands/text/user_commands.txt")
	if err != nil {
		fmt.Println("File 'user_commands.txt' not found!")
		return nil, err
	}
	str := fmt.Sprintf("доступные команды:\n%s", string(file))
	return &str, nil
}
