package main

type Player struct {
	Name      string
	Victories int
}

func NewPlayer(name string) Player {
	return Player{
		Name:      name,
		Victories: 0,
	}
}

type Score map[Player]int

func NewScore() Score {
	return make(map[Player]int)
}

type Game struct {
	Winner  Player
	Players []Player
	Score   Score
}

func NewGame(winner Player, score map[Player]int) Game {
	return Game{
		Winner: winner,
		Score:  score,
	}
}
