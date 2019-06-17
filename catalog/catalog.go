package catalog

import (
	"github.com/jasiu001/maestro/bucket"
	"github.com/pkg/errors"
)

const maxBucketsPerFile = 30

type FileManager interface {
	GetFileContent() ([]bucket.Bundle, error)
	GenerateNewName()
	SaveContent([]bucket.Bundle) error
}

type CatalogManager struct {
	file FileManager
}

func NewCatalogManager(fm FileManager) *CatalogManager {
	return &CatalogManager{
		file: fm,
	}
}

func (cm *CatalogManager) SaveBucket(bct bucket.Bundle) error {
	buckets, err := cm.file.GetFileContent()
	if err != nil {
		return errors.Wrap(err, "during fetch last file content")
	}

	if len(buckets) >= maxBucketsPerFile {
		cm.file.GenerateNewName()
	}

	buckets = append(buckets, bct)
	return cm.file.SaveContent(buckets)
}
