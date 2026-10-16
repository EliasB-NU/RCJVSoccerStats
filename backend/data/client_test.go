package data

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCatigoalRequests(t *testing.T) {
	responses := map[string]string{
		"/leagues":                              `{"tournament_name":"Test Cup","leagues":[{"name":"Open","abbrev":"open","league_stages":[{"name":"Final","standings_publishing_state":"PUBLISHED"}]}]}`,
		"/matches?league=open&league_stage=0":   `{"matches":[{"number":7,"league":"open"}]}`,
		"/standings?league=open&league_stage=0": `{"league_stage_name":"Final","standings":[{"points":3}]}`,
		"/standings?league=open":                `{"standings":[{"points":0}]}`,
		"/matches":                              `{"matches":[{"number":7,"league":"open","referees":[{"first_name":"Jane","last_name":"Doe"}]}]}`,
	}
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, ok := responses[r.RequestURI]
		if !ok || r.Method != http.MethodGet {
			t.Errorf("unexpected request: %s %s", r.Method, r.RequestURI)
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, body)
	}))
	defer server.Close()
	defer catigoalClient.CloseIdleConnections()
	url := server.URL + "/"
	leagues := GetLeagues(url)
	if leagues == nil || leagues.TournamentName != "Test Cup" || len(leagues.Leagues) != 1 {
		t.Fatalf("unexpected leagues: %+v", leagues)
	}
	leagues = GetMatches(url, leagues)
	matches := leagues.Leagues[0].Stages[0].Matches
	if len(matches) != 1 || matches[0].Number != 7 {
		t.Fatalf("unexpected matches: %+v", matches)
	}
	leagues = GetStandings(url, leagues)
	standings := leagues.Leagues[0].LatestStandings
	if standings.LeagueStageName != "Final" || len(standings.Standings) != 1 || standings.Standings[0].Points != 3 {
		t.Fatalf("unexpected standings: %+v", standings)
	}
	leagues.Leagues[0].Stages[0].StandingsPublished = "UNPUBLISHED"
	leagues = GetStandings(url, leagues)
	seeding := leagues.Leagues[0].LatestStandings
	if seeding.LeagueStageName != "Seeding" || len(seeding.Standings) != 1 {
		t.Fatalf("unexpected seeding: %+v", seeding)
	}
	referees := GetReferees(url)
	if referees == nil || len(*referees) != 1 || (*referees)[0].Name != "Jane Doe" || len((*referees)[0].Matches) != 1 {
		t.Fatalf("unexpected referees: %+v", referees)
	}
}

func TestCatigoalRequestFailures(t *testing.T) {
	for _, tt := range []struct {
		name   string
		status int
		body   string
		closed bool
	}{
		{"HTTP error", http.StatusBadGateway, `{}`, false},
		{"invalid JSON", http.StatusOK, `{`, false},
		{"connection error", http.StatusOK, `{}`, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				fmt.Fprint(w, tt.body)
			}))
			defer server.Close()
			defer catigoalClient.CloseIdleConnections()
			if tt.closed {
				server.Close()
			}
			url := server.URL + "/"
			if got := GetLeagues(url); got != nil {
				t.Errorf("GetLeagues = %+v, want nil", got)
			}
			if got := getMatchesLeague(url, "open", 0); got != nil {
				t.Errorf("getMatchesLeague = %+v, want nil", got)
			}
			if got := GetReferees(url); got != nil {
				t.Errorf("GetReferees = %+v, want nil", got)
			}
			if got := getStandingsLeague(url, "open", 0); got.Standings != nil || got.LeagueStageName != "" {
				t.Errorf("getStandingsLeague = %+v, want empty response", got)
			}
			if got := getSeedingLeague(url, "open"); got.Standings != nil || got.LeagueStageName != "" {
				t.Errorf("getSeedingLeague = %+v, want empty response", got)
			}
		})
	}
}
