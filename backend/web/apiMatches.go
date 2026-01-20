package web

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (a *API) getMatches(c *fiber.Ctx) error {
	// Check if its a valid league
	var leagueParam = c.Params("league")
	var leagueValid = false
	var leagueID int
	for i, l := range a.Leagues.Leagues {
		if leagueParam == l.LeagueAbbreviation {
			leagueID = i
			leagueValid = true
		}
	}
	if !leagueValid {
		log.Printf("Invalid league parameter: %s\n", leagueParam)
		return c.Status(fiber.StatusBadRequest).JSON("Invalid league parameter")
	}

	return c.Status(fiber.StatusOK).JSON(a.Leagues.Leagues[leagueID].Stages)
}
