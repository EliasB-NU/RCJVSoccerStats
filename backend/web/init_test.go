package web

import (
	"RCJV-SoccerStats/backend/data"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestAppRoutes(t *testing.T) {
	var leagues data.League
	if err := json.Unmarshal([]byte(`{"leagues":[{"name":"Open","abbrev":"open","league_stages":[{"name":"Final","matches":[]}],"LatestStandings":{"league_stage_name":"Final","standings":[]}}]}`), &leagues); err != nil {
		t.Fatal(err)
	}
	referees := []data.RefereeConverted{{Name: "Jane Doe"}}
	frontendDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(frontendDir, "index.html"), []byte("<html>Soccer Stats</html>"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(frontendDir, "app.js"), []byte("console.log('soccer');"), 0600); err != nil {
		t.Fatal(err)
	}
	app := newApp(&API{Leagues: &leagues, Referees: &referees}, frontendDir)
	t.Cleanup(func() { _ = app.Shutdown() })

	for _, tt := range []struct {
		path   string
		status int
		body   string
		json   bool
	}{
		{"/healthcheck", 200, `"OK"`, true},
		{"/livez", 200, "OK", false},
		{"/readyz", 200, "OK", false},
		{"/api/v1/leagues", 200, `[{"league_name":"Open","abbreviation":"open"}]`, true},
		{"/api/v1/matches/open", 200, `[{"name":"Final","standings_publishing_state":"","standings_last_published":"","matches":[]}]`, true},
		{"/api/v1/standings/open", 200, `{"league":"Open","league_abbrev":"open","last_published":"","stage_name":"Final","standings":[]}`, true},
		{"/api/v1/referees", 200, `["Jane Doe"]`, true},
		{"/api/v1/referees/janedoe/matches", 200, `{"name":"Jane Doe","matches":null}`, true},
		{"/api/v1/matches/unknown", 400, `"Invalid league parameter"`, true},
		{"/api/v1/standings/unknown", 400, `"Invalid league parameter"`, true},
		{"/", 200, "<html>Soccer Stats</html>", false},
		{"/app.js", 200, "console.log('soccer');", false},
		{"/missing.js", 404, "", false},
		{"/api/v1/missing", 404, "", false},
	} {
		t.Run(tt.path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			req.Header.Set("Origin", "https://example.com")
			resp, err := app.Test(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != tt.status {
				t.Fatalf("status = %d, want %d; body = %s", resp.StatusCode, tt.status, body)
			}
			if got := resp.Header.Get("Access-Control-Allow-Origin"); got != "*" {
				t.Errorf("CORS origin = %q, want *", got)
			}
			if tt.json {
				var got, want any
				if err := json.Unmarshal(body, &got); err != nil {
					t.Fatal(err)
				}
				if err := json.Unmarshal([]byte(tt.body), &want); err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(got, want) {
					t.Errorf("body = %s, want %s", body, tt.body)
				}
			} else if tt.body != "" && string(body) != tt.body {
				t.Errorf("body = %q, want %q", body, tt.body)
			}
		})
	}

	t.Run("CORS preflight", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodOptions, "/api/v1/leagues", nil)
		req.Header.Set("Origin", "https://example.com")
		req.Header.Set("Access-Control-Request-Method", "GET")
		req.Header.Set("Access-Control-Request-Headers", "Content-Type")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusNoContent {
			t.Errorf("status = %d, want 204", resp.StatusCode)
		}
		for name, want := range map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET,POST,DELETE",
			"Access-Control-Allow-Headers": "Origin,Content-Type,Accept",
		} {
			if got := resp.Header.Get(name); strings.ReplaceAll(got, " ", "") != want {
				t.Errorf("%s = %q, want %q", name, got, want)
			}
		}
	})

	t.Run("static HEAD", func(t *testing.T) {
		resp, err := app.Test(httptest.NewRequest(http.MethodHead, "/app.js", nil))
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK || len(body) != 0 {
			t.Errorf("HEAD status = %d, body = %q", resp.StatusCode, body)
		}
	})
}
