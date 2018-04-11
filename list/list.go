package list

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type List struct {
	rows      []row
	activeRow row
}

func CreateList(rowsData []RowData) *List {
	var list List

	for _, rowData := range rowsData {
		newUUID, _ := uuid.NewRandom()
		r := row{
			id:           newUUID,
			kind:         rowData.GetKind(),
			words:        rowData.GetWords(),
			translations: rowData.GetTranslations(),
			description:  rowData.GetDescription(),
			repeat:       rowData.GetRepeat(),
		}

		list.rows = append(list.rows, r)
	}
	list.activateRow()

	return &list
}

func (l *List) activateRow() {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	index := r.Intn(len(l.rows))

	l.activeRow = l.rows[index]
}

func (l *List) GetTranslations() []string {
	return l.activeRow.translations
}

func (l List) GetWords() (uuid.UUID, []string) {
	return l.activeRow.id, l.activeRow.words
}

func (l List) GetDescription() string {
	return l.activeRow.description
}

func (l List) IsEmpty() bool {
	return len(l.rows) == 0
}

func (l *List) Remove(index uuid.UUID) {
	for key, element := range l.rows {
		if element.id != index {
			continue
		}
		// remove element
		l.rows = append(l.rows[:key], l.rows[key+1:]...)
		if l.IsEmpty() {
			return
		}
		l.activateRow()
		return
	}
}
