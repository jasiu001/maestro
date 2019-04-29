package list

import (
	"github.com/jasiu001/maestro/list/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_ExecuteResponse(t *testing.T) {
	// Given
	comparer := mocks.Comparer{}
	comparer.On("Compare", []string{"a", "b"}, []string{"a", "b"}).Return(0)
	comparer.On("Compare", []string{"a", "b"}, []string{"c", "b"}).Return(1)
	ls := CreateList(oneBucketTest(), &comparer)

	// When
	ls.ExecuteResponse([]string{"c", "b"})

	// Then
	assert.False(t, ls.IsFinished())

	// When
	ls.ExecuteResponse([]string{"a", "b"})

	// Then
	assert.True(t, ls.IsFinished())
}

func TestFlow(t *testing.T) {
	// Given
	comparer := mocks.Comparer{}
	comparer.On("Compare", []string{"a", "b"}, []string{"a", "b"}).Return(0)
	comparer.On("Compare", []string{"a", "b"}, []string{"c", "d"}).Return(0)
	comparer.On("Compare", []string{"c", "d"}, []string{"a", "b"}).Return(0)
	comparer.On("Compare", []string{"c", "d"}, []string{"c", "d"}).Return(0)

	// When
	ls := CreateList(twoBucketsTest(), &comparer)

	// Then
	assert.False(t, ls.IsFinished())

	assert.NotEmpty(t, ls.GetDescription())
	assert.Equal(t, ls.NumberOfWords(), 2)
	ls.ExecuteResponse([]string{"a", "b"})
	assert.NotEmpty(t, ls.GetResult())

	assert.False(t, ls.IsFinished())

	assert.NotEmpty(t, ls.GetDescription())
	assert.Equal(t, ls.NumberOfWords(), 2)
	ls.ExecuteResponse([]string{"c", "d"})
	assert.NotEmpty(t, ls.GetResult())

	assert.True(t, ls.IsFinished())
	assert.NotEmpty(t, ls.GetSummary())
}

func oneBucketTest() []Bucket {
	return []Bucket{
		&BucketMock{
			transaltion: "trans#1",
			description: "desc#1",
			words:       []string{"a", "b"},
			wordsAmount: 2,
		},
	}
}

func twoBucketsTest() []Bucket {
	return []Bucket{
		&BucketMock{
			transaltion: "trans#1",
			description: "desc#1",
			words:       []string{"a", "b"},
			wordsAmount: 2,
		},
		&BucketMock{
			transaltion: "trans#2",
			description: "desc#2",
			words:       []string{"c", "d"},
			wordsAmount: 2,
		},
	}
}

type BucketMock struct {
	transaltion string
	description string
	words       []string
	wordsAmount int
}

func (b BucketMock) GetTranslation() string {
	return b.transaltion
}

func (b BucketMock) GetDescription() string {
	return b.description
}

func (b BucketMock) AmountOfWords() int {
	return b.wordsAmount
}

func (b BucketMock) GetWords() []string {
	return b.words
}
