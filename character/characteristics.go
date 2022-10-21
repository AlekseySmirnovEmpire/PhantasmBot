package character

import (
	"PhantasmBot/db"
	"fmt"
)

type Characteristics struct {
	ID           int `db:"c_id"`
	PlayerID     int `db:"p_id"`
	Strength     int `db:"strength"`
	Dexterity    int `db:"dexterity"`
	Constitution int `db:"constitution"`
	Intelligance int `db:"intelligance"`
	Wisdom       int `db:"wisdom"`
	Charisma     int `db:"charisma"`
	PhisArmour   int `db:"phis_armour"`
	MagicArmour  int `db:"magic_armour"`
	HealthMax    int `db:"health_max"`
	ManaMax      int `db:"mana_max"`
}

func InitCharacteristics(playerID int) *Characteristics {
	sql := fmt.Sprintf(
		`SELECT c.c_id, c.p_id, c.strength, c.dexterity, c.constitution, c.intelligance, c.wisdom, c.charisma,
       		c.phis_armour, c.magic_armour, c.health_max, c.mana_max
			FROM characteristics AS c
			JOIN player AS p ON p.p_id = c.p_id
			WHERE p.p_id = %d`, playerID)
	characteristics, err := db.Select[Characteristics](&sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(characteristics) == 0 {
		return nil
	}
	c := characteristics[0]
	return &c
}
