package character

import (
	"PhantasmBot/db"
	"fmt"
)

type Money struct {
	ID       int `db:"m_id"`
	PlayerID int `db:"p_id"`
	Gold     int `db:"gold"`
}

func InitMoney(playerID int) *Money {
	sql := fmt.Sprintf(
		`SELECT m.m_id, m.p_id, m.gold
			FROM money AS m
			JOIN player AS p ON p.p_id = m.p_id
			WHERE p.p_id = %d`, playerID)
	moneys, err := db.Select[Money](&sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(moneys) == 0 {
		return nil
	}
	m := moneys[0]
	return &m
}
