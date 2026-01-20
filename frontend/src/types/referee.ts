export interface Match {
    field: string
    team1: string
    team2: string
    league: string
    start: string
    referees: string
}

export interface RefereeResponse {
    name: string
    matches: Match[]
}