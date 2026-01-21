package data

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetMatches(url string, leagues *League) *League {
	for i, l := range leagues.Leagues {
		for stageIndex := range l.Stages {
			leagues.Leagues[i].Stages[stageIndex].Matches = getMatchesLeague(url, l.LeagueAbbreviation, stageIndex)
		}
	}
	return leagues
}

func getMatchesLeague(url string, leagueAbbrev string, stageIndex int) []Match {
	agent := fiber.Get(fmt.Sprintf("%smatches?league=%s&league_stage=%d", url, leagueAbbrev, stageIndex)).InsecureSkipVerify()
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 || statusCode != 200 {
		log.Printf("Error getting matches with code %d and error: %v\n", statusCode, errs)
		return nil
	}
	var resp MatchResponse
	err := json.Unmarshal(body, &resp)
	if err != nil {
		log.Printf("Error unmarshalling matches: %v\n", err)
		return nil
	}
	return resp.Matches
}
