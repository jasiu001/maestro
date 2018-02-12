package bucket

type UserWord string

func (w UserWord) Val() string {
	return string(w)
}
