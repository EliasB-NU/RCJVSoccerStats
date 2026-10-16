package data

import (
	"encoding/json"
	"fmt"
	"log"
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
	response, err := catigoalClient.Get(fmt.Sprintf("%smatches?league=%s&league_stage=%d", url, leagueAbbrev, stageIndex))
	if err != nil {
		log.Printf("Error fetching Catigoal data: %v\n", err)
		return nil
	}
	defer response.Close()
	statusCode, body := response.StatusCode(), response.Body()
	if statusCode != 200 {
		log.Printf("Error getting matches with code %d\n", statusCode)
		return nil
	}
	var resp MatchResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Printf("Error unmarshalling matches: %v\n", err)
		return nil
	}
	return resp.Matches
}
