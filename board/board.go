package board

import (
	"errors"
)

type Color int

const (
	Blank             Color = 0
	Red               Color = 1
	Yellow            Color = 2
	boardHeight             = 6
	boardWidth              = 7
	requiredSequalNum       = 4
)

var ErrColumnFull = errors.New("column is fulled")

type Board struct {
	State [][]Color
}

func NewBoard() *Board {
	// initialized with all
	newState := [][]Color{}
	for i := 0; i < boardHeight; i++ {
		row := make([]Color, boardWidth)
		for j := 0; j < boardWidth; j++ {
			row[j] = Blank
		}
		newState = append(newState, row)
	}

	return &Board{
		State: newState,
	}

}

func (b *Board) InsertColor(color Color, columnNum int) error {
	for i := boardHeight - 1; i >= 0; i-- {
		if b.State[i][columnNum] == Blank {
			b.State[i][columnNum] = color
			return nil
		}
	}

	return ErrColumnFull
}

func (b *Board) HasWinner() (bool, Color) {
	// some logic to check if
	// check is there 4 sequal of colors to 縦・横・斜
	if found, color := b.hasWinnerHorizontal(); found {
		return found, color
	}
	if found, color := b.hasWinnerVertical(); found {
		return found, color
	}
	if found, color := b.hasWinnerDiagonal(); found {
		return found, color
	}
	return false, Blank
}

func (b *Board) hasWinnerHorizontal() (bool, Color) {
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth-requiredSequalNum+1; j++ {
			color := b.State[i][j]
			if color == Blank {
				continue
			}
			if color == b.State[i][j+1] && color == b.State[i][j+2] && color == b.State[i][j+3] {
				return true, color
			}
		}
	}
	return false, Blank
}

func (b *Board) hasWinnerVertical() (bool, Color) {
	for i := 0; i < boardHeight-requiredSequalNum+1; i++ {
		for j := 0; j < boardWidth; j++ {
			color := b.State[i][j]
			if color == Blank {
				continue
			}

			if color == b.State[i+1][j] && color == b.State[i+2][j] && color == b.State[i+3][j] {
				return true, color
			}
		}
	}
	return false, Blank
}

func (b *Board) hasWinnerDiagonal() (bool, Color) {
	// left-top to right-down check
	for i := 0; i < boardHeight-requiredSequalNum+1; i++ {
		for j := 0; j < boardWidth-requiredSequalNum+1; j++ {
			color := b.State[i][j]
			if color == Blank {
				continue
			}
			if color == b.State[i+1][j+1] && color == b.State[i+2][j+2] && color == b.State[i+3][j+3] {
				return true, color
			}
		}
	}

	for i := 0; i < boardHeight-requiredSequalNum+1; i++ {
		for j := boardWidth - 1; j >= requiredSequalNum-1; j-- {
			color := b.State[i][j]
			if color == Blank {
				continue
			}
			if color == b.State[i+1][j-1] && color == b.State[i+2][j-2] && color == b.State[i+3][j-3] {
				return true, color
			}
		}
	}
	return false, Blank
}
