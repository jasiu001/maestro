package comparison

import (
	"github.com/jasiu001/maestro/bucket"
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

func CompareWords(word *bucket.Word, userWord bucket.UserWord) {
	diff := levenshtein.DistanceForStrings([]rune(word.Val()), []rune(userWord.Val()), option)

	word.Rating().UpdateMark(diff)
}

func RunComparison(bucket *bucket.Bucket) {

	for _, pair := range bucket.GetPairs() {
		CompareWords(pair.Word(), pair.UserWord())
	}

	for _, word := range bucket.Words() {
		if word.Rating().Pass() {
			bucket.PassWord()
		}
	}
}
