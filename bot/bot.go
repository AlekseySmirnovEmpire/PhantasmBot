package bot

import (
	"PhantasmBot/commands"
	"PhantasmBot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var BotId string

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
	<-make(chan struct{})
	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, config.BotPrefix) {
		if m.Author.ID == BotId {
			return
		}

		m.Content = strings.TrimPrefix(m.Content, config.BotPrefix)

		if m.Content == "" {
			msg, err := s.ChannelMessages(m.ChannelID, 1, "", "", m.ID)
			if err != nil {
				fmt.Errorf("unable to get messages: %s", err)
				return
			}
			m.Content = msg[0].Content
			m.Attachments = msg[0].Attachments
		}
		msg := strings.Split(m.Content, " ")
		switch msg[0] {
		case "старт":
			commands.StartCommand(s, m)
			break
		case "команды":
			commands.ListCommand(s, m)
			break
		case "д":
			commands.Dice(s, msg, &m.Author.ID, &m.ChannelID)
			break
		case "данет":
			commands.FiftyFiftyRandom(s, m)
			break
		}
	}
}
