package commands

import "github.com/bwmarrin/discordgo"

var (
	userId string
)

func StartCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	userId = m.Author.ID
	str := "Привет! Я - бот для проведения игр по системе Phantasm! Чтобы узнать что я могу напиши \"!команды\"."
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str))
}
