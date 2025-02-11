package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	source := "/home/felipe/Documentos/dev_lab/golang/quicknotes/cmd/http/main.go:95"

	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := filepath.Base(workdir)
	baseDirIndex := strings.Index(source, baseDir)

	fmt.Println(source[baseDirIndex:])
}
