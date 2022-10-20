package commands

import "fmt"

func makeMessageWithPing(sP *string) string {
	return fmt.Sprintf("<@%s> %s", userId, *sP)
}
