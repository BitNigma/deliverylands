package main

import (
	"delivery/cmd/app"
	"delivery/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("can't get config", err)
	}
	err = app.Run(cfg)
	if err != nil {
		log.Fatal("can't run app")
	}
}
