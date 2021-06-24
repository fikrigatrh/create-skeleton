package gtr_skeleton

import (
	"log"
	"os"
)

func CreateTon(name string) {
	if name == "create_ton" {

		dirs := []string{
			"repo",
			"usecase",
			"controllers",
			"models",
			"config",
			"router",
			"server",
			"middlewares",
		}

		for _, dir := range dirs {
			createDir(dir)
		}
	}
}

func createDir(dir string) {
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
