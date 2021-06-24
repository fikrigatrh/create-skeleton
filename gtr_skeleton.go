package gtr_skeleton

import (
	"log"
	"os"
)

func CreateTon(name string) {
	if name == "create_ton" {
		_, err := os.Stat("repo")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("repo", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		_, err = os.Stat("usecase")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("usecase", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		_, err = os.Stat("controllers")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("controllers", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		_, err = os.Stat("models")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("models", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		_, err = os.Stat("config")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("config", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		_, err = os.Stat("router")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("router", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		_, err = os.Stat("server")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("server", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}

		_, err = os.Stat("middlewares")

		if os.IsNotExist(err) {
			errDir := os.MkdirAll("middlewares", 0755)
			if errDir != nil {
				log.Fatal(err)
			}

		}
	}
}
