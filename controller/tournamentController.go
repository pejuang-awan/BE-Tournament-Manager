package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pejuang-awan/BE-Tournament-Manager/models"
	"github.com/pejuang-awan/BE-Tournament-Manager/service"
)

func Index(c *gin.Context) {
	response := models.Response{
		Data:  "Welcome to Tournament Manager API",
		Error: nil,
	}
	c.JSON(http.StatusOK, response)
}

func CreateTournament(c *gin.Context) {
	var tournament models.Tournament

	if err := c.ShouldBindJSON(&tournament); err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tournament = service.CreateTournament(tournament)

	response := models.Response{
		Data:  tournament,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}

func GetTournaments(c *gin.Context) {
	tournaments := service.GetTournaments()
	response := models.Response{
		Data:  tournaments,
		Error: nil,
	}
	c.JSON(http.StatusOK, response)
}

func GetTournamentByGameId(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tournament := service.GetTournamentsByGameID(uint(id))

	response := models.Response{
		Data:  tournament,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}

func GetTournamentById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tournament := service.GetTournamentById(uint(id))

	response := models.Response{
		Data:  tournament,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}

func UpdateTournament(c *gin.Context) {
	var tournament models.Tournament

	if err := c.ShouldBindJSON(&tournament); err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tournament = service.UpdateTournament(tournament)

	response := models.Response{
		Data:  tournament,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}

func DeleteTournament(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	service.DeleteTournament(uint(id))
	response := models.Response{
		Data:  nil,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}

func GetTournamentsByGameID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tournaments := service.GetTournamentsByGameID(uint(id))
	response := models.Response{
		Data:  tournaments,
		Error: nil,
	}
	c.JSON(http.StatusOK, response)
}

func UpdateTournamentRegisteredTeam(c *gin.Context) {
	var updateTournament models.UpdateTournament

	if err := c.ShouldBindJSON(&updateTournament); err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tournament := service.UpdateTournamentRegisteredTeam(updateTournament.TournamentId, updateTournament.TeamCount)

	response := models.Response{
		Data:  tournament,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}

func GetTotalTournamentAndTeam(c *gin.Context) {
	gameId, err := strconv.ParseUint(c.Param("gameId"), 10, 64)
	if err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	tournamentTotal := service.GetTotalTournament(uint(gameId))
	teamTotal := service.GetTotalTeamByGameId(uint(gameId))

	total := models.Total{
		TournamentTotal: tournamentTotal,
		TeamTotal:       teamTotal,
	}

	response := models.Response{
		Data:  total,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}
