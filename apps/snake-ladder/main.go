package main

import (
	"math/rand"
)

type Dice struct{}

func (d Dice) Roll() int {
	return rand.Intn(6) + 1
}

type Pos int

type Player struct {
	Name string
	Pos  Pos
}

func (p Player) Move(pos Pos) {

}

type Board struct {
	Size            int
	SnakesPositions []Range
	LadderPositions []Range
}

func NewBoard(size int) *Board {
	return &Board{Size: size}
}

func (b *Board) AddSnakePositions() {

}

func (b *Board) AddLadderPositions() {

}

type SnakeAndLadder struct {
	Board   Board
	Dice    Dice
	Players map[string]Player
}

func (s *SnakeAndLadder) AddPlayers(player Player) {
	s.Players[player.Name] = player
}

func (s *SnakeAndLadder) StartGame() {

}

type Range struct {
	from int
	to   int
}

func main() {
	dice := Dice{}

	board := Board{
		Size:            50,
		SnakesPositions: []Range{{6, 10}},
		LadderPositions: []Range{{11, 49}},
	}

	app := SnakeAndLadder{
		Dice:  dice,
		Board: board,
	}

	app.AddPlayers(Player{Name: "x", Pos: 0})
	app.AddPlayers(Player{Name: "y", Pos: 0})
}
