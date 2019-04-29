package comparison

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

var option = levenshtein.Options{
	InsCost: 1,
	DelCost: 1,
	SubCost: 1,
	Matches: func(sourceCharacter rune, targetCharacter rune) bool {
		return sourceCharacter == targetCharacter
	},
}

type Comparison struct {
	words []*Word
}

type Word struct {
	word           string
	bestResult     int
	bestWord       string
	wordsToCompare []ComparedWord
}

type ComparedWord struct {
	word   string
	result int
}

func NewComparison() *Comparison {
	return &Comparison{}
}

func (c *Comparison) Compare(patterns, words []string) int {
	c.makePairs(patterns, words)
	c.findBestResult()

	result := 0
	for _, pattern := range c.words {
		result += pattern.bestResult
	}

	return result
}

func (c *Comparison) makePairs(patterns, words []string) {
	for _, ptr := range patterns {
		w := Word{
			word:       ptr,
			bestResult: 9999,
		}
		for _, wrd := range words {
			comparedWord := ComparedWord{
				word:   wrd,
				result: levenshtein.DistanceForStrings([]rune(ptr), []rune(wrd), option),
			}
			w.wordsToCompare = append(w.wordsToCompare, comparedWord)
		}
		c.words = append(c.words, &w)
	}
}

func (c *Comparison) findBestResult() {
	for _, pattern := range c.words {
		for _, word := range pattern.wordsToCompare {
			if word.result < pattern.bestResult {
				pattern.bestResult = word.result
				pattern.bestWord = word.word
			}
		}
	}
}
