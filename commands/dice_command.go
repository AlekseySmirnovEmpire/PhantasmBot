package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"time"
)

func Dice(s *discordgo.Session, msg []string, ID *string, chanId *string) {
	var str string
	switch len(msg) {
	case 2:
		n, err := diceRandom(&msg[1])
		if err != nil {
			str = "вы указали не число!"
		} else {
			str = fmt.Sprintf("%d!", n)
		}
		break
	case 3:
		n, err := multipleDiceRandom(&msg[1], &msg[2])
		if err != nil {
			str = "вы указали не число!"
		} else {
			str = ""
			c := 0
			for _, v := range n {
				str += fmt.Sprintf("%d, ", v)
				c += v
			}
			str += fmt.Sprintf(" сумма: %d", c)
		}
		break
	default:
		str = "неверный формат ввода для кубиков!"
	}
	_, _ = s.ChannelMessageSend(*chanId, makeMessageWithPing(&str, ID))
}

func diceRandom(str *string) (int, error) {
	rand.Seed(time.Now().UnixNano())
	max, err := strconv.Atoi(*str)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return rand.Intn(max-0) + 1, nil
}

func multipleDiceRandom(f *string, s *string) ([]int, error) {
	rand.Seed(time.Now().UnixNano())
	max, err := strconv.Atoi(*f)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	num, err := strconv.Atoi(*s)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = rand.Intn(max-0) + 1
	}
	return arr, nil
}

func FiftyFiftyRandom(s *discordgo.Session, m *discordgo.MessageCreate) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	var str string
	if n >= 50 {
		str = "ДА!"
	} else {
		str = "НЕТ!"
	}
	_, _ = s.ChannelMessageSend(m.ChannelID, makeMessageWithPing(&str, &m.Author.ID))
}
