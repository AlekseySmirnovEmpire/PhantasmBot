package commands

import (
	"PhantasmBot/player"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func FindCharList(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	if player.IsInGame(&m.Author.ID) {
		str = "вы уже в игре!"
	} else {
		msg := strings.Split(m.Content, " ")
		if len(msg) != 2 {
			str = "неверный ввод команды!"
		} else if err := player.FindCharacter(&m.Author.ID, &msg[1]); err != nil {
			str = "ой-ой, не могу загрузить персонажа!"
			fmt.Println(err.Error())
		} else {
			if str, err = player.PrintCharList(&m.Author.ID); err != nil {
				str = "не получилось загрузить твой чарлист :("
				fmt.Println(err.Error())
			}
		}
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func ShowTitle(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	if !player.IsInGame(&m.Author.ID) {
		str = "для начала зайдите за своего персонажа!"
	} else {
		str = player.ShowTitle(&m.Author.ID)
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func ShowAttributes(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	if !player.IsInGame(&m.Author.ID) {
		str = "для начала зайдите за своего персонажа!"
	} else {
		str = player.ShowAttributes(&m.Author.ID)
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func ShowMoney(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	if !player.IsInGame(&m.Author.ID) {
		str = "для начала зайдите за своего персонажа!"
	} else {
		str = player.ShowMoney(&m.Author.ID)
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func Quite(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	if !player.IsInGame(&m.Author.ID) {
		str = "чтобы выйти из персонажа, надо сначала за него зайти, мистер \"мозги\" ..."
	} else {
		str = player.QuiteChar(&m.Author.ID)
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func MakeNote(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	if !player.IsInGame(&m.Author.ID) {
		str = "чтобы оставить заметку - зайдите за персонажа!"
	} else {
		msg := strings.Split(m.Content, " ")
		if len(msg) < 2 {
			str = "неправильно ввёл команду!"
		} else {
			nT := strings.Join(msg[1:], " ")
			str = player.MakeNote(&m.Author.ID, &nT)
		}
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func CreateNewPlayer(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	if player.IsInGame(&m.Author.ID) {
		str = "выйдите из персонажа, чтобы создать нового!"
		_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
	}
	msg := strings.Split(m.Content, " ")
	if len(msg[1:]) != 6 {
		str = "неправильно введены данные!"
		_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
	}
	msg = msg[1:]
	str, err := player.AddNewPlayer(&msg, &m.Author.ID)
	if err != nil {
		str = "что-то пошло не так!"
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}
