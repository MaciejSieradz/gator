package main

import (
	"fmt"
	"log"

	"github.com/MaciejSieradz/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %v\n", cfg)

	err = cfg.SetUser("maciek")
	if err != nil {
		log.Fatalf("error setting userName: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config again: %v\n", cfg)
}
