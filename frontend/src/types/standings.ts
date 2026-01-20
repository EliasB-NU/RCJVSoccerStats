export interface LeagueResponse {
    league_name: string
    abbreviation: string
}

export interface Team {
    id: number
    name: string
    affiliation: string
    startnumber?: number | null
    external_key?: string | null
}

export interface StandingEntry {
    team: Team
    group: number
    wins: number
    draws: number
    losses: number
    points: number
    scored: number
    conceded: number
    difference: number
    rank: number
    qualification: string
    last_published: string
}

export interface StandingsMessage {
    league: string
    league_abbrev: string
    last_published: string
    stage_name: string
    standings: StandingEntry[]
}
