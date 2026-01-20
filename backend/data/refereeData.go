package data

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetReferees(url string) *[]RefereeConverted {
	// Get all matches
	agent := fiber.Get(fmt.Sprintf("%smatches", url)).InsecureSkipVerify()
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 || statusCode != 200 {
		log.Printf("Error fetching refereeMatches status code %d with body %s\n", statusCode, body)
		return nil
	}
	// Unmarshal matches
	var resp MatchResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		log.Printf("Error unmarshalling refereeMatches: %v\n", err)
		return nil
	}

	// Convert referees to RefereeConverted
	return ConvertMatchesToReferees(resp.Matches)
}

func ConvertMatchesToReferees(matches []Match) *[]RefereeConverted {
	refMap := make(map[string]*RefereeConverted)

	for _, match := range matches {
		for i, ref := range match.Referees {
			refName := ref.FirstName + " " + ref.LastName

			// Create referee entry if it doesn't exist
			if _, exists := refMap[refName]; !exists {
				refMap[refName] = &RefereeConverted{
					Name: refName,
				}
			}

			// Determine second referee
			secondRef := ""
			if len(match.Referees) > 1 {
				if i == 1 {
					other := match.Referees[0]
					secondRef = other.FirstName + " " + other.LastName
				} else {
					other := match.Referees[1]
					secondRef = other.FirstName + " " + other.LastName
				}

			}

			// Safely extract fields
			field := ""
			if match.Pitch != nil {
				field = *match.Pitch
			}

			start := ""
			if match.Start != nil {
				start = *match.Start
			}

			team1 := ""
			if match.Team1 != nil {
				team1 = match.Team1.Name
			}

			team2 := ""
			if match.Team2 != nil {
				team2 = match.Team2.Name
			}

			// Append match to referee
			refMap[refName].Matches = append(
				refMap[refName].Matches,
				struct {
					Field         string `json:"field"`
					Team1         string `json:"team1"`
					Team2         string `json:"team2"`
					League        string `json:"league"`
					Start         string `json:"start"`
					SecondReferee string `json:"referees"`
				}{
					Field:         field,
					Team1:         team1,
					Team2:         team2,
					League:        match.League,
					Start:         start,
					SecondReferee: secondRef,
				},
			)
		}
	}

	// Convert map to slice
	result := make([]RefereeConverted, 0, len(refMap))
	for _, ref := range refMap {
		result = append(result, *ref)
	}

	return &result
}
