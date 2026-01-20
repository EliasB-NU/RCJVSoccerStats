package web

import "RCJV-SoccerStats/backend/data"

type StandingsMessage struct {
	League        string           `json:"league"`
	LeagueAbbrev  string           `json:"league_abbrev"`
	LastPublished string           `json:"last_published"`
	StageName     string           `json:"stage_name"`
	Standings     []data.Standings `json:"standings"`
}
