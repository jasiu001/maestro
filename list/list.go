package list

import (
	"fmt"
	"strings"
)

// #### MOCK
type List struct {
	Parts  []string
	result []string
}

func (l *List) IsFinished() bool {
	return len(l.Parts) == 0
}

func (l *List) GetDescription() string {
	part := l.Parts[0]
	l.Parts = l.Parts[1:]

	return part
}

func (l *List) NumberOfWords() int {
	return 2
}

func (l *List) ExecuteResponse(data []string) {
	for _, w := range data {
		l.result = append(l.result, w)
	}
}

func (l *List) GetResult() string {
	msg := fmt.Sprintf("you words: %s \n\n", strings.Join(l.result, ","))
	l.result = []string{}
	return msg
}

func (l *List) GetSummary() string {
	return "Here will be summary"
}
