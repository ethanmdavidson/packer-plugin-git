package common

import (
	"log"
	"path/filepath"
)

func PrintOpeningRepo(path string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Printf("Finding repository from '%s' (unable to determine absolute path, %s)\n", path, err.Error())
	} else {
		log.Printf("Finding repository from '%s' (%s)\n", path, absPath)
	}
}
