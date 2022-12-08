package models

type Response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type UpdateTournament struct {
	TournamentId uint `json:"tournamentId"`
	TeamCount    int  `json:"teamCount"`
}

type Total struct {
	TournamentTotal int `json:"tournamentTotal"`
	TeamTotal       int `json:"teamTotal"`
}
