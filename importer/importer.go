package importer

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/jasiu001/maestro/bucket"
	"github.com/pkg/errors"
)

type bucketWriter struct {
	bct bucket.Bundle
}

func NewBucketWriter() *bucketWriter {
	return &bucketWriter{bucket.Bundle{}}
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
	spew.Dump(b.bct)

	return nil
}
