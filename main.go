package main

import (
	"github.com/jasiu001/maestro/bucket"
	"github.com/jasiu001/maestro/cli"
	"github.com/jasiu001/maestro/comparison"
	"github.com/jasiu001/maestro/importer"
	"github.com/jasiu001/maestro/list"
	"log"
)

func main() {
	pathToFile := "./data/data.json"
	words, err := bucket.NewBundleCollectionFromFile(pathToFile)
	if err != nil {
		log.Fatalf("Failed during fetch data from file %q: %s", pathToFile, err)
	}

	var buckets []list.Bucket
	for _, word := range words {
		buckets = append(buckets, word)
	}
	wordsList := list.CreateList(buckets, comparison.NewComparison())
	cli.RunMaestro(wordsList, importer.NewBucketWriter())
}
