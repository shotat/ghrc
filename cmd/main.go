package main

import (
	"fmt"

	"github.com/shotat/ghrc"
)

func main() {
	err := func() error {
		conf, err := ghrc.LoadRepositoryConfigFromFile("./example/sample.yaml")
		if err != nil {
			return err
		}

		return conf.Apply()
	}()

	if err != nil {
		fmt.Println(err.Error())
	}
}
