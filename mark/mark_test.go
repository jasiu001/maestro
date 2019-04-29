package mark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSpecifyMark(t *testing.T) {
	assert.Equal(t, CORRECT, SpecifyMark(0))
	assert.Equal(t, PROPER, SpecifyMark(1))
	assert.Equal(t, SIMILAR, SpecifyMark(2))
	assert.Equal(t, WRONG, SpecifyMark(3))
	assert.Equal(t, WRONG, SpecifyMark(5))
}

func TestMark_UpdateMarkToCorrectState(t *testing.T) {
	mark := InitMark()
	mark.UpdateMark(0)
	assert.Equal(t, CORRECT, mark.value)
	assert.True(t, mark.Pass())
	assert.False(t, mark.Fail())
}

func TestMark_UpdateMarkToProperState(t *testing.T) {
	mark := InitMark()
	mark.UpdateMark(1)
	assert.Equal(t, PROPER, mark.value)
	assert.False(t, mark.Pass())
	assert.True(t, mark.Fail())
}

func TestMark_UpdateMarkToSimilarState(t *testing.T) {
	mark := InitMark()
	mark.UpdateMark(2)
	assert.Equal(t, SIMILAR, mark.value)
	assert.False(t, mark.Pass())
	assert.True(t, mark.Fail())
}

func TestMark_UpdateMarkToWrongState(t *testing.T) {
	mark := InitMark()
	mark.UpdateMark(3)
	assert.Equal(t, WRONG, mark.value)
	assert.False(t, mark.Pass())
	assert.True(t, mark.Fail())
}

func TestMark_UpdateMarkWithManyMistakes(t *testing.T) {
	mark := InitMark()
	mark.UpdateMark(6)
	assert.Equal(t, WRONG, mark.value)
	assert.False(t, mark.Pass())
	assert.True(t, mark.Fail())
}

func TestMark_UpdateMarkFromSimilarToWrongState(t *testing.T) {
	mark := InitMark()
	mark.UpdateMark(2)
	mark.UpdateMark(6)
	assert.Equal(t, SIMILAR, mark.value)
	assert.False(t, mark.Pass())
	assert.True(t, mark.Fail())
}

func TestMark_UpdateMarkFromWrongToProperState(t *testing.T) {
	mark := InitMark()
	mark.UpdateMark(5)
	mark.UpdateMark(1)
	assert.Equal(t, PROPER, mark.value)
	assert.False(t, mark.Pass())
	assert.True(t, mark.Fail())
}
