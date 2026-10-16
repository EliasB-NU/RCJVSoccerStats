package data

import (
	"encoding/json"
	"fmt"
	"log"
)

func GetStandings(url string, leagues *League) *League {
	for i, l := range leagues.Leagues {
		var stageNumber int
		for i2, stage := range l.Stages {
			if stage.StandingsPublished == "PUBLISHED" {
				stageNumber = i2
			}
		}
		if stageNumber == 0 && l.Stages[0].StandingsPublished != "PUBLISHED" {
			leagues.Leagues[i].LatestStandings = getSeedingLeague(url, l.LeagueAbbreviation)
		} else {
			leagues.Leagues[i].LatestStandings = getStandingsLeague(url, l.LeagueAbbreviation, stageNumber)
		}
		// log.Println(leagues.Leagues[i].LatestStandings.Standings)
	}

	return leagues
}

func getStandingsLeague(url string, leagueAbbrev string, stageNumber int) StandingsResponse {
	response, err := catigoalClient.Get(fmt.Sprintf("%sstandings?league=%s&league_stage=%d", url, leagueAbbrev, stageNumber))
	if err != nil {
		log.Printf("Error fetching Catigoal data: %v\n", err)
		return StandingsResponse{}
	}
	defer response.Close()
	statusCode, body := response.StatusCode(), response.Body()
	if statusCode != 200 {
		log.Printf("Error getting standings with code %d\n", statusCode)
		return StandingsResponse{}
	}
	var standings StandingsResponse
	err = json.Unmarshal(body, &standings)
	if err != nil {
		log.Printf("Error unmarshalling standings: %v\n", err)
		return StandingsResponse{}
	}
	return standings
}

func getSeedingLeague(url string, leagueAbbrev string) StandingsResponse {
	response, err := catigoalClient.Get(fmt.Sprintf("%sstandings?league=%s", url, leagueAbbrev))
	if err != nil {
		log.Printf("Error fetching Catigoal data: %v\n", err)
		return StandingsResponse{}
	}
	defer response.Close()
	statusCode, body := response.StatusCode(), response.Body()
	if statusCode != 200 {
		log.Printf("Error getting standings with code %d\n", statusCode)
		return StandingsResponse{}
	}
	var standings StandingsResponse
	err = json.Unmarshal(body, &standings)
	if err != nil {
		log.Printf("Error unmarshalling standings: %v\n", err)
		return StandingsResponse{}
	}
	standings.LeagueStageName = "Seeding"
	return standings
}
