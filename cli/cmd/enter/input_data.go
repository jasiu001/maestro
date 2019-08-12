package enter

import (
	"fmt"
	"regexp"
)

const (
	labelType        = "Type"
	labelTranslation = "Translation"
	labelDescription = "Description"
)

type inputData struct {
	words []string
	index int
	items map[int]*Element
}

type Element struct {
	Label   string
	value   string
	isValid func(input string) (bool, string)
}

func NewInputData() *inputData {
	return &inputData{
		index: 0,
		items: map[int]*Element{
			0: {
				Label: labelType,
				isValid: func(input string) (b bool, s string) {
					types := []string{"v", "n", "a", "o"}
					for _, t := range types {
						if t == input {
							return true, ""
						}
					}
					return false, fmt.Sprintf("Available types: %v", types)
				},
			},
			1: {
				Label: labelTranslation,
				isValid: func(input string) (b bool, s string) {
					if input == "" {
						return false, "Translation is required"
					}
					re := regexp.MustCompile(`^[a-zA-Z\s\-,\:]*$`)
					if re.MatchString(input) {
						return true, ""
					}
					return false, "Translation has to contain only letters (and chars: -,:)"
				},
			},
			2: {
				Label: labelDescription,
				isValid: func(input string) (b bool, s string) {
					return true, ""
				},
			},
		},
	}
}

func (id *inputData) GetElement() (*Element, bool) {
	item, ok := id.items[id.index]
	if !ok {
		return &Element{}, false
	}
	id.index++

	return item, true
}

func (id *inputData) SetWord(input string) error {
	if input == "" {
		return fmt.Errorf("word cannot be empty")
	}

	re := regexp.MustCompile(`^[a-zA-Z\s]*$`)
	if re.MatchString(input) {
		id.words = append(id.words, input)
		return nil
	}

	return fmt.Errorf("word has to contain only letters")
}

func (e *Element) SetValue(input string) error {
	valid, msg := e.isValid(input)
	if valid {
		e.value = input
		return nil
	}

	return fmt.Errorf("Value %q is incorrect: %s", e.Label, msg)
}

func (id *inputData) GetType() string {
	return id.getValueByLabel(labelType)
}

func (id *inputData) GetWords() []string {
	return id.words
}

func (id *inputData) GetTranslation() string {
	return id.getValueByLabel(labelTranslation)
}

func (id *inputData) GetDescription() string {
	return id.getValueByLabel(labelDescription)
}

func (id *inputData) getValueByLabel(label string) string {
	for _, element := range id.items {
		if element.Label == label {
			return element.value
		}
	}

	return ""
}
