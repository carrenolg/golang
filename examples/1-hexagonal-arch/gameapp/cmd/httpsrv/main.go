package main

import (
	"gameapp/internal/core/services/gamesrv"
	"gameapp/internal/handlers/gamehdl"
	"gameapp/internal/repositories/gamesrepo"
	"gameapp/pkg/uidgen"

	"github.com/gin-gonic/gin"
)

func main() {
	// init repos
	gameRepo := gamesrepo.NewMemKVS()
	// init services
	gameSrv := gamesrv.New(gameRepo, uidgen.New())
	// init handlers
	gamehdl := gamehdl.NewGameHandler(gameSrv)

	// init server
	router := gin.New()
	router.POST("/games", gamehdl.Create)
	router.GET("/games/:id", gamehdl.Get)
	router.PUT("/games/:id", gamehdl.Reveal)

	// run server
	router.Run(":8080")
}
