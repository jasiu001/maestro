package comparison

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparison_Compare(t *testing.T) {
	testCases := map[string]struct {
		pattern []string
		words   []string
		result  int
	}{
		"full match one word": {
			pattern: []string{"aaa"},
			words:   []string{"aaa"},
			result:  0,
		},
		"full match two words": {
			pattern: []string{"aaa", "bbb"},
			words:   []string{"bbb", "aaa"},
			result:  0,
		},
		"full match three words": {
			pattern: []string{"aaa", "ccc", "bbb"},
			words:   []string{"ccc", "bbb", "aaa"},
			result:  0,
		},
		"one mistake one word": {
			pattern: []string{"aaa"},
			words:   []string{"aab"},
			result:  1,
		},
		"one mistake two words": {
			pattern: []string{"aaa", "bbb"},
			words:   []string{"bbb", "aab"},
			result:  1,
		},
		"one mistake three words": {
			pattern: []string{"aaa", "bbb", "ccc"},
			words:   []string{"ccc", "bbb", "aab"},
			result:  1,
		},
		"two mistakes one word": {
			pattern: []string{"aaa"},
			words:   []string{"acb"},
			result:  2,
		},
		"two mistakes two words": {
			pattern: []string{"aaa", "bbb"},
			words:   []string{"bcb", "aab"},
			result:  2,
		},
		"two mistakes three words": {
			pattern: []string{"aaa", "bbb", "ccc"},
			words:   []string{"bcb", "ccc", "aab"},
			result:  2,
		},
		"three mistakes one word": {
			pattern: []string{"aaa"},
			words:   []string{"bbb"},
			result:  3,
		},
		"three mistakes two words": {
			pattern: []string{"aaa", "bbb"},
			words:   []string{"bbc", "bba"},
			result:  3,
		},
		"three mistakes three words": {
			pattern: []string{"aaa", "bbb", "ccc"},
			words:   []string{"bbb", "add", "ccw"},
			result:  3,
		},
		"four mistakes one word": {
			pattern: []string{"aaaa"},
			words:   []string{"bbbb"},
			result:  4,
		},
		"four mistakes two words": {
			pattern: []string{"aaa", "bbb"},
			words:   []string{"byy", "xxa"},
			result:  4,
		},
		"four mistakes three words": {
			pattern: []string{"aaa", "bbb", "ccc"},
			words:   []string{"bcb", "add", "ccw"},
			result:  4,
		},
	}

	for name, tn := range testCases {
		t.Run(name, func(t *testing.T) {
			cmp := NewComparison()
			assert.Equal(t, tn.result, cmp.Compare(tn.pattern, tn.words))
		})
	}
}
