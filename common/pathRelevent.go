package common

import (
	"log"
	"os"
	"path/filepath"
)

func CurrentBasePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("get exec path error: %v", err)
	}
	return dir
}
