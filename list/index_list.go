package list

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type IndexList struct {
	last *LastIndexList
	list []string
}

func NewIndexList() *IndexList {
	rand.Seed(time.Now().Unix())
	return &IndexList{
		last: &LastIndexList{},
	}
}

func (i *IndexList) GenerateID() string {
	id := uuid.New().String()
	i.list = append(i.list, id)

	return id
}

func (i *IndexList) RemoveID(id string) {
	for index, value := range i.list {
		if value == id {
			i.list = append(i.list[:index], i.list[index+1:]...)
		}
	}
}

func (i IndexList) RandomID() string {
	switch len(i.list) {
	case 0:
		return ""
	case 1:
		return i.list[0]
	default:
		return i.random()
	}
}

func (i IndexList) random() string {
	randomIndex := rand.Intn(len(i.list))
	randomElement := i.list[randomIndex]

	if len(i.list) == 2 {
		i.last.add("")
	}
	if i.last.exist(randomElement) {
		return i.random()
	}

	i.last.add(randomElement)
	return randomElement
}
