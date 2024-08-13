package gamehdl

import (
	"gameapp/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gameService ports.GamesService
}

func NewGameHandler(gameService ports.GamesService) *GameHandler {
	return &GameHandler{gameService: gameService}
}

func (hdl *GameHandler) Get(c *gin.Context) {
	game, err := hdl.gameService.Get(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, game)
}

func (hdl *GameHandler) Create(c *gin.Context) {
	body := BodyCreate{}
	c.BindJSON(&body)

	game, err := hdl.gameService.Create(body.Name, body.Size, body.Bombs)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(201, BuildResponseCreate(game))
}

func (hdl *GameHandler) Reveal(c *gin.Context) {
	body := BodyRevealCell{}
	c.BindJSON(&body)

	game, err := hdl.gameService.Reveal(c.Param("id"), body.Row, body.Col)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, BuildResponseRevealCell(game))
}
