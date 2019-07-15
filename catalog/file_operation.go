package catalog

import "io/ioutil"

type FileReadWrite struct{}

func (f FileReadWrite) ReadFile(fileName string) ([]byte, error) {
	return ioutil.ReadFile(fileName)
}

func (f FileReadWrite) WriteFile(fileName string, content []byte) error {
	return ioutil.WriteFile(fileName, content, 0644)
}
