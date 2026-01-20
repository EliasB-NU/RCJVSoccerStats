package web

import "github.com/gofiber/fiber/v2"

func (a *API) getLeagues(c *fiber.Ctx) error {
	type Response struct {
		LeagueName string `json:"league_name"`
		Abbrev     string `json:"abbreviation"`
	}

	var msg []Response
	for _, m := range a.Leagues.Leagues {
		msg = append(msg, Response{
			LeagueName: m.LeagueName,
			Abbrev:     m.LeagueAbbreviation,
		})
	}

	return c.Status(fiber.StatusOK).JSON(msg)
}
