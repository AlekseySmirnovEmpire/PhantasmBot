package character

import (
	"PhantasmBot/db"
	"fmt"
	"sync"
)

type noPlayerErr struct {
	err error
	val string
}

func (n noPlayerErr) Error() string {
	return n.val
}

func (n noPlayerErr) Unwrap() error {
	return n.err
}

type Character struct {
	ID                     int    `db:"p_id"`
	Name                   string `db:"p_name"`
	LastName               string `db:"last_name"`
	Klass                  string `db:"klass"`
	Race                   string `db:"race"`
	Height                 int    `db:"height"`
	Sex                    string `db:"sex"`
	Level                  int    `db:"p_level"`
	CurrentHealth          int    `db:"current_health"`
	CurrentMana            int    `db:"current_mana"`
	CreatedAt              string `db:"created_at"`
	Characteristics        *Characteristics
	Attributes             *Attributes
	CharacteristicsUpgrade *CharacteristicsUpgrade
	InventoryItems         *[]InventoryItems
	Money                  *Money
	Skills                 *[]Skills
}

func Init(name *string, userID int) (*Character, error) {
	sql := fmt.Sprintf(
		`SELECT p.p_id, p.p_name, p.last_name, p.klass, p.race, p.height, 
       		p.sex, p.p_level, p.current_health, p.current_mana, p.created_at
			FROM player AS p
			JOIN user_info AS u ON u.p_id = p.p_id
			WHERE u.u_id = %d AND p.p_name = '%s'`, userID, *name)
	players, err := db.Select[Character](&sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if len(players) == 0 {
		return nil, noPlayerErr{val: "нет такого персонажа!", err: err}
	}
	c := players[0]

	wg := new(sync.WaitGroup)
	wg.Add(6)

	go func() {
		c.Characteristics = asyncCreate[Characteristics](InitCharacteristics, c.ID, wg)
		c.Attributes = asyncCreate[Attributes](InitAttr, c.ID, wg)
		c.CharacteristicsUpgrade = asyncCreate[CharacteristicsUpgrade](InitCharUp, c.ID, wg)
		c.InventoryItems = asyncCreateArr[InventoryItems](InitInvItems, c.ID, wg)
		c.Money = asyncCreate[Money](InitMoney, c.ID, wg)
		c.Skills = asyncCreateArr[Skills](InitSkills, c.ID, wg)
	}()

	wg.Wait()
	return &c, nil
}

func asyncCreate[T comparable](f func(int) *T, playerID int, wg *sync.WaitGroup) *T {
	defer wg.Done()
	return f(playerID)
}

func asyncCreateArr[T comparable](f func(int) *[]T, playerID int, wg *sync.WaitGroup) *[]T {
	defer wg.Done()
	return f(playerID)
}
