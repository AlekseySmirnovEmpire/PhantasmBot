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
