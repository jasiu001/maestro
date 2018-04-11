package spreadsheets

import (
	"log"
	"strconv"
	"strings"
)

type Data struct {
	kind         string
	words        []string
	translations []string
	description  string
	repeat       byte
}

func separate(full string) []string {
	words := strings.Split(full, ";")

	for key, val := range words {
		words[key] = strings.TrimSpace(val)
	}

	return words
}

func prepareRepeatValue(value string) byte {
	repeat, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Cannot convert '%s' to integer", value)
	}

	return byte(repeat)
}

func InitData(rows []interface{}) Data {

	return Data{
		kind:         rows[0].(string),
		words:        separate(rows[1].(string)),
		translations: separate(rows[2].(string)),
		description:  rows[3].(string),
		repeat:       prepareRepeatValue(rows[4].(string)),
	}
}

func (d Data) GetKind() string {
	return d.kind
}

func (d Data) GetWords() []string {
	return d.words
}

func (d Data) GetTranslations() []string {
	return d.translations
}

func (d Data) GetDescription() string {
	return d.description
}

func (d Data) GetRepeat() byte {
	return d.repeat
}

func (d Data) IncreaseRepeat() {
	d.repeat++
}
