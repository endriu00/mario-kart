package main

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GameHandler struct {
	log *logrus.Logger
	db  DB
}

func NewGameHandler(log *logrus.Logger, db DB) *GameHandler {
	return &GameHandler{
		log: log,
		db:  db,
	}
}

func (h *GameHandler) GetLeaderboard(c *gin.Context) {
	players, err := h.db.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}

	// calculate the leaderboard
	sort.Slice(players, func(i, j int) bool {
		return players[i].Victories > players[j].Victories
	})

	c.JSON(http.StatusOK, players)
}

type AddGameReq struct {
	Score map[string]int `json:"score" binding:"required"`
}

func (h *GameHandler) AddGame(c *gin.Context) {
	var addGameReq AddGameReq
	if err := c.ShouldBindJSON(&addGameReq); err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	playersScoreFromReq := addGameReq.Score

	maxScore := 0
	score := NewScore()
	var winner Player
	for playerName, playerScore := range playersScoreFromReq {
		player, err := h.db.GetPlayer(playerName)
		if err != nil {
			// TODO: should probably differentiate an error with user not found
			// player didn't exist
			if err := h.db.AddPlayer(NewPlayer(playerName)); err != nil {
				h.log.WithFields(logrus.Fields{"player": player}).Error("failed to create player")
				c.JSON(http.StatusInternalServerError, "Internal Server Error")
				return
			}
			h.log.WithFields(logrus.Fields{"player": player}).Info("player created")
			player, _ = h.db.GetPlayer(playerName)
		}
		if playerScore > maxScore {
			maxScore = playerScore
			winner = *player
		}
		score[*player] = playerScore
	}
	h.log.WithFields(logrus.Fields{"winner": winner}).Info("determined winner")

	game := NewGame(winner, score)
	if err := h.db.AddGame(game); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if err := h.db.AddPlayerVictory(winner.Name); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusCreated, "Game Created")
}
