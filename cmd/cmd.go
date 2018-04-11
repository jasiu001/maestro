package cmd

import (
	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
	b "github.com/jasiu001/maestro/bucket"
	"github.com/jasiu001/maestro/comparison"
	l "github.com/jasiu001/maestro/list"
	"strings"
)

const exit = "exit!"

type uWords struct {
	words []string
}

func Translate(c *ishell.Context) {
	list := c.Get("list").(*l.List)
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for {
		// get first random row
		index, words := list.GetWords()
		c.Printf("Translete %d words (%s) \n", len(words), strings.Join(list.GetTranslations(), ", "))
		c.Printf("[%s] \n \n", list.GetDescription())
		c.SetPrompt("--> ")
		inputData := []string{}

		// take from user as many words as bucket has
		for i := 0; i < len(words); i++ {
			inputData = append(inputData, c.ReadLine())
		}

		// check if bucket pass if yes remove bucket from list
		bucket := b.InitBucket(words, inputData)
		comparison.RunComparison(bucket)
		if bucket.Pass() {
			c.Println(green("You pass this word"))
			list.Remove(index)
		} else {
			c.Println(red("You fail this word"))
		}

		// break all if list is empty
		if list.IsEmpty() {
			break
		}

		// hard exit
		if exist, err := in_array(exit, inputData); exist || err != -1 {
			break
		}

		// clear data from user
		inputData = []string{}
	}

	c.Println("Exit translation")
	c.SetPrompt(">>> ")
}

func in_array(val string, array []string) (exists bool, index int) {
	exists = false
	index = -1

	for i, v := range array {
		if val == v {
			index = i
			exists = true
			return
		}
	}

	return
}
