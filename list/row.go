package list

import "github.com/google/uuid"

type RowData interface {
	GetKind() string
	GetWords() []string
	GetTranslations() []string
	GetDescription() string
	GetRepeat() byte
	IncreaseRepeat()
}

type row struct {
	id           uuid.UUID
	kind         string
	words        []string
	translations []string
	description  string
	repeat       byte
}
