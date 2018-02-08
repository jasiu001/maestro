package comparison

import (
	"testing"
	"github.com/jasiu001/maestro/bucket"
)

func TestCompareWordsWithResultCorrect(t *testing.T) {
	word := bucket.NewWord("ccccc")
	uWord := bucket.UserWord("ccccc")

	CompareWords(word, uWord)
	if word.Rating().NameMark() != "CORRECT" {
		t.Errorf("Word mark should be CORRECT but is: %s", word.Rating().NameMark())
	}
}

func TestCompareWordsWithResultProper(t *testing.T) {
	word := bucket.NewWord("aaaaa")
	uWord := bucket.UserWord("aaaab")

	CompareWords(word, uWord)
	if word.Rating().NameMark() != "PROPER" {
		t.Errorf("Word mark should be PROPER but is: %s", word.Rating().NameMark())
	}
}

func TestCompareWordsWithResultSimilar(t *testing.T) {
	word := bucket.NewWord("xxxxxxx")
	uWord := bucket.UserWord("xxyxxxy")

	CompareWords(word, uWord)
	if word.Rating().NameMark() != "SIMILAR" {
		t.Errorf("Word mark should be SIMILAR but is: %s", word.Rating().NameMark())
	}
}

func TestCompareWordsWithResultWrong(t *testing.T) {
	word := bucket.NewWord("aaaaaa")
	uWord := bucket.UserWord("bbbbbb")

	CompareWords(word, uWord)
	if word.Rating().NameMark() != "WRONG" {
		t.Errorf("Word mark should be WRONG but is: %s", word.Rating().NameMark())
	}
}

func TestCompareWordsWithResultCorrectAfterSecondTry(t *testing.T) {
	word := bucket.NewWord("ccccc")
	uWord := bucket.UserWord("ccbbc")

	CompareWords(word, uWord)
	if word.Rating().NameMark() != "SIMILAR" {
		t.Errorf("Word mark should be SIMILAR but is: %s", word.Rating().NameMark())
	}

	uWord2 := bucket.UserWord("ccccc")
	CompareWords(word, uWord2)
	if word.Rating().NameMark() != "CORRECT" {
		t.Errorf("Word mark should be CORRECT but is: %s", word.Rating().NameMark())
	}
}
