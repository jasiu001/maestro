package list

import (
	"testing"

	"github.com/jasiu001/maestro/list/mocks"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestList_ExecuteResponse(t *testing.T) {
	// Given
	idList := mocks.IDList{}
	idList.On("GenerateID").Return("1234-abcd")
	idList.On("RandomID").Return("1234-abcd")
	idList.On("RemoveID", testMock.AnythingOfType("string")).Return()

	comparer := mocks.Comparer{}
	comparer.On("Compare", []string{"a", "b"}, []string{"a", "b"}).Return(0)
	comparer.On("Compare", []string{"a", "b"}, []string{"c", "b"}).Return(1)
	ls := CreateList(oneBucketTest(), &comparer, &idList)

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
	buckets := twoBucketsTest()

	idList := mocks.IDList{}
	idList.On("GenerateID").Return("1234-abcd").Once()
	idList.On("GenerateID").Return("5678-efgh").Once()
	idList.On("RandomID").Return("1234-abcd").Once()
	idList.On("RandomID").Return("5678-efgh")
	idList.On("RemoveID", testMock.AnythingOfType("string")).Return()

	comparer := mocks.Comparer{}
	comparer.On("Compare", []string{"a", "b"}, []string{"a", "b"}).Return(0)
	comparer.On("Compare", []string{"a", "b"}, []string{"c", "d"}).Return(0)
	comparer.On("Compare", []string{"c", "d"}, []string{"a", "b"}).Return(0)
	comparer.On("Compare", []string{"c", "d"}, []string{"c", "d"}).Return(0)

	// When
	ls := CreateList(buckets, &comparer, &idList)

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
