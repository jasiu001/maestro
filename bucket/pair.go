package bucket

type pair struct {
	word  *Word
	uWord UserWord
}

func (p pair) Word() *Word {
	return p.word
}

func (p pair) UserWord() UserWord {
	return p.uWord
}
