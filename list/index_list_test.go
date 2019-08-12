package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexList_GenerateID(t *testing.T) {
	// Given
	il := NewIndexList()

	// When
	id := il.GenerateID()

	// Then
	assert.NotEmpty(t, id)
	assert.Len(t, id, 36)
	assert.IsType(t, "string", id)
}

func TestIndexList_RemoveID(t *testing.T) {
	for name, c := range map[string]struct {
		input          []string
		removedElement string
		expected       []string
	}{
		"last element": {
			input:          []string{"a", "b", "c"},
			removedElement: "c",
			expected:       []string{"a", "b"},
		},
		"first element": {
			input:          []string{"a", "b", "c"},
			removedElement: "a",
			expected:       []string{"b", "c"},
		},
		"middle element from 3 elements slice": {
			input:          []string{"a", "b", "c"},
			removedElement: "b",
			expected:       []string{"a", "c"},
		},
		"middle element from 5 elements slice": {
			input:          []string{"a", "b", "c", "d", "e"},
			removedElement: "d",
			expected:       []string{"a", "b", "c", "e"},
		},
	} {
		t.Run(name, func(t *testing.T) {
			// Given
			il := NewIndexList()
			il.list = c.input

			// When
			il.RemoveID(c.removedElement)

			// Then
			assert.Equal(t, c.expected, il.list)
		})
	}
}

func TestIndexList_RandomID_ZeroElements(t *testing.T) {
	// Given
	il := NewIndexList()
	il.GenerateID()
	il.GenerateID()

	// When
	for _, val := range il.list {
		il.RemoveID(val)
	}

	// Then
	assert.Equal(t, "", il.RandomID())
}

func TestIndexList_RandomID_OneElements(t *testing.T) {
	// Given
	il := NewIndexList()
	il.GenerateID()
	il.GenerateID()

	// When
	element := il.RandomID()
	il.RemoveID(element)

	// Then
	assert.NotEqual(t, element, il.RandomID())
}

func TestIndexList_RandomID_TwoElements(t *testing.T) {
	// Given
	il := NewIndexList()
	il.GenerateID()
	il.GenerateID()

	// When
	elementOne := il.RandomID()

	// Then
	elementTwo := il.RandomID()
	assert.NotEqual(t, elementOne, elementTwo)
}

func TestIndexList_RandomID_ThreeElements(t *testing.T) {
	// Given
	il := NewIndexList()
	il.GenerateID()
	il.GenerateID()
	il.GenerateID()

	// When
	elementOne := il.RandomID()

	// Then
	elementTwo := il.RandomID()
	assert.NotEqual(t, elementOne, elementTwo)

	// When
	elementThree := il.RandomID()
	assert.NotEqual(t, elementOne, elementThree)
	assert.NotEqual(t, elementTwo, elementThree)
}
