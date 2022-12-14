package service

import (
	"github.com/pejuang-awan/BE-Tournament-Manager/models"
	"github.com/pejuang-awan/BE-Tournament-Manager/utils"
)

func GetTournaments() []models.Tournament {
	var tournaments []models.Tournament
	utils.Database.Find(&tournaments)
	return tournaments
}

func GetTournamentById(id uint) models.Tournament {
	var tournament models.Tournament
	utils.Database.Find(&tournament, id)
	return tournament
}

func CreateTournament(tournament models.Tournament) models.Tournament {
	utils.Database.Create(&tournament)
	game := GetGameById(tournament.GameID)
	game.RegisteredTournament++
	UpdateGame(game)
	return tournament
}

func UpdateTournament(tournament models.Tournament) models.Tournament {
	utils.Database.Save(&tournament)
	return tournament
}

func DeleteTournament(tid uint) {
	var tournament models.Tournament
	utils.Database.Where("id = ?", tid).Delete(&tournament)
}

func GetTournamentsByGameID(gameID uint) []models.Tournament {
	var tournaments []models.Tournament
	utils.Database.Where("game_id = ?", gameID).Find(&tournaments)
	return tournaments
}

func GetTournamentsByGameName(gameName string) []models.Tournament {
	var tournaments []models.Tournament
	utils.Database.Where("game_id = ?", gameName).Find(&tournaments)
	return tournaments
}

func UpdateTournamentRegisteredTeam(tid uint, team int) models.Tournament {
	tournament := GetTournamentById(tid)
	tournament.RegisteredTeam = team
	UpdateTournament(tournament)
	return tournament
}

func GetTotalTeamByGameId(gid uint) int {
	tournaments := GetTournamentsByGameID(gid)
	totalTeam := 0
	for _, tournament := range tournaments {
		totalTeam += tournament.RegisteredTeam
	}

	return totalTeam
}
