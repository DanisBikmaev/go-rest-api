package main

import (
	"flag"
	"log"

	"github.com/DanisBikmaev/go-rest-api/internal/app/apiserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "configPath", "config/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	s := apiserver.NewAPIServer(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
