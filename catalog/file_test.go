package catalog

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jasiu001/maestro/bucket"
	"github.com/jasiu001/maestro/catalog/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFile_GenerateNewName(t *testing.T) {
	// Given
	fo := mocks.FileOperation{}
	dir := mocks.DirectoryReader{}

	tc := NewFile("", &fo, &dir)

	// When
	tc.GenerateNewName()

	// Then
	assert.Equal(t, fmt.Sprintf("%s.json", time.Now().Format(datetimeFormatForFiles)), tc.name)
}

func TestFile_findFilePath(t *testing.T) {
	for name, state := range map[string]struct {
		input        string
		files        []string
		expected     string
		error        bool
		errorMsg     string
		directoryErr error
	}{
		"not empty input, no error": {
			input:    "20190611200712.json",
			files:    []string{"20190611200712.json", "20190611201012.json"},
			expected: "20190611200712.json",
		},
		"not empty input, error": {
			input:    "20190611200712.json",
			files:    []string{"20190611200715.json", "20190611201016.json"},
			error:    true,
			errorMsg: "file 20190611200712.json does not exist",
		},
		"not empty input, error from directory": {
			input:        "20190611200712.json",
			files:        nil,
			error:        true,
			directoryErr: errors.New("directory fatal error"),
			errorMsg:     "while reading files from directory: directory fatal error",
		},
		"empty input, empty files": {
			input:    "",
			files:    []string{},
			expected: fmt.Sprintf("%s.json", time.Now().Format(datetimeFormatForFiles)),
		},
		"empty input, no error, sort by time": {
			input: "",
			files: []string{
				"20190611200715.json",
				"20190611211016.json",
				"20190611201016.json",
			},
			expected: "20190611211016.json",
		},
		"empty input, no error, sort by date": {
			input: "",
			files: []string{
				"20190511200715.json",
				"20190611200715.json",
				"20190622200715.json",
				"20190616200715.json",
				"20190515200715.json",
				"20190615200715.json",
			},
			expected: "20190622200715.json",
		},
	} {
		t.Run(name, func(t *testing.T) {
			// Given
			fo := mocks.FileOperation{}
			fo.On("WriteFile", state.expected, []byte{}).Return(nil)

			dir := mocks.DirectoryReader{}
			dir.On("GetFiles").Return(state.files, state.directoryErr)
			dir.On("GetFullPathToFile", mock.AnythingOfType("string")).Return(state.expected, state.directoryErr)

			// When
			tc := NewFile(state.input, &fo, &dir)
			err := tc.findFilePath()

			// Then
			if state.error {
				assert.Error(t, err)
				assert.EqualError(t, err, state.errorMsg)
				return
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, state.expected, tc.name)
		})
	}
}

func TestFile_GetFileContent(t *testing.T) {
	for name, state := range map[string]struct {
		fileName       string
		content        string
		expected       []bucket.Bundle
		err            bool
		readFileErrMsg error
		expectedErrMsg string
	}{
		"no error": {
			fileName: "00001.json",
			content: `[{
				"id": "dd4b3bad-4671-4a66-99f3-608e6682a408",
				"type": "o",
				"words": [
					"eng word"
				],
				"translation": "trans",
				"description": "desc",
				"points": 0
			}]`,
			expected: []bucket.Bundle{
				{
					ID:          "dd4b3bad-4671-4a66-99f3-608e6682a408",
					Kind:        "o",
					Words:       []string{"eng word"},
					Translation: "trans",
					Description: "desc",
					Points:      0,
				},
			},
		},
		"no error, more than one items": {
			content: `[{
				"id": "dd4b3bad-4671-4a66-99f3-608e6682a408",
				"type": "o",
				"words": [
					"eng word"
				],
				"translation": "trans",
				"description": "desc",
				"points": 0
			},{
				"id": "ddbaa1b5-5541-47d6-940d-6a32889c1182",
				"type": "v",
				"words": [
					"spa word",
					"other word"
				],
				"translation": "trans#2",
				"description": "desc#2",
				"points": 5
			}]`,
			expected: []bucket.Bundle{
				{
					ID:          "dd4b3bad-4671-4a66-99f3-608e6682a408",
					Kind:        "o",
					Words:       []string{"eng word"},
					Translation: "trans",
					Description: "desc",
					Points:      0,
				},
				{
					ID:          "ddbaa1b5-5541-47d6-940d-6a32889c1182",
					Kind:        "v",
					Words:       []string{"spa word", "other word"},
					Translation: "trans#2",
					Description: "desc#2",
					Points:      5,
				},
			},
		},
		"empty content": {
			fileName: "00001.json",
			content:  "",
			expected: []bucket.Bundle{},
		},
		"read file error": {
			fileName:       "00002.json",
			err:            true,
			readFileErrMsg: errors.New("read error"),
			expectedErrMsg: `while reading file "00002.json" content: read error`,
		},
		"unmarshal error": {
			fileName:       "00001.json",
			content:        ".",
			err:            true,
			expectedErrMsg: `while unmarshaling file "00001.json" content: invalid character '.' looking for beginning of value`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			// Given
			fo := mocks.FileOperation{}
			fo.On("ReadFile", state.fileName).Return([]byte(state.content), state.readFileErrMsg)

			dir := mocks.DirectoryReader{}
			dir.On("GetFiles").Return([]string{state.fileName}, nil)
			dir.On("GetFullPathToFile", state.fileName).Return(state.fileName)

			// When
			tc := NewFile("", &fo, &dir)
			buckets, err := tc.GetFileContent()

			// Then
			if state.err {
				assert.Error(t, err)
				assert.EqualError(t, err, state.expectedErrMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, state.expected, buckets)
			}
		})
	}
}

func TestFile_SaveContent(t *testing.T) {
	// Given
	fo := mocks.FileOperation{}
	fo.On("WriteFile", "", []byte(`[
  {
    "id": "dd4b3bad-4671-4a66-99f3-608e6682a408",
    "type": "o",
    "words": [
      "eng word"
    ],
    "translation": "trans",
    "description": "desc",
    "points": 0
  }
]`)).Return(nil)

	dir := mocks.DirectoryReader{}
	dir.On("GetFullPathToFile", "").Return("")

	// When
	tc := NewFile("", &fo, &dir)
	err := tc.SaveContent([]bucket.Bundle{
		{
			ID:          "dd4b3bad-4671-4a66-99f3-608e6682a408",
			Kind:        "o",
			Words:       []string{"eng word"},
			Translation: "trans",
			Description: "desc",
			Points:      0,
		},
	})

	// Then
	assert.NoError(t, err)
}
