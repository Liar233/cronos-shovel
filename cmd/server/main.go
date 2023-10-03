package main

import (
	"fmt"
	"os"

	"github.com/Liar233/cronos-shovel/internal/server"
	"gopkg.in/yaml.v3"
)

func main() {
	config, err := loadConfig(os.Args[1])

	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	cronoServer := server.NewCronoServer(config)

	cronoServer.Run()
}

func loadConfig(filePath string) (*server.CronosServerConfig, error) {

	if _, err := os.Stat(filePath); err != nil {
		return nil, fmt.Errorf("config file `%s` does not exist", filePath)
	}

	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	var config server.CronosServerConfig

	if err = yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
