package player

import (
	"PhantasmBot/character"
	"PhantasmBot/db"
	"fmt"
)

var (
	users []User
)

type User struct {
	UserKey  string `db:"user_key"`
	UserID   int    `db:"u_id"`
	PlayerID int    `db:"p_id"`
	Player   *character.Character
}

func (u User) String() string {
	var str string
	if u.Player.Note == nil {
		str = "Заметок нет."
	} else {
		str = u.Player.Note.String()
	}
	return fmt.Sprintf(
		"---------------------------------\nВаш ID: %s\n---------------------------------\n%s\n%s",
		u.UserKey,
		u.Player,
		str)
}

func Refresh() string {
	var str string
	findUser := func() <-chan error {
		ch := make(chan error, len(users))

		go func() {
			defer close(ch)
			for _, u := range users {
				if u.Player == nil {
					str += fmt.Sprintf("%s - игрок не подгружен, не обновился!\n")
					continue
				}
				ch <- FindCharacter(&u.UserKey, &u.Player.Name)
			}
		}()

		return ch
	}

	for val := range findUser() {
		if val != nil {
			str += val.Error()
		}
	}

	if str == "" {
		str += "данные успещно загружены!"
	}

	return str
}

func Any() bool {
	return len(users) > 0
}

func GetPlayer(ID *string) (*character.Character, error) {
	for _, p := range users {
		if p.UserKey == *ID {
			return p.Player, nil
		}
	}
	return nil, character.NoPlayerErr{}
}

func HealPlayer(heal int, playerName *string) string {
	var user *character.Character
	for _, u := range users {
		if u.Player.Name == *playerName {
			user = u.Player
		}
	}
	if user == nil {
		return "персонаж не в игре!"
	}
	user.CurrentHealth += heal
	if user.Characteristics == nil {
		return "характеристики не прогрузились!"
	}
	if user.CurrentHealth > user.Characteristics.HealthMax {
		user.CurrentHealth = user.Characteristics.HealthMax
	}
	ch := user.CurrentHealth
	query := fmt.Sprintf(`UPDATE player SET current_health = %d WHERE p_name = '%s'`, ch, user.Name)
	err := db.InsertOrUpdate(&query)
	if err != nil {
		return "что-то пошло не так, урон не нанесён!"
	}
	return fmt.Sprintf("%s получил %d хила, текущие ХП = %d.", user.Name, heal, user.CurrentHealth)
}

func DealDamage(dmg int, playerName *string, isFull bool) string {
	var user *character.Character
	for _, u := range users {
		if u.Player.Name == *playerName {
			user = u.Player
		}
	}
	if user == nil {
		return "персонаж не в игре!"
	}
	user.CurrentHealth -= dmg
	ch := user.CurrentHealth
	query := fmt.Sprintf(`UPDATE player SET current_health = %d WHERE p_name = '%s'`, ch, user.Name)
	err := db.InsertOrUpdate(&query)
	if err != nil {
		return "что-то пошло не так, урон не нанесён!"
	}
	if ch <= -user.Characteristics.HealthMax*2 || isFull {
		ch = -user.Characteristics.HealthMax * 2
		return fmt.Sprintf("%s умер! Он выведен из игры, его хп = %d.", user.Name, ch)
	} else {
		return fmt.Sprintf("%s получил %d урона, текущие ХП = %d.", user.Name, dmg, user.CurrentHealth)
	}
}

func ShowTitle(ID *string) string {
	player, err := GetPlayer(ID)
	if err != nil {
		return err.Error()
	}
	return player.String()
}

func ShowAttributes(ID *string) string {
	player, err := GetPlayer(ID)
	if err != nil {
		return err.Error()
	}
	if player.Attributes == nil {
		return "ой-ой что-то не так! (атрибуты не загрузились)"
	}
	return player.Attributes.String()
}

func ShowMoney(ID *string) string {
	player, err := GetPlayer(ID)
	if err != nil {
		return err.Error()
	}
	if player.Money == nil {
		return "ой-ой что-то не так! (деньги не подгрузились)"
	}
	return player.Money.String()
}

func ShowSkills(ID *string) string {
	player, err := GetPlayer(ID)
	if err != nil {
		return err.Error()
	}
	if player.Skills == nil || len(*player.Skills) == 0 {
		return "список ваших способностей пуст!"
	}
	str := ""

	f := func(skills *[]character.Skills) <-chan string {
		ch := make(chan string, len(*skills))

		go func() {
			defer close(ch)
			for _, v := range *skills {
				ch <- v.String()
			}
		}()

		return ch
	}

	for s := range f(player.Skills) {
		str += s
	}

	return str
}

func IsInGame(ID *string) bool {
	for _, u := range users {
		if u.UserKey == *ID {
			return true
		}
	}
	return false
}

func ShowInventory(ID *string) string {
	player, err := GetPlayer(ID)
	if err != nil {
		return err.Error()
	}
	if player.InventoryItems == nil || len(*player.InventoryItems) == 0 {
		return "ваш инвентарь пуст!"
	}
	str := ""

	for s := range getItem(player.InventoryItems) {
		str += s
	}

	return str
}

func getItem(ii *[]character.InventoryItems) <-chan string {
	ch := make(chan string, len(*ii))

	go func() {
		for _, i := range *ii {
			ch <- i.String()
		}
		close(ch)
	}()

	return ch
}

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
	var us *User
	for _, u := range users {
		if u.UserKey == *ID {
			us = &u
		}
	}
	if us == nil {
		str = "что-то пошло не так при печати персонажа! (его нет в игре)"
	} else {
		str = fmt.Sprintf("ВЫ ВОШЛИ В ИГРУ!\n%s", *us)
	}
	return str, nil
}

func QuiteChar(ID *string) string {
	var us *User
	var c int
	for i, u := range users {
		if u.UserKey == *ID {
			us = &u
			c = i
		}
	}
	if us == nil {
		return "а тебя и нет в списках!"
	}
	name := us.Player.Name
	users = append(users[:c], users[c+1:]...)
	return fmt.Sprintf("%s покинул игру!", name)
}

func ShowPlayers() string {
	if len(users) == 0 {
		return "никого в игре нет!"
	}
	var str string
	for i, u := range users {
		if u.Player == nil {
			str += fmt.Sprintf("%d. *неопознан* UID: %s; PID: %d;\n", i+1, u.UserKey, u.PlayerID)
		} else {
			str += fmt.Sprintf("%d. Персонаж: %s; UID:%s\n", i+1, u.Player.Name, u.UserKey)
		}
	}

	return str
}

func MakeNote(playerID, text *string) string {
	player, err := GetPlayer(playerID)
	if err != nil {
		return "вы не в игре!"
	}
	query := fmt.Sprintf(`INSERT INTO notes (p_id, note) VALUES (%d, '%s')`, player.ID, *text)
	err = db.InsertOrUpdate(&query)
	if err != nil {
		return "заметка не добавлена!"
	}
	return "заметка успешно добавлена!"
}
