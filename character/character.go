package character

import (
	"PhantasmBot/db"
	"fmt"
	"sync"
)

type NoPlayerErr struct {
	err error
	val string
}

func (n NoPlayerErr) Error() string {
	return "нет такого персонажа!"
}

func (n NoPlayerErr) Unwrap() error {
	return n
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
	Note                   *Note
}

func (c *Character) String() string {
	pr := c.Level / 40
	return fmt.Sprintf(
		"Имя:%s\nПрозвище:%s\nКласс:%s\nРаса:%s\nРост:%d\nПол:%s\nПрестиж:%d\nУровень:%d\nХП:%d\nМП:%d\n%s\n",
		c.Name, c.LastName, c.Klass, c.Race, c.Height, c.Sex, pr, c.Level-pr*40, c.CurrentHealth, c.CurrentMana,
		c.Characteristics)
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
		return nil, NoPlayerErr{}
	}
	c := players[0]

	wg := new(sync.WaitGroup)
	wg.Add(7)

	go func() {
		c.Characteristics = asyncCreate[Characteristics](InitCharacteristics, c.ID, wg)
		c.Attributes = asyncCreate[Attributes](InitAttr, c.ID, wg)
		c.CharacteristicsUpgrade = asyncCreate[CharacteristicsUpgrade](InitCharUp, c.ID, wg)
		c.InventoryItems = asyncCreateArr[InventoryItems](InitInvItems, c.ID, wg)
		c.Money = asyncCreate[Money](InitMoney, c.ID, wg)
		c.Skills = asyncCreateArr[Skills](InitSkills, c.ID, wg)
		c.Note = asyncCreate[Note](InitNotes, c.ID, wg)
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
