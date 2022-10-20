package commands

import "github.com/bwmarrin/discordgo"

func StartCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	str := "Привет! Я - бот для проведения игр по системе Phantasm! Чтобы узнать что я могу напиши \"!команды\"."
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func NotFound(s *discordgo.Session, m *discordgo.MessageCreate) {
	str := "не знаю такой команды или пока не умею её делать!"
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}
