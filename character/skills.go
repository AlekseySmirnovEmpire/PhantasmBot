package character

import (
	"PhantasmBot/db"
	"fmt"
)

type Skills struct {
	ID          int    `db:"s_id"`
	PlayerID    int    `db:"p_id"`
	SkillName   string `db:"skill_name"`
	Description string `db:"description"`
	Rarity      string `db:"rarity"`
}

func (s *Skills) String() string {
	return fmt.Sprintf(
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nТип: %s\nНазвание: %s\nОписание: %s\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n",
		s.Rarity, s.SkillName, s.Description)
}

func InitSkills(playerID int) *[]Skills {
	sql := fmt.Sprintf(
		`SELECT s.s_id, stp.p_id, s.skill_name, s.description, s.rarity
			FROM skills AS s
			JOIN skill_to_player AS stp ON stp.s_id = s.s_id
			JOIN player AS p ON p.p_id = stp.p_id	
			WHERE p.p_id = %d`, playerID)
	skills, err := db.Select[Skills](&sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(skills) == 0 {
		return nil
	}
	return &skills
}
