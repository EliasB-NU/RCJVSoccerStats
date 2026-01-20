package main

import (
	"RCJV-SoccerStats/backend/config"
	"RCJV-SoccerStats/backend/web"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Starting server...")

	cfg := config.GetConfig()
	url := fmt.Sprintf("%s%s/", cfg.URL, cfg.TournamentAbbreviation)

	// Start Webserver
	web.Init(cfg, url)
}
