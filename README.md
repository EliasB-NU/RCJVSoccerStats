# RCJV Soccer Stats

The RCJV Soccer Stats project is an open source web application designed
to show the upcoming matches and current standings from the [Catigoal](https://catigoal.com/)
application. 
There is a total of three different sites available:

```
Sites/
 ├─ https://example.com/#/referees               (Referee overview)
 ├─ https://example.com/#/matches?leagues=       (Match overview)
 └─ https://example.com/#/standings              (Standings overview)
```

The matches site have a league argument, where you can specify the 
leagues with the league abbreviation (e.g. `et,lwl,open`).

For the standings and matches site you can also specify the cycle time.
The parameter is called `time` and is specified in seconds (default: 20).

Matches and Standings sites cycle through the different leagues and stages
every 20 seconds, that's also the interval new data gets fetched from the backend, which fetches the data from Catigoal every 10 seconds.

## Development
To run the project locally, you need to have [Node.js](https://nodejs.org/) and [Golang](https://go.dev) installed.
Then, follow these steps:
```bash
# First build the frontend
cd frontend
npm install
npm run build

# Now start the backend
go run ./backend/main.go (TOURNAMENT_ABBREVIATION)

# You can now access all the sites, e.g. http://localhost:3030/#/standings
# If you change something in the frontend code, you need to rebuild it
cd frontend
npm run build
```

If you have any questions, feel free to open an issue or contact me directly.

## Deployment
You can either run it directly or use the docker setup provided below.
```bash
docker run -d -p 3030:3030 ghcr.io/eliasb-nu/rcjvsoccerstats:latest --env TOURNAMENT_ABBREVIATION=(LeagueAbbreviation)
```