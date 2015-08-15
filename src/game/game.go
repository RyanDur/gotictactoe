package game

import (
	gb "board"
	"errors"
)

type Game interface {
	Set(x, y int) error
	GetBoard() [][]string
	IsOver() bool
	Winner() string
}

type tictactoe struct {
	board  gb.Board
	winner string
}

func NewGame(b gb.Board) Game {
	return &tictactoe{b, ""}
}

func (t *tictactoe) Set(x, y int) error {
	if t.IsOver() {
		return errors.New("game is over")
	}
	if t.board.Get(x, y) != "" {
		return errors.New("spot taken")
	}

	piece := "O"
	if t.board.Entries()%2 == 0 {
		piece = "X"
	}

	t.board.Set(piece, x, y)
	if t.checkForWin() {
		t.winner = piece
	}
	return nil
}

func (t *tictactoe) GetBoard() [][]string {
	return t.board.Contents()
}

func (t *tictactoe) IsOver() bool {
	return t.Winner() != "" || t.board.IsFull()
}

func (t *tictactoe) Winner() string {
	return t.winner
}

func (t *tictactoe) checkForWin() bool {
	return t.forEach(func(row int) func(int) string {
		return func(col int) string {
			return t.board.Get(row, col)
		}
	}) || t.forEach(func(row int) func(int) string {
		return func(col int) string {
			return t.board.Get(col, row)
		}
	}) || t.groupContains(func(row int) string {
		return t.board.Get(row, row)
	}) || t.groupContains(func(row int) string {
		return t.board.Get(row, (t.board.Size()-1)-row)
	})
}

func (t *tictactoe) groupContains(move func(int) string) bool {
	size := t.board.Size()
	pieces := make(map[string]int)
	for i := 0; i < size; i++ {
		piece := move(i)
		if piece != "" {
			pieces[piece]++
		}
	}
	return contains(pieces, size)
}

func (t *tictactoe) forEach(move func(int) func(int) string) bool {
	for row := 0; row < t.board.Size(); row++ {
		if t.groupContains(move(row)) {
			return true
		}
	}
	return false
}

func contains(m map[string]int, value int) bool {
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}
