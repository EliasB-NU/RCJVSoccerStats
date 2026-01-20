package web

import (
	"RCJV-SoccerStats/backend/data"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (a *API) getReferees(c *fiber.Ctx) error {
	var msg []string
	for _, r := range *a.Referees {
		msg = append(msg, r.Name)
	}
	return c.Status(fiber.StatusOK).JSON(msg)
}

func (a *API) getRefereeMatches(c *fiber.Ctx) error {
	// Check if its a valid referee
	var refereeParam = c.Params("referee")
	if refereeParam == "" {
		return c.Status(fiber.StatusBadRequest).JSON("Referee parameter is required")
	}

	var referee data.RefereeConverted
	for _, m := range *a.Referees {
		var convertedName = strings.ToLower(strings.ReplaceAll(m.Name, " ", ""))

		if refereeParam == convertedName {
			referee = m
		}
	}

	return c.Status(fiber.StatusOK).JSON(referee)
}
