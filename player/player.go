package player

import "PhantasmBot/character"

type Player struct {
	ID   string
	char character.Character
}

func Any() bool {
	return len(players) > 0
}

func IsInGame(ID *string) bool {
	for _, p := range players {
		if p.ID == *ID {
			return true
		}
	}
	return false
}

var (
	players []Player
)
