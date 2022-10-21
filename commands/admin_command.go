package commands

import (
	"PhantasmBot/config"
	"PhantasmBot/player"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type notAllowed struct {
	val string
}

func (n notAllowed) Error() string {
	if n.val == "" {
		return "не так быстро, котик, это команды админа."
	}
	return n.val
}

func (n notAllowed) Unwrap() error {
	return n
}

func checkForAdmin(ID *string) (string, error) {
	if config.Admin != *ID {
		return "", notAllowed{}
	}
	return "", nil
}

func KickPlayer(s *discordgo.Session, m *discordgo.MessageCreate) {
	str, err := checkForAdmin(&m.Author.ID)
	if err == nil {
		msg := strings.Split(m.Content, " ")
		if len(msg) != 2 {
			str = "неа, неправильно ввёл команду, алёша!"
		} else {
			str = player.QuiteChar(&msg[1])
		}
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func ShowPlayers(s *discordgo.Session, m *discordgo.MessageCreate) {
	str, err := checkForAdmin(&m.Author.ID)
	if err == nil {
		str = player.ShowPlayers()
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}
