package character

import (
	"PhantasmBot/db"
	"fmt"
)

type Note struct {
	ID       int    `db:"n_id"`
	PlayerID int    `db:"p_id"`
	Text     string `db:"note"`
}

func (n Note) String() string {
	return fmt.Sprintf("ЗАМЕТКИ:\n%s", n.Text)
}

func InitNotes(playerID int) *Note {
	query := fmt.Sprintf(`SELECT n.n_id, n.p_id, n.note 
		FROM notes AS n JOIN player AS p ON p.p_id = n.p_id WHERE p.p_id = %d`, playerID)
	notes, err := db.Select[Note](&query)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	if len(notes) == 0 {
		return nil
	}
	return &notes[len(notes)-1]
}
