package list

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jasiu001/maestro/mark"
)

type Bucket interface {
	GetTranslation() string
	GetDescription() string
	GetWords() []string
	AmountOfWords() int
}

type Comparer interface {
	Compare(pattern, word []string) int
}

func CreateList(data []Bucket, comp Comparer) *List {
	list := &List{}
	list.comparer = comp
	list.items = make(map[string]*Item)
	list.passedItems = make(map[string]*Item)
	for _, element := range data {
		var i Item
		i.translation = element.GetTranslation()
		i.description = element.GetDescription()
		i.amoutOfWords = element.AmountOfWords()
		i.words = element.GetWords()
		i.resultMark = mark.InitMark()
		list.items[uuid.New().String()] = &i
	}
	list.currentItemId = randomItem(list.items)

	return list
}

type Item struct {
	translation  string
	description  string
	amoutOfWords int
	words        []string
	resultMark   *mark.Mark
}

type List struct {
	comparer      Comparer
	currentItemId string
	items         map[string]*Item
	passedItems   map[string]*Item
}

func (l *List) IsFinished() bool {
	return len(l.items) == 0
}

func (l *List) GetDescription() string {
	if l.currentItem() != nil {
		return fmt.Sprintf("%s \n %s", l.currentItem().translation, l.currentItem().description)
	}
	return ""
}

func (l *List) NumberOfWords() int {
	if l.currentItem() != nil {
		return l.currentItem().amoutOfWords
	}
	return 0
}

func (l *List) ExecuteResponse(data []string) {
	result := l.comparer.Compare(l.currentItem().words, data)
	l.currentItem().resultMark.UpdateMark(result)
	if l.currentItem().resultMark.Pass() {
		l.passedItems[l.currentItemId] = l.currentItem()
		delete(l.items, l.currentItemId)
		return
	}
	l.currentItem().resultMark.UpdateMark(result)
}

func (l *List) GetResult() string {
	if l.currentItem() != nil {
		return l.currentItem().resultMark.NameMark()
	}
	passItem := l.passedItems[l.currentItemId]
	l.currentItemId = randomItem(l.items)
	return passItem.resultMark.NameMark()
}

func (l *List) GetSummary() string {
	return fmt.Sprintf("There were %d words", len(l.passedItems))
}

func (l *List) currentItem() *Item {
	if item, ok := l.items[l.currentItemId]; ok {
		return item
	}
	return nil
}

func randomItem(items map[string]*Item) string {
	// TODO: implement randomness based on some rules (e.g. not two the same element in a row)
	for id, _ := range items {
		return id
	}
	return ""
}
