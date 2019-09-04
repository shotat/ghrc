package main

import (
	"fmt"

	"bytes"
	"github.com/shotat/ghrc"
	"gopkg.in/yaml.v2"
)

const (
	filepath = ".ghrc.yaml"
)

func main() {
	err := export()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func apply() error {
	conf, err := ghrc.LoadRepositoryConfigFromFile(filepath)
	if err != nil {
		return err
	}

	return conf.Apply()
}

func export() error {
	meta := &ghrc.RepositoryMetadata{"shotat", "ghrc"}
	conf, err := ghrc.ImportConfig(meta)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	err = yaml.NewEncoder(buf).Encode(conf)
	fmt.Println(buf.String())
	return nil
}
