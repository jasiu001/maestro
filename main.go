package main

import (
	"github.com/abiosoft/ishell"
	"github.com/jasiu001/maestro/cmd"
	"github.com/jasiu001/maestro/list"
	"github.com/jasiu001/maestro/spreadsheets"
)

func main() {
	// fetch data from e.g. google spreadsheets
	data := spreadsheets.CreateData()
	// create list based on fetched data
	newList := list.CreateList(data)

	// init shell
	shell := ishell.New()
	// set list to shell context
	shell.Set("list", newList)

	// add shell command
	shell.AddCmd(&ishell.Cmd{
		Name: "start",
		Help: "lets translate something",
		Func: cmd.Translate,
	})

	// run shell
	shell.Run()
}
