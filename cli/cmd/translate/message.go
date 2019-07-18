package translate

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jasiu001/maestro/mark"
	"strings"
)

func Result(result string, words []string) {
	switch result {
	case mark.CORRECT_STR:
		color.Green(fmt.Sprintf("Result: %s", result))
	case mark.PROPER_STR, mark.SIMILAR_STR:
		color.Yellow(fmt.Sprintf("Best result: %s \n", result))
		properWords(words)
	default:
		color.Red(fmt.Sprintf("Best result: %s \n", result))
		properWords(words)
	}

	separator()
}

func separator() {
	color.White("--------------------")
}

func properWords(words []string) {
	fmt.Printf("(Correct words: %s) \n", strings.Join(words, ", "))
}
