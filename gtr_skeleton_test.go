package gtr_skeleton

import (
	"os"
	"testing"
)

func TestMkdir(t *testing.T) {
	createDir("lorem")

	_, err := os.Stat("lorem")

	if os.IsNotExist(err) {
		t.Error("unable to find directory")
	} else {
		os.RemoveAll("lorem")
	}
}

func TestCreateTon(t *testing.T) {
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

	CreateTon("create_ton")

	for _, dir := range dirs {
		_, err := os.Stat(dir)

		if os.IsNotExist(err) {
			t.Error("unable to find directory", dir)
		} else {
			os.RemoveAll(dir)
		}
	}
}
