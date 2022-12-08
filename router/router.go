package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pejuang-awan/BE-Tournament-Manager/controller"
)

func InitializeRouter() *gin.Engine {
	router := gin.Default()

	apiRoute := router.Group("/api")
	{
		apiRoute.GET("/", controller.Index)
		apiRoute.GET("/tournament/:id", controller.GetTournamentById)
		apiRoute.POST("/tournament", controller.CreateTournament)
		apiRoute.PUT("/tournament", controller.UpdateTournament)
		apiRoute.DELETE("/tournament/:id", controller.DeleteTournament)

		apiRoute.GET("/tournaments", controller.GetTournaments)
		apiRoute.GET("/tournaments/:id", controller.GetTournamentByGameId)

		apiRoute.POST("/update-count", controller.UpdateTournamentRegisteredTeam)
		apiRoute.GET("/get-count/:gameId", controller.GetTotalTournamentAndTeam)

		apiRoute.GET("/games", controller.GetGames)

		apiRoute.GET("/game/:id", controller.GetGameById)
		apiRoute.POST("/game", controller.CreateGame)
	}

	return router
}
