package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLastIndexList_add(t *testing.T) {
	// Given
	list := &LastIndexList{}

	// When
	list.add("ef1a95b8-cf30-4f18-9565-ecc1099bfba2")
	list.add("1d096c33-1a27-4af0-9d4e-7a9fef3184fc")

	// Then
	assert.Equal(t, "1d096c33-1a27-4af0-9d4e-7a9fef3184fc", list.newer)
	assert.Equal(t, "ef1a95b8-cf30-4f18-9565-ecc1099bfba2", list.older)

	// When
	list.add("")

	// Then
	assert.Equal(t, "", list.newer)
	assert.Equal(t, "1d096c33-1a27-4af0-9d4e-7a9fef3184fc", list.older)
}

func TestLastIndexList_exist(t *testing.T) {
	// Given
	list := &LastIndexList{}

	// When
	list.add("6e04132b-c400-47b8-8a24-5bcfb1c0a282")
	list.add("00abb473-8768-48fa-a7ba-2475bbdf8516")

	// Then
	assert.True(t, list.exist("6e04132b-c400-47b8-8a24-5bcfb1c0a282"))
	assert.True(t, list.exist("00abb473-8768-48fa-a7ba-2475bbdf8516"))

	// When
	list.add("")

	// Then
	assert.False(t, list.exist("6e04132b-c400-47b8-8a24-5bcfb1c0a282"))
	assert.True(t, list.exist("00abb473-8768-48fa-a7ba-2475bbdf8516"))

	// When
	list.add("")

	// Then
	assert.False(t, list.exist("6e04132b-c400-47b8-8a24-5bcfb1c0a282"))
	assert.False(t, list.exist("00abb473-8768-48fa-a7ba-2475bbdf8516"))
}
