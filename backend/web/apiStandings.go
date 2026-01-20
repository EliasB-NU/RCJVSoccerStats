package web

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (a *API) getStandings(c *fiber.Ctx) error {
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

	// Get the standings for the league
	var msg StandingsMessage
	msg.League = a.Leagues.Leagues[leagueID].LeagueName
	msg.LeagueAbbrev = a.Leagues.Leagues[leagueID].LeagueAbbreviation
	msg.LastPublished = a.Leagues.Leagues[leagueID].LatestStandings.LastPublished
	msg.Standings = a.Leagues.Leagues[leagueID].LatestStandings.Standings

	return c.Status(fiber.StatusOK).JSON(msg)
}
