package comparison

import (
	"github.com/jasiu001/maestro/bucket"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

var option = levenshtein.Options{
	InsCost: 1,
	DelCost: 1,
	SubCost: 1,
	Matches: func (sourceCharacter rune, targetCharacter rune) bool {
	return sourceCharacter == targetCharacter
	},
}

func CompareWords(word *bucket.Word, userWord bucket.UserWord) {
	diff := levenshtein.DistanceForStrings([]rune(word.GetValue()), []rune(string(userWord)), option)

	word.Rating().UpdateMark(diff)
}
