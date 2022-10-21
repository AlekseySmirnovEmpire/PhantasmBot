package player

import (
	"PhantasmBot/character"
	"PhantasmBot/db"
	"fmt"
)

type User struct {
	UserKey  string `db:"user_key"`
	UserID   int    `db:"u_id"`
	PlayerID int    `db:"p_id"`
	Player   *character.Character
}

func Any() bool {
	return len(users) > 0
}

func IsInGame(ID *string) bool {
	for _, u := range users {
		if u.UserKey == *ID {
			return true
		}
	}
	return false
}

var (
	users []User
)

func FindCharacter(ID *string, name *string) error {
	sql := fmt.Sprintf(
		`SELECT u.u_id, u.user_key,u.p_id 
			FROM user_info AS u
			JOIN player AS p ON p.p_id = u.p_id
			WHERE u.user_key = '%s' AND p.p_name = '%s'`, *ID, *name)
	chars, err := db.Select[User](&sql)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	char := chars[0]
	char.Player = new(character.Character)
	if char.Player, err = character.Init(name, char.UserID); err != nil {
		fmt.Println(err.Error())
		return err
	}

	users = append(users, char)
	return nil
}

func PrintCharList(ID *string) (string, error) {
	var str string
	return str, nil
}
