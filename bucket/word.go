package bucket

import (
	"github.com/jasiu001/maestro/mark"
)

type Word struct {
	value  string
	rating *mark.Mark
}

func NewWord(word string) *Word {
	return &Word{
		value:  word,
		rating: mark.InitMark(),
	}
}

func (w Word) GetValue() string {
	return w.value
}

func (w Word) Rating() *mark.Mark {
	return w.rating
}
