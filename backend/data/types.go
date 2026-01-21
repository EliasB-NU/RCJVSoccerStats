package data

type League struct {
	TournamentName       string `json:"tournament_name"`
	TournamentAbbrev     string `json:"tournament_abbrev"`
	TournamentHeaderType string `json:"tournament_header_type"`
	Leagues              []struct {
		LeagueName           string `json:"name"`
		LeagueAbbreviation   string `json:"abbrev"`
		SeedingPublished     string `json:"seeding_publishing_state"`
		SeedingLastPublished string `json:"seeding_last_published"`
		Stages               []struct {
			Name                   string  `json:"name"`
			StandingsPublished     string  `json:"standings_publishing_state"`
			StandingsLastPublished string  `json:"standings_last_published"`
			Matches                []Match `json:"matches"`
		} `json:"league_stages"`
		LatestStandings StandingsResponse
	} `json:"leagues"`
}

type StandingsResponse struct {
	TournamentAbbrev     string      `json:"tournament_abbrev"`
	TournamentName       string      `json:"tournament_name"`
	TournamentHeaderType string      `json:"tournament_header_type"`
	LeagueAbbrev         string      `json:"league_abbrev"`
	LeagueName           string      `json:"league_name"`
	LeagueStageNumber    string      `json:"league_stage_number"`
	LeagueStageName      string      `json:"league_stage_name"`
	LastPublished        string      `json:"last_published"`
	Standings            []Standings `json:"standings"`
}

type Standings struct {
	Team          Team   `json:"team"`
	Group         int    `json:"group"`
	Wins          int    `json:"wins"`
	Draws         int    `json:"draws"`
	Losses        int    `json:"losses"`
	Points        int    `json:"points"`
	Scored        int    `json:"scored"`
	Conceded      int    `json:"conceded"`
	Difference    int    `json:"difference"`
	Rank          int    `json:"rank"`
	Qualification string `json:"qualification"`
	LastPublished string `json:"last_published"`
}

type Match struct {
	Number      int     `json:"number"`
	Team1       *Team   `json:"team1"`
	Team2       *Team   `json:"team2"`
	League      string  `json:"league"`
	LeagueStage *string `json:"league_stage"`
	GroupName   *string `json:"group_name"`
	Start       *string `json:"start"`
	Duration    *string `json:"duration"`
	Pitch       *string `json:"pitch"`
	Goals1      int     `json:"goals1"`
	Goals2      int     `json:"goals2"`
	Points1     int     `json:"points1"`
	Points2     int     `json:"points2"`
	Referees    []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"referees"`
	LastPublished string `json:"last_published"`
}

type MatchResponse struct {
	TournamentAbbrev     string  `json:"tournament_abbrev"`
	TournamentName       string  `json:"tournament_name"`
	TournamentHeaderType string  `json:"tournament_header_type"`
	LeagueAbbrev         string  `json:"league_abbrev"`
	LeagueName           string  `json:"league_name"`
	LeagueStageNumber    string  `json:"league_stage_number"`
	LeagueStageName      string  `json:"league_stage_name"`
	Matches              []Match `json:"matches"`
}

type Team struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Affiliation string  `json:"affiliation"`
	StartNumber *int    `json:"startnumber"`
	ExternalKey *string `json:"external_key"`
}

type RefereeConverted struct {
	Name    string `json:"name"`
	Matches []struct {
		Field         string `json:"field"`
		Team1         string `json:"team1"`
		Team2         string `json:"team2"`
		League        string `json:"league"`
		Start         string `json:"start"`
		SecondReferee string `json:"referees"`
	} `json:"matches"`
}
