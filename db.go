package main

import "errors"

type DB interface {
	AddPlayer(player Player) error
	AddPlayerVictory(name string) error
	GetPlayers() ([]*Player, error)
	GetPlayer(name string) (*Player, error)

	AddGame(game Game) error
	GetGames(limit int) ([]Game, error)
}

type InMemoryDB struct {
	Players []*Player
	Games   []Game
}

func (db *InMemoryDB) AddPlayer(player Player) error {
	for _, p := range db.Players {
		if player.Name == p.Name {
			return errors.New("player already exists")
		}
	}
	db.Players = append(db.Players, &player)
	return nil
}

func (db *InMemoryDB) AddPlayerVictory(name string) error {
	for _, player := range db.Players {
		if player.Name == name {
			player.Victories += 1
			return nil
		}
	}
	return errors.New("no player found")
}

func (db *InMemoryDB) AddGame(game Game) error {
	db.Games = append(db.Games, game)
	return nil
}

func (db *InMemoryDB) GetPlayers() ([]*Player, error) {
	return db.Players, nil
}

func (db *InMemoryDB) GetPlayer(name string) (*Player, error) {
	for _, player := range db.Players {
		if player.Name == name {
			return player, nil
		}
	}
	return &Player{}, errors.New("no player found")
}
func (db *InMemoryDB) GetGames(limit int) ([]Game, error) {
	return db.Games, nil
}

func NewInMemoryDB() DB {
	return &InMemoryDB{
		Players: make([]*Player, 0),
		Games:   make([]Game, 0),
	}
}

type PostgresDB struct{}
