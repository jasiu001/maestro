package main

import (
	"log"

	"github.com/jasiu001/maestro/bucket"
	"github.com/jasiu001/maestro/catalog"
	"github.com/jasiu001/maestro/cli"
	"github.com/jasiu001/maestro/comparison"
	"github.com/jasiu001/maestro/importer"
	"github.com/jasiu001/maestro/list"
)

func main() {
	pathToFile := "./data/20190611200712.json"
	words, err := bucket.NewBundleCollectionFromFile(pathToFile)
	if err != nil {
		log.Fatalf("Failed during fetch data from file %q: %s", pathToFile, err)
	}

	var buckets []list.Bucket
	for _, word := range words {
		buckets = append(buckets, word)
	}
	wordsList := list.CreateList(buckets, comparison.NewComparison())
	cm := catalog.NewCatalogManager(catalog.NewFile("", catalog.FileReadWrite{}, catalog.Directory{}))
	cli.RunMaestro(wordsList, importer.NewBucketWriter(cm))
}
