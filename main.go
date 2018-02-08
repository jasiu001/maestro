package main

import (
	"fmt"
	levenshtein "github.com/texttheater/golang-levenshtein/levenshtein"
)

// export GOPATH=$HOME/workspace/maestro-project

func main() {
	fmt.Println(levenshtein.DistanceForStrings([]rune("aaaa"), []rune("aabb"), levenshtein.DefaultOptions))
}
