package enter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInputData_GetElement(t *testing.T) {
	tid := NewInputData()

	element, exist := tid.GetElement()
	assert.True(t, exist)
	assert.Contains(t, element.Label, labelType)

	element, exist = tid.GetElement()
	assert.True(t, exist)
	assert.Contains(t, element.Label, labelTranslation)

	element, exist = tid.GetElement()
	assert.True(t, exist)
	assert.Contains(t, element.Label, labelDescription)

	_, exist = tid.GetElement()
	assert.False(t, exist)
}

func TestInputData_SetWord(t *testing.T) {
	for name, state := range map[string]struct {
		input      string
		err        bool
		errMessage string
	}{
		"no error": {
			input: "properWord",
			err:   false,
		},
		"no empty error": {
			input:      "",
			err:        true,
			errMessage: "word cannot be empty",
		},
		"wrong input error": {
			input:      "123fer",
			err:        true,
			errMessage: "word has to contain only letters",
		},
	} {
		t.Run(name, func(t *testing.T) {
			tid := NewInputData()

			err := tid.SetWord(state.input)
			if state.err {
				assert.Error(t, err)
				assert.EqualError(t, err, state.errMessage)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestElement_SetValue(t *testing.T) {
	for name, state := range map[string]struct {
		typeValue        string
		typeError        bool
		TranslationValue string
		TranslationError bool
		DescriptionValue string
	}{
		"no error": {
			typeValue:        "v",
			typeError:        false,
			TranslationValue: "some",
			TranslationError: false,
			DescriptionValue: "some desc",
		},
		"wrong type": {
			typeValue: "z",
			typeError: true,
		},
		"empty translation": {
			typeValue:        "v",
			typeError:        false,
			TranslationValue: "",
			TranslationError: true,
		},
		"wrong translation": {
			typeValue:        "v",
			typeError:        false,
			TranslationValue: "334ty",
			TranslationError: true,
		},
	} {
		t.Run(name, func(t *testing.T) {
			tid := NewInputData()

			element, _ := tid.GetElement()
			err := element.SetValue(state.typeValue)
			if state.typeError {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}

			element, _ = tid.GetElement()
			err = element.SetValue(state.TranslationValue)
			if state.TranslationError {
				assert.Error(t, err)
				return
			} else {
				assert.NoError(t, err)
			}

			element, _ = tid.GetElement()
			err = element.SetValue(state.DescriptionValue)
			assert.NoError(t, err)
		})
	}
}

func TestInputData_GetType(t *testing.T) {
	tid := NewInputData()

	element, _ := tid.GetElement()
	assert.NoError(t, element.SetValue("a"))

	assert.Equal(t, "a", tid.GetType())
}

func TestInputData_GetWords(t *testing.T) {
	tid := NewInputData()

	assert.NoError(t, tid.SetWord("word one"))
	assert.NoError(t, tid.SetWord("word two"))

	assert.Contains(t, tid.GetWords(), "word one")
	assert.Contains(t, tid.GetWords(), "word two")
}

func TestInputData_GetTranslation(t *testing.T) {
	tid := NewInputData()

	element, _ := tid.GetElement()
	assert.NoError(t, element.SetValue("a"))

	element, _ = tid.GetElement()
	assert.NoError(t, element.SetValue("translation input"))

	assert.Equal(t, "translation input", tid.GetTranslation())
}

func TestInputData_GetDescription(t *testing.T) {
	tid := NewInputData()

	element, _ := tid.GetElement()
	assert.NoError(t, element.SetValue("a"))

	element, _ = tid.GetElement()
	assert.NoError(t, element.SetValue("translation input"))

	element, _ = tid.GetElement()
	assert.NoError(t, element.SetValue("description input"))

	assert.Equal(t, "description input", tid.GetDescription())
}
