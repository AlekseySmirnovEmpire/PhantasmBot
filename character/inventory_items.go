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
	Decription     string         `db:"decription"`
	ArmourPiercing int            `db:"armour_piercing"`
	PhisArmour     int            `db:"phis_Armour"`
	MagicArmour    int            `db:"magic_Armour"`
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
