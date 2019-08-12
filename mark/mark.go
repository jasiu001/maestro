package mark

const (
	NONE    = 0
	CORRECT = 1
	PROPER  = 2
	SIMILAR = 3
	WRONG   = 4
)

const (
	NONE_STR    = "NONE"
	CORRECT_STR = "CORRECT"
	PROPER_STR  = "PROPER"
	SIMILAR_STR = "SIMILAR"
	WRONG_STR   = "WRONG"
)

func SpecifyMark(difference int) int {
	switch difference {
	case 0:
		return CORRECT
	case 1:
		return PROPER
	case 2:
		return SIMILAR
	default:
		return WRONG
	}
}

type Mark struct {
	value int
}

func InitMark() *Mark {
	return &Mark{value: NONE}
}

func (m Mark) NameMark() string {
	switch m.value {
	case 0:
		return NONE_STR
	case 1:
		return CORRECT_STR
	case 2:
		return PROPER_STR
	case 3:
		return SIMILAR_STR
	case 4:
		return WRONG_STR
	default:
		return "Undefined"
	}
}

func (m *Mark) UpdateMark(diff int) {
	markValue := SpecifyMark(diff)

	if m.value == NONE {
		m.value = markValue
		return
	}
	if markValue < m.value {
		m.value = markValue
	}
}

func (m Mark) Pass() bool {
	return m.value == CORRECT
}

func (m Mark) Fail() bool {
	return m.value != CORRECT
}
