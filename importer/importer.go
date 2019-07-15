package importer

import (
	"github.com/google/uuid"
	"github.com/jasiu001/maestro/bucket"
	"github.com/pkg/errors"
)

// mockery -name BuckerSaver
type BuckerSaver interface {
	SaveBucket(b bucket.Bundle) error
}

type bucketWriter struct {
	saver BuckerSaver
	bct   bucket.Bundle
}

func NewBucketWriter(bs BuckerSaver) *bucketWriter {
	return &bucketWriter{
		saver: bs,
		bct:   bucket.Bundle{},
	}
}

func (b *bucketWriter) SetType(t string) {
	b.bct.Kind = t
}

func (b *bucketWriter) SetWords(words []string) {
	b.bct.Words = words
}

func (b *bucketWriter) SetTranslation(translation string) {
	b.bct.Translation = translation
}

func (b *bucketWriter) SetDescription(desc string) {
	b.bct.Description = desc
}

func (b *bucketWriter) Save() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return errors.Wrap(err, "failed during create UUID for bucket")
	}

	b.bct.ID = id.String()
	err = b.saver.SaveBucket(b.bct)
	if err != nil {
		return errors.Wrap(err, "failed during save bucket")
	}

	b.bct = bucket.Bundle{}
	return nil
}
