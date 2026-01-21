package config

type Config struct {
	URL                    string
	TournamentAbbreviation string
	Host                   string
	Port                   int
}

func GetConfig() *Config {
	return &Config{
		URL:                    "https://catigoal.com/rest/v1/",
		TournamentAbbreviation: "test2",
		Host:                   "0.0.0.0",
		Port:                   3030,
	}
}
