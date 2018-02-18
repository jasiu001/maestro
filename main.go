package main

import (
	"github.com/jasiu001/maestro/list"
	"github.com/jasiu001/maestro/spreadsheets"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	data := spreadsheets.CreateData()
	newList := list.CreateList(data)
	spew.Dump(newList)
}
