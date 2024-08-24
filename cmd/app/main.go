package main

import (
	"flag"
	"log"

	"github.com/tonitaga/url-shortener/internal/app/server"
	"github.com/tonitaga/url-shortener/internal/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/config.yaml", "path to YAML config file")
}

func main() {
	flag.Parse()

	// Get configurations from file
	config := config.NewDefault()
	if err := config.LoadFromYaml(configPath); err != nil {
		log.Fatal(err)
	}

	// Creating and launching server
	server := server.New(config)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
