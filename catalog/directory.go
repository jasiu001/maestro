package catalog

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

const (
	filesDirectory = "./data"
)

type Directory struct{}

func (dr Directory) GetFiles() ([]string, error) {
	files, err := ioutil.ReadDir(filesDirectory)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading directory %q", filesDirectory)
	}

	var dataFiles []string
	for _, f := range files {
		dataFiles = append(dataFiles, f.Name())
	}

	return dataFiles, nil
}

func (dr Directory) GetFullPathToFile(file string) string {
	return fmt.Sprintf("%s/%s", filesDirectory, file)
}
