package list

import (
	"fmt"
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

type IDList interface {
	GenerateID() string
	RemoveID(string)
	RandomID() string
}

func CreateList(data []Bucket, comp Comparer, idl IDList) *List {
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

		id := idl.GenerateID()
		list.items[id] = &i
	}
	list.indexList = idl
	list.currentItemId = list.indexList.RandomID()

	return list
}

// Item represents one bucket of words with all data e.g. description and translation
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
	indexList     IDList
	items         map[string]*Item
	passedItems   map[string]*Item
}

// IsFinished checks if item list is empty, return true if is
func (l *List) IsFinished() bool {
	return len(l.items) == 0
}

// GetDescription returns full item details: description plus translation
func (l *List) GetDescription() string {
	if l.currentItem() != nil {
		return fmt.Sprintf("%s \n %s", l.currentItem().translation, l.currentItem().description)
	}
	return ""
}

// NumberOfWords return amount of words of current item
func (l *List) NumberOfWords() int {
	if l.currentItem() != nil {
		return l.currentItem().amoutOfWords
	}
	return 0
}

// GetWords return currem item words
func (l *List) GetWords() []string {
	return l.currentItem().words
}

// ExecuteResponse executes current item process, check words with external words by comparer
// if words pass remove current item from list and add it to pass items list
func (l *List) ExecuteResponse(data []string) {
	result := l.comparer.Compare(l.currentItem().words, data)
	l.currentItem().resultMark.UpdateMark(result)
	if l.currentItem().resultMark.Pass() {
		l.passedItems[l.currentItemId] = l.currentItem()
		delete(l.items, l.currentItemId)
		l.indexList.RemoveID(l.currentItemId)
		return
	}
}

// GetResult returns result of current or passed item
// if item exists on item list return current not passed item result
// if item not exists on item list it means item passed and return item result from passed list
// in both cases chose other item as current item
func (l *List) GetResult() string {
	if l.currentItem() != nil {
		currentItemMark := l.currentItem().resultMark.NameMark()
		l.currentItemId = l.indexList.RandomID()
		return currentItemMark
	}
	passItem := l.passedItems[l.currentItemId]
	l.currentItemId = l.indexList.RandomID()
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
