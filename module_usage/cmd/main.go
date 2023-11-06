package main

import (
	"fmt"
	"os"

	"module_usage/internal/app"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func main() {
	configData, err := os.ReadFile("config/files.yml")
	if err != nil {
		log.Fatal(err)
	}
	var fileList struct {
		Files []app.File
	}

	err = yaml.Unmarshal(configData, &fileList)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileList.Files {
		contains, err := app.Search(file.Filename, file.Substring)
		if err != nil {
			log.Warn(err)
			continue
		}
		if contains {
			fmt.Printf("file %s contains %s: Yes\n", file.Filename, file.Substring)
		} else {
			fmt.Printf("file %s contains %s: No\n", file.Filename, file.Substring)
		}
	}

}
