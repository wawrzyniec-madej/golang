package game

import (
	"ball-game/board"
	"time"
)

type game struct {
	board *board.Board
}

func NewGame() *game {
	return &game{
		board: board.NewBoard(10, 10),
	}
}

func (game *game) Start() {
	for {
		game.board.Tick()
		time.Sleep(time.Millisecond * 123)
	}
}
