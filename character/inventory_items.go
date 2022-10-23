package character

import (
	"PhantasmBot/db"
	"database/sql"
	"fmt"
)

type InventoryItems struct {
	ID             int            `db:"ii_id"`
	PlayerID       int            `db:"p_id"`
	IsEquiped      bool           `db:"is_equiped"`
	ItemName       string         `db:"item_name"`
	Description    string         `db:"decription"`
	ArmourPiercing int            `db:"armour_piercing"`
	PhisArmour     int            `db:"phis_armour"`
	MagicArmour    int            `db:"magic_armour"`
	Rarity         string         `db:"rarity"`
	Damage         sql.NullString `db:"damage"`
	Upgrades       sql.NullString `db:"upgrades"`
	Amount         int            `db:"amount"`
	Head           bool           `db:"head"`
	Body           bool           `db:"body"`
	Neck           bool           `db:"neck"`
	LeftHand       bool           `db:"left_hand"`
	RightHand      bool           `db:"right_hand"`
	SpecialOne     bool           `db:"special_one"`
	SpecialTwo     bool           `db:"special_two"`
	SpecialThree   bool           `db:"special_three"`
}

func (ii *InventoryItems) String() string {
	str := fmt.Sprint("\n______________________________\n")
	str += fmt.Sprintf("Редкость: %s\n", ii.Rarity)
	str += fmt.Sprintf("Название: %s\n", ii.ItemName)
	str += fmt.Sprint("Экипировано: ")
	if ii.IsEquiped {
		str += "Да\n"
	} else {
		str += "Нет\n"
	}
	if ii.Damage.Valid {
		str += fmt.Sprintf("Урон: %s\n", ii.Damage.String)
	}
	if ii.Upgrades.Valid {
		str += fmt.Sprintf("Сокеты: %s\n", ii.Upgrades.String)
	}
	if ii.PhisArmour != 0 {
		str += fmt.Sprintf("Физ. армор: %d\n", ii.PhisArmour)
	}
	if ii.MagicArmour != 0 {
		str += fmt.Sprintf("Маг. армор: %d\n", ii.MagicArmour)
	}
	str += fmt.Sprintf("Количество: %d\n", ii.Amount)
	var pl string
	switch {
	case ii.Head:
		pl = "Голова"
		break
	case ii.Body:
		pl = "Тело"
		break
	case ii.LeftHand:
		pl = "Левая рука"
		break
	case ii.RightHand:
		pl = "Правая рука"
		break
	case ii.Neck:
		pl = "Шея"
		break
	case ii.SpecialOne:
		pl = "Спец. слот"
		break
	case ii.SpecialTwo:
		pl = "Спец. слот"
		break
	case ii.SpecialThree:
		pl = "Спец. слот"
		break
	default:
		pl = "НИКУДА"
		break
	}
	str += fmt.Sprintf("Одевается: %s\n", pl)
	str += fmt.Sprintf("Описание: %s\n", ii.Description)
	str += fmt.Sprint("______________________________\n")
	return str
}

func InitInvItems(playerID int) *[]InventoryItems {
	sql := fmt.Sprintf(
		`SELECT ii.ii_id, ii.p_id, ii.is_equiped, ii.item_name, ii.decription, ii.armour_piercing, ii.phis_Armour,
       		ii.magic_Armour, ii.rarity, ii.damage, ii.upgrades, ii.amount, ii.head, ii.body, ii.neck, ii.left_hand,
       		ii.right_hand, ii.special_one, ii.special_two, ii.special_three
			FROM inventory_items AS ii
			JOIN player AS p ON p.p_id = ii.p_id
			WHERE p.p_id = %d`, playerID)
	invItems, err := db.Select[InventoryItems](&sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(invItems) == 0 {
		return nil
	}
	return &invItems
}
