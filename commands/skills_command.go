package commands

import (
	"PhantasmBot/player"
	"github.com/bwmarrin/discordgo"
)

func ShowSkills(s *discordgo.Session, m *discordgo.MessageCreate) {
	var str string
	if !player.IsInGame(&m.Author.ID) {
		str = "для начала зайдите за своего персонажа!"
	} else {
		str = player.ShowSkills(&m.Author.ID)
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}
