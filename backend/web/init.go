package web

import (
	"RCJV-SoccerStats/backend/config"
	"RCJV-SoccerStats/backend/data"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

type API struct {
	Leagues *data.League
}

func Init(cfg *config.Config, url string) {
	var (
		addr = fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
		a    = API{}

		rcjvSoccerStats = fiber.New(fiber.Config{
			ServerHeader: "rcjv_soccerStats:fiber",
			AppName:      "rcjv_soccerStats",
		})

		// Cors
		c = cors.New(cors.Config{
			AllowOrigins: strings.Join([]string{
				"*",
			}, ","),

			AllowHeaders: strings.Join([]string{
				"Origin",
				"Content-Type",
				"Accept",
			}, ","),

			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodDelete,
			}, ","),

			AllowCredentials: false,
		})
	)
	rcjvSoccerStats.Use(c)
	rcjvSoccerStats.Use(healthcheck.New(healthcheck.ConfigDefault))
	rcjvSoccerStats.Get("/healthcheck", getHealthCheck)

	// Get data
	// Get Data
	ticker := time.NewTicker(10 * time.Second)
	a.Leagues = data.GetLeagues(url)
	go func() {
		for {
			<-ticker.C
			a.Leagues = data.GetStandings(url, a.Leagues)
			a.Leagues = data.GetMatches(url, a.Leagues)

			//for _, l := range leagues.Leagues {
			//	log.Printf("Updated standings for league: %s\n", l.LeagueName)
			//	log.Println(l.LatestStandings.Standings)
			//}
			//for _, l := range leagues.Leagues {
			//	for _, stage := range l.Stages {
			//		log.Printf("Updated matches for league: %s, stage: %s\n", l.LeagueName, stage.Name)
			//		log.Println(stage.Matches)
			//	}
			//}
		}
	}()

	// API Endpoints
	apiV1 := fiber.New()
	rcjvSoccerStats.Mount("/api/v1", apiV1)
	apiV1.Get("/standings/:league", a.getStandings)
	apiV1.Get("/matches/:league", a.getMatches)

	// Website
	rcjvSoccerStats.Static("/", "./frontend/dist")

	// Start Server
	err := rcjvSoccerStats.Listen(addr)
	if err != nil {
		log.Fatalf("Error starting webserver: %v\n", err)
	}
}
