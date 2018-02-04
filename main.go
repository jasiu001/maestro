package main

import (
	"fmt"
	levenshtein "github.com/texttheater/golang-levenshtein/levenshtein"
)

func main() {
	fmt.Println(levenshtein.DistanceForStrings([]rune("aaaa"), []rune("aabb"), levenshtein.DefaultOptions))
}