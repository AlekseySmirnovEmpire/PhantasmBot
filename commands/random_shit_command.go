package commands

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
	"time"
)

var (
	randomShit = []string{
		"сам такой!",
		"а твоя мама сосёт kappa",
		"обиделся? Ну ничего, бывает, можешь посасать, мб пройдёт!",
		"а ты жирный!",
		"маму в кино водил!",
		"сынок, ну ты понял про маму, да?",
		"попка подгорела?",
		"потушить твою попку?",
		"welcome to the deep dark fantasies, fucking slave",
	}
	mandiWords = []string{
		"пидр",
		"сука",
		"алёша",
		"пнх",
		"пошёлнахуй",
		"ебать",
		"ебал",
		"матьебал",
	}
)

func IsMandiWord(str *string) bool {
	v := strings.ToLower(*str)
	for _, s := range mandiWords {
		if s == v {
			return true
		}
	}
	return false
}

func GetRandomShitAnswer(s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(randomShit))
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&randomShit[n], &m.Author.ID))
}
