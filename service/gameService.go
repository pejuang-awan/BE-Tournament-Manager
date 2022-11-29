package service

import (
	"github.com/pejuang-awan/BE-Tournament-Manager/models"
	"github.com/pejuang-awan/BE-Tournament-Manager/utils"
)

func GetGames() []models.Game {
	var games []models.Game
	utils.Database.Find(&games)
	return games
}

func GetGameById(id uint) models.Game {
	var game models.Game
	utils.Database.Find(&game, id)
	return game
}

func CreateGame(game []models.Game) []models.Game {
	utils.Database.Create(&game)
	return game
}
