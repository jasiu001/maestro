package list

import "github.com/google/uuid"

type list struct {
	rows []row
}

func CreateList(rowsData []RowData) *list {

	var list list

	for _, rowData := range rowsData {
		newUUID, _ := uuid.NewRandom()
		row := row{
			id:           newUUID,
			kind:         rowData.GetKind(),
			words:        rowData.GetWords(),
			translations: rowData.GetTranslations(),
			description:  rowData.GetDescription(),
			repeat:       rowData.GetRepeat(),
		}

		list.rows = append(list.rows, row)
	}

	return &list
}
