package bucket

import (
	"reflect"
	"testing"
)

func TestBucket_GetPairs(t *testing.T) {
	bucket := InitBucket([]string{"word1", "word2", "word3"}, []string{"uw1", "uw2"})
	testSlice := [][]string{
		{"word1", "uw1"},
		{"word1", "uw2"},
		{"word2", "uw1"},
		{"word2", "uw2"},
		{"word3", "uw1"},
		{"word3", "uw2"},
	}

	if len(bucket.GetPairs()) != 6 {
		t.Errorf("There should be 6 pairs but is %d", len(bucket.GetPairs()))
	}

	for _, testPair := range bucket.GetPairs() {
		if !contain(testPair, testSlice) {
			t.Errorf(
				"Pair buckets (%s - %s) should not exist in bucket",
				testPair.Word().Val(),
				testPair.UserWord().Val())
		}
	}
}

func TestBucket_GetPairsWithCorrectWords(t *testing.T) {
	bucket := InitBucket([]string{"testWord", "commonWord"}, []string{"commonWord", "testWord"})
	testSlice := [][]string{
		{"testWord", "commonWord"},
		{"testWord", "testWord"},
		{"commonWord", "commonWord"},
		{"commonWord", "testWord"},
	}

	if len(bucket.GetPairs()) != 4 {
		t.Errorf("There should be 4 pairs but is %d", len(bucket.GetPairs()))
	}

	for _, testPair := range bucket.GetPairs() {
		if !contain(testPair, testSlice) {
			t.Errorf(
				"Pair buckets (%s - %s) should not exist in bucket",
				testPair.Word().Val(),
				testPair.UserWord().Val())
		}
	}
}

func TestBucket_GetPairsWithOneUserWord(t *testing.T) {
	bucket := InitBucket([]string{"word1", "word2", "word3"}, []string{"test"})
	testSlice := [][]string{
		{"word1", "test"},
		{"word2", "test"},
		{"word3", "test"},
	}

	if len(bucket.GetPairs()) != 3 {
		t.Errorf("There should be 3 pairs but is %d", len(bucket.GetPairs()))
	}

	for _, testPair := range bucket.GetPairs() {
		if !contain(testPair, testSlice) {
			t.Errorf(
				"Pair buckets (%s - %s) should not exist in bucket",
				testPair.Word().Val(),
				testPair.UserWord().Val())
		}
	}
}

func contain(needle pair, test [][]string) bool {
	testPair := []string{
		needle.Word().Val(),
		needle.UserWord().Val(),
	}

	for _, pair := range test {
		if reflect.DeepEqual(pair, testPair) {
			return true
		}
	}

	return false
}
