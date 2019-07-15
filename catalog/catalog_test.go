package catalog

import (
	"github.com/jasiu001/maestro/bucket"
	"testing"

	"github.com/jasiu001/maestro/catalog/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCatalogManager_SaveBucket(t *testing.T) {
	// Given
	fm := &mocks.FileManager{}
	fm.On("GetFileContent").Return([]bucket.Bundle{
		{
			ID:          "1234-abcd",
			Kind:        "o",
			Words:       nil,
			Translation: "trans#1",
			Description: "desc#1",
			Points:      0,
		},
		{
			ID:          "5678-efgh",
			Kind:        "o",
			Words:       nil,
			Translation: "trans#2",
			Description: "desc#2",
			Points:      0,
		},
	}, nil)
	fm.On("SaveContent", []bucket.Bundle{
		{
			ID:          "1234-abcd",
			Kind:        "o",
			Words:       nil,
			Translation: "trans#1",
			Description: "desc#1",
			Points:      0,
		},
		{
			ID:          "5678-efgh",
			Kind:        "o",
			Words:       nil,
			Translation: "trans#2",
			Description: "desc#2",
			Points:      0,
		},
		{
			ID:          "9012-ijkl",
			Kind:        "o",
			Words:       nil,
			Translation: "trans#3",
			Description: "desc#3",
			Points:      0,
		},
	}).Return(nil)

	cm := NewCatalogManager(fm)

	// Then
	assert.NoError(t, cm.SaveBucket(bucket.Bundle{
		ID:          "9012-ijkl",
		Kind:        "o",
		Words:       nil,
		Translation: "trans#3",
		Description: "desc#3",
		Points:      0,
	}))
}
