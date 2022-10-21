package character

import (
	"PhantasmBot/db"
	"fmt"
)

type CharacteristicsUpgrade struct {
	ID           int `db:"cu_id"`
	PlayerID     int `db:"p_id"`
	Strength     int `db:"strength_up"`
	Dexterity    int `db:"dexterity_up"`
	Constitution int `db:"constitution_up"`
	Intelligance int `db:"intelligance_up"`
	Wisdom       int `db:"wisdom_up"`
	Charisma     int `db:"charisma_up"`
}

func InitCharUp(playerID int) *CharacteristicsUpgrade {
	sql := fmt.Sprintf(
		`SELECT cu.cu_id, cu.p_id, cu.strength_up, cu.dexterity_up, cu.constitution_up, cu.intelligance_up,
       		cu.wisdom_up, cu.charisma_up
			FROM characteristics_upgrade AS cu
			JOIN player AS p ON p.p_id = cu.p_id
			WHERE p.p_id = %d`, playerID)
	charUps, err := db.Select[CharacteristicsUpgrade](&sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(charUps) == 0 {
		return nil
	}
	cu := charUps[0]
	return &cu
}
