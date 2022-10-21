package character

import (
	"PhantasmBot/db"
	"fmt"
)

type Attributes struct {
	ID           int `db:"a_id"`
	PlayerID     int `db:"p_id"`
	Athletic     int `db:"athletic"`
	MeleeFight   int `db:"melee_fight"`
	SwordFight   int `db:"sword_fight"`
	Acrobatic    int `db:"acrobatic"`
	HandAgil     int `db:"hand_agil"`
	Hack         int `db:"hack"`
	Sneaky       int `db:"sneaky"`
	Constitution int `db:"constitution"`
	Analitic     int `db:"analitic"`
	History      int `db:"history"`
	Religion     int `db:"religion"`
	Attention    int `db:"attention"`
	Medicine     int `db:"medicine"`
	Insight      int `db:"insight"`
	Animals      int `db:"animals"`
	Perceptions  int `db:"perceptions"`
}

func InitAttr(playerID int) *Attributes {
	sql := fmt.Sprintf(
		`SELECT a.a_id, a.p_id, a.athletic, a.melee_fight, a.sword_fight, a.acrobatic, a.hand_agil, a.hack, 
       		a.sneaky, a.constitution, a.analitic, a.history, a.religion, a.attention, a.medicine, a.insight, a.animals,
			a.perceptions
			FROM attributes AS a
			JOIN player AS p ON p.p_id = a.p_id
			WHERE p.p_id = %d`, playerID)
	attrs, err := db.Select[Attributes](&sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(attrs) == 0 {
		return nil
	}
	a := attrs[0]
	return &a
}
