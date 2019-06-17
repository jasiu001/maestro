package catalog

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/jasiu001/maestro/bucket"
	"github.com/pkg/errors"
)

const datetimeFormatForFiles = "20060102150405"

type DirectoryManager interface {
	GetFiles() ([]string, error)
	GetFullPathToFile(string) string
}

type FileOperation interface {
	ReadFile(string) ([]byte, error)
	WriteFile(string, []byte) error
}

type file struct {
	name      string
	frw       FileOperation
	directory DirectoryManager
}

func NewFile(name string, f FileOperation, dir DirectoryManager) *file {
	return &file{
		name:      name,
		frw:       f,
		directory: dir,
	}
}

func (f *file) GetFileContent() ([]bucket.Bundle, error) {
	err := f.findFilePath()
	if err != nil {
		return nil, errors.Wrap(err, "while finding path file")
	}

	currentFile := f.directory.GetFullPathToFile(f.name)
	content, err := f.frw.ReadFile(currentFile)
	if err != nil {
		return nil, errors.Wrapf(err, "while reading file %q content", currentFile)
	}

	if len(content) == 0 {
		return []bucket.Bundle{}, nil
	}

	var bc []bucket.Bundle
	err = json.Unmarshal(content, &bc)
	if err != nil {
		return nil, errors.Wrapf(err, "while unmarshaling file %q content", currentFile)
	}

	return bc, nil
}

func (f *file) GenerateNewName() {
	currentTime := time.Now()

	f.name = fmt.Sprintf("%s.json", currentTime.Format(datetimeFormatForFiles))
}

func (f *file) SaveContent(buckets []bucket.Bundle) error {
	content, err := json.MarshalIndent(buckets, "", "  ")
	if err != nil {
		return errors.Wrap(err, "while marshaling buckets")
	}

	err = f.frw.WriteFile(f.directory.GetFullPathToFile(f.name), content)
	if err != nil {
		return errors.Wrapf(err, "while saving content in %q file", f.name)
	}

	return nil
}

func (f *file) findFilePath() error {
	files, err := f.directory.GetFiles()
	if err != nil {
		return errors.Wrap(err, "while reading files from directory")
	}

	if len(files) == 0 {
		f.GenerateNewName()
		return nil
	}

	if f.name == "" {
		sort.Strings(files)
		lastFile := files[len(files)-1]
		f.name = lastFile
		return nil
	}

	for _, file := range files {
		if f.name == file {
			return nil
		}
	}

	return fmt.Errorf("file %s does not exist", f.name)
}
