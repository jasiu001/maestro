package config

import (
	"fmt"
	"log"
	"path"
	"runtime"
)

func GetFile(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("Runtime cannot read current file")
	}

	return fmt.Sprintf("%s/%s", path.Join(path.Dir(filename)), name)
}
