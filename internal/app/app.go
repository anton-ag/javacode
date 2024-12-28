package app

import (
	"fmt"
	"log"

	"github.com/anton-ag/javacode/internal/config"
	"github.com/anton-ag/javacode/pkg/postgres"
)

func Run(configPath string) {
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := postgres.Connect(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	// TODO: launch application
	fmt.Printf("%+v\n", cfg)
}
