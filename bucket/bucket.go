package bucket

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

type Bundle struct {
	Kind        string   `json:"type"`
	Words       []string `json:"words"`
	Translation string   `json:"translation"`
	Description string   `json:"description"`
	Points      int      `json:"points"`
}

func NewBundleCollectionFromFile(pathToFile string) ([]Bundle, error) {
	jsonFile, err := os.Open(pathToFile)
	if err != nil {
		return nil, errors.Wrap(err, "during open file")
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Wrap(err, "during read file")
	}

	var collection []Bundle
	err = json.Unmarshal(byteValue, &collection)
	if err != nil {
		return nil, errors.Wrap(err, "during unmarshal file content")
	}

	return collection, nil
}

func (b Bundle) GetTranslation() string {
	return b.Translation
}

func (b Bundle) GetDescription() string {
	return b.Description
}

func (b Bundle) GetWords() []string {
	return b.Words
}

func (b Bundle) AmountOfWords() int {
	return len(b.Words)
}
