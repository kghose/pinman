package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const game_data_dir = "games"

type PinmanOrientation string

const (
	Up    PinmanOrientation = "UP"
	Horiz PinmanOrientation = "HORIZ"
	Vert  PinmanOrientation = "VERT"
)

type Pinman struct {
	row, col    int
	orientation PinmanOrientation
}

func (p *Pinman) move_left() {
	switch p.orientation {
	case Up:
		p.orientation = Horiz
		p.col -= 2
	case Horiz:
		p.orientation = Up
		p.col -= 1
	case Vert:
		p.col -= 1
	}
}

func (p *Pinman) move_right() {
	switch p.orientation {
	case Up:
		p.orientation = Horiz
		p.col += 1
	case Horiz:
		p.orientation = Up
		p.col += 2
	case Vert:
		p.col += 1
	}
}

func (p *Pinman) move_up() {
	switch p.orientation {
	case Up:
		p.orientation = Vert
		p.row -= 2
	case Horiz:
		p.row -= 1
	case Vert:
		p.orientation = Up
		p.row -= 1
	}
}

func (p *Pinman) move_down() {
	switch p.orientation {
	case Up:
		p.orientation = Vert
		p.row += 1
	case Horiz:
		p.row += 1
	case Vert:
		p.orientation = Up
		p.row += 2
	}
}

type PinmanGameBoard struct {
	board_no                  int
	board                     []string
	board_width, board_height int
}

func (g *PinmanGameBoard) load(board_no int) {
	g.board_no = board_no
	file_name := fmt.Sprintf("game_%02d.txt", g.board_no)
	f, err := os.Open(filepath.Join(game_data_dir, file_name))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	g.board_width = 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Clone(scanner.Text())
		g.board_width = max(g.board_width, len(line))
		g.board = append(g.board, line)
	}

	g.board_height = len(g.board)
}

func (b *PinmanGameBoard) get_pinman_position() (int, int) {
	for r := 0; r < b.board_height; r++ {
		c := strings.Index(b.board[r], "P")
		if c > -1 {
			return r, c
		}
	}
	panic("Can't find pinman on board! Pinman's start position has to be marked with a 'P' on the board.")
}

type PinmanGame struct {
	board PinmanGameBoard
	man   Pinman
}

func (g *PinmanGame) load(board_no int) {
	g.board.load(board_no)
	r, c := g.board.get_pinman_position()
	g.man.row = r
	g.man.col = c
	g.man.orientation = Up
}
