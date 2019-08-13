package main

import (
	"avsanagrams/internal/app/avsanagrams"
	"avsanagrams/internal/app/avsanagrams/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	app := avsanagrams.New(cfg)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
