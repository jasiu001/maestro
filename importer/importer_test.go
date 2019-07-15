package importer

import (
	"testing"

	"github.com/jasiu001/maestro/bucket"
	"github.com/jasiu001/maestro/importer/mocks"
	"github.com/stretchr/testify/assert"
	m "github.com/stretchr/testify/mock"
)

func TestBucketWriter_Save(t *testing.T) {
	// Given
	bs := &mocks.BuckerSaver{}
	bs.On("SaveBucket", m.AnythingOfType("bucket.Bundle")).Return(nil)

	bw := NewBucketWriter(bs)
	bw.SetType("o")
	bw.SetWords([]string{"word#1", "word#2"})
	bw.SetTranslation("trans#1")
	bw.SetDescription("desc#1")

	// Then
	assert.Equal(t, "o", bw.bct.Kind)
	assert.Equal(t, []string{"word#1", "word#2"}, bw.bct.Words)
	assert.Equal(t, "trans#1", bw.bct.Translation)
	assert.Equal(t, "desc#1", bw.bct.Description)

	assert.NoError(t, bw.Save())
	assert.Equal(t, bucket.Bundle{}, bw.bct)
}
