package cmd

import (
	"github.com/abiosoft/ishell"
	b "github.com/jasiu001/maestro/bucket"
	"github.com/jasiu001/maestro/comparison"
	l "github.com/jasiu001/maestro/list"
	"strings"
)

type uWords struct {
	words []string
}

func Translate(c *ishell.Context) {
	list := c.Get("list").(*l.List)

	for {
		// get first random row
		index, words := list.GetWords()
		c.Printf("Translete %d words (%s) \n", len(words), strings.Join(list.GetTranslations(), ", "))
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
			list.Remove(index)
		}

		// break all if list is empty
		if list.IsEmpty() {
			break
		}

		// clear data from user
		inputData = []string{}
	}

	c.Println("Exit translation")
	c.SetPrompt(">>> ")
}
