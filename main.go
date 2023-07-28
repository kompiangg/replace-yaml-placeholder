package main

import (
	"log"
	"os"
)

func main() {
	cfg := loadConfig()

	buffer, err := cfg.loadConfigFromYaml()
	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("config.yaml", buffer, 0644)
}
