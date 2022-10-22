package commands

import (
	"PhantasmBot/config"
	"PhantasmBot/player"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
	"time"
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
	} else {
		str = err.Error()
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}

func ClearChat(s *discordgo.Session, m *discordgo.MessageCreate, botID *string) {
	str, err := checkForAdmin(&m.Author.ID)
	if err != nil {
		str = err.Error()
		_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
		return
	}
	msg := strings.Split(m.Content, " ")
	if len(msg) < 2 || len(msg) > 3 {
		str = "неа, неправильно ввёл команду, алёша!"
		_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
		return
	}

	flags, count, err := checkFlags(msg[1:])
	if err != nil {
		var naErr notAllowed
		if errors.As(err, &naErr) {
			str = naErr.Error()
		} else {
			str = "неа, неправильно ввёл команду, алёша!"
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
		return
	}

	var messages []*discordgo.Message
	t := time.Now()
	if !checkTargetFlag(&flags, "-a") && count < 100 {
		messages, _ = s.ChannelMessages(m.ChannelID, count, "", "", "")
	} else if checkTargetFlag(&flags, "-a") {
		messages, _ = s.ChannelMessages(m.ChannelID, 100, "", "", "")
		if len(messages) == 100 {
			mID := messages[len(messages)-1].ID
			for {
				mm, _ := s.ChannelMessages(m.ChannelID, 100, mID, "", "")
				if mm != nil {
					messages = append(messages, mm...)
					if len(mm) < 100 || t.Sub(mm[len(mm)-1].Timestamp).Hours() > 24*14 {
						break
					}
					mID = mm[len(mm)-1].ID
				} else {
					break
				}
			}
		}
	} else {
		messages, _ = s.ChannelMessages(m.ChannelID, 100, "", "", "")
		mID := messages[len(messages)-1].ID
		for {
			count -= 100
			if count <= 0 {
				break
			}
			if count <= 100 {
				mm, _ := s.ChannelMessages(m.ChannelID, count, mID, "", "")
				if mm != nil {
					messages = append(messages, mm...)
				}
				break
			} else {
				mm, _ := s.ChannelMessages(m.ChannelID, 100, mID, "", "")
				if mm != nil {
					if len(mm) == 0 {
						break
					}
					mID = mm[len(mm)-1].ID
					messages = append(messages, mm...)
					if len(mm) < 100 || t.Sub(mm[len(mm)-1].Timestamp).Hours() > 24*14 {
						break
					}
				} else {
					break
				}
			}
		}
	}

	if checkTargetFlag(&flags, "-b") {
		for i := 0; i < len(messages); i++ {
			if messages[i].ID != *botID {
				messages = append(messages[:i], messages[i+1:]...)
			}
		}
	}

	mOut := make([]string, 0)
	for _, v := range messages {
		if t.Sub(v.Timestamp).Hours() < 24*14 {
			mOut = append(mOut, v.ID)
		} else {
			break
		}
	}

	for {
		if len(mOut) <= 100 {
			err = s.ChannelMessagesBulkDelete(m.ChannelID, mOut)
			break
		}
		err = s.ChannelMessagesBulkDelete(m.ChannelID, mOut[:100])
		mOut = mOut[100:]
	}
}

func checkFlags(flags []string) ([]string, int, error) {
	out := make([]string, 0)
	count := 0
	for _, f := range flags {
		val, err := strconv.Atoi(f)
		if err == nil && val > 0 {
			count = val
		} else {
			switch f {
			case "-a":
				out = append(out, f)
				break
			case "-b":
				out = append(out, f)
				break
			default:
				return nil, 0, notAllowed{val: fmt.Sprintf("\"%s\" - нет такого флага!", f)}
			}
		}
	}

	if count != 0 {
		if checkTargetFlag(&flags, "-a") {
			return nil, 0, notAllowed{val: "указан флаг \"-a\" и число!"}
		}
		return flags, count, nil
	}
	if checkTargetFlag(&flags, "-b") && !checkTargetFlag(&flags, "-a") {
		return nil, 0, notAllowed{val: "не указан флаг \"-a\" или число!"}
	}
	return flags, 0, nil
}

func checkTargetFlag(flags *[]string, flag string) bool {
	for _, f := range *flags {
		if f == flag {
			return true
		}
	}
	return false
}
