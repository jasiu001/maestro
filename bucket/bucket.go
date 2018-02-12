package bucket

type Bucket struct {
	words     []*Word
	uWords    []UserWord
	passWords int
}

func InitBucket(words []string, uWords []string) *Bucket {

	var bucketWords []*Word
	for _, word := range words {
		bucketWords = append(bucketWords, NewWord(word))
	}

	var bucketUserWords []UserWord
	for _, uWord := range uWords {
		bucketUserWords = append(bucketUserWords, UserWord(uWord))
	}

	return &Bucket{
		words:     bucketWords,
		uWords:    bucketUserWords,
		passWords: 0,
	}
}

func (b Bucket) GetPairs() []pair {
	var bucketPair []pair

	for _, word := range b.words {
		for _, uWord := range b.uWords {
			bucketPair = append(bucketPair, pair{word, uWord})
		}
	}

	return bucketPair
}

func (b Bucket) Words() []*Word {
	return b.words
}

func (b Bucket) Pass() bool {
	return len(b.words) == b.passWords
}

func (b *Bucket) PassWord() {
	b.passWords++
}
