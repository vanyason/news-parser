package main

import (
	"fmt"

	"github.com/vanyason/news-parser/internal/config"
	"github.com/vanyason/news-parser/internal/log"
	"github.com/vanyason/news-parser/internal/scratcher"
)

func app() (err error) {
	log.Info("Scratcher App started")

	config, err := config.Parse("./config/")
	if err != nil {
		return fmt.Errorf("error while parsing config : %w", err.Error())
	}

	scratchers, err := scratcher.CreateScratchers(config)
	if err != nil {
		return fmt.Errorf("error while creating scratchers : %w", err.Error())
	}

	for _, s := range scratchers {
		_, _ = s.Scratch()
	}
	return nil

}

func main() {
	err := app()

	if err != nil {
		log.Fatal(err.Error())
	}
}
