package main

import (
	"github.com/jasiu001/maestro/cli"
	"github.com/jasiu001/maestro/list"
)

func main() {
	app := list.List{Parts: []string{
		"this is description #1",
		"this is description #2",
	}}
	cli.RunCmdApplication(&app)
}
