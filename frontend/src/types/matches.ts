export interface Team {
    id: number
    name: string
    affiliation: string
    startnumber?: number | null
    external_key?: string | null
}

export interface Match {
    number: number
    team1: Team | null
    team2: Team | null
    league: string
    league_stage?: string | null
    group_name?: string | null
    start: string
    duration?: string | null
    pitch?: string | null
    goals1: number
    goals2: number
    points1: number
    points2: number
}

export interface Stage {
    name: string
    standings_publishing_state: string
    standings_last_published: string
    matches: Match[]
    last_published: string
}

export type Stages = Stage[]

