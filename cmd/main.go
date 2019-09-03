package main

import (
	"fmt"

	"github.com/shotat/ghrc"
)

const (
	filepath = "./example/sample.yaml"

// filepath = ".ghrc.yaml"
)

func main() {
	err := func() error {
		conf, err := ghrc.LoadRepositoryConfigFromFile(filepath)
		if err != nil {
			return err
		}

		return conf.Apply()
	}()

	if err != nil {
		fmt.Println(err.Error())
	}
}
