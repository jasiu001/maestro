package comparison

import (
	"github.com/jasiu001/maestro/bucket"
	"testing"
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

func TestBucket_RunComparison(t *testing.T) {
	testBucket := bucket.InitBucket([]string{"word1", "word2"}, []string{"word2", "word1"})
	RunComparison(testBucket)

	if !testBucket.Pass() {
		t.Errorf("Bucket should has state Pass but has not")
	}
}

func TestBucket_RunComparisonWithOneMistakeBucket(t *testing.T) {
	testBucket := bucket.InitBucket([]string{"word1", "word2"}, []string{"word", "word1"})
	RunComparison(testBucket)

	if testBucket.Pass() {
		t.Errorf("Bucket should not has state Pass but it has")
	}
}

func TestBucket_RunComparisonWithTwoMistakesBucket(t *testing.T) {
	testBucket := bucket.InitBucket([]string{"word1", "word2"}, []string{"word", "ward1"})
	RunComparison(testBucket)

	if testBucket.Pass() {
		t.Errorf("Bucket should not has state Pass but it has")
	}
}

func TestBucket_RunComparisonWithThreeCorrectWords(t *testing.T) {
	testBucket := bucket.InitBucket([]string{"word1", "word2", "word3"}, []string{"word2", "word3", "word1"})
	RunComparison(testBucket)

	if !testBucket.Pass() {
		t.Errorf("Bucket should has state Pass but has not")
	}
}
