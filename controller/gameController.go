package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pejuang-awan/BE-Tournament-Manager/models"
	"github.com/pejuang-awan/BE-Tournament-Manager/service"
)

func GetGameById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response := models.Response{
			Data:  nil,
			Error: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	game := service.GetGameById(uint(id))

	response := models.Response{
		Data:  game,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}

func GetGames(c *gin.Context) {
	games := service.GetGames()

	response := models.Response{
		Data:  games,
		Error: nil,
	}

	c.JSON(http.StatusOK, response)
}
