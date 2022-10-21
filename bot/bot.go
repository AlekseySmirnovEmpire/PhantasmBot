package bot

import (
	"PhantasmBot/commands"
	"PhantasmBot/config"
	"PhantasmBot/db"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var BotId string

func Start() error {
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)

	if err = db.InitDB(); err != nil {
		return err
	}

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Bot is running!")
	return nil
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
		case "чарлист":
			commands.FindCharList(s, m)
			break
		case "инвентарь":
			commands.ShowInventory(s, m)
			break
		case "скилы":
			commands.ShowSkills(s, m)
			break
		case "титул":
			commands.ShowTitle(s, m)
			break
		case "навыки":
			commands.ShowAttributes(s, m)
			break
		case "золото":
			commands.ShowMoney(s, m)
			break
		case "выйти":
			commands.Quite(s, m)
			break
		case "кикнуть":
			commands.KickPlayer(s, m)
			break
		case "игроки":
			commands.ShowPlayers(s, m)
			break
		default:
			if commands.IsMandiWord(&msg[0]) {
				commands.GetRandomShitAnswer(s, m)
			} else {
				commands.NotFound(s, m)
			}
		}
	}
}
