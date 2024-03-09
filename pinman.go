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

type PinmanGameBoardSquare string

const (
	Abyss PinmanGameBoardSquare = "ABYSS"
	Land  PinmanGameBoardSquare = "LAND"
	Exit  PinmanGameBoardSquare = "EXIT"
)

type PinmanGameBoard struct {
	board_no                  int
	board                     []string
	board_width, board_height int
	ideal_steps               int
}

func (g *PinmanGameBoard) load(board_no int) {
	g.board_no = board_no
	g.board = nil
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

func (g *PinmanGameBoard) square(r int, c int) PinmanGameBoardSquare {
	if r < 0 || r >= g.board_height {
		return Abyss
	}
	if c < 0 || c >= len(g.board[r]) {
		return Abyss
	}
	b := g.board[r][c]
	if b == '.' || b == 'P' {
		return Land
	}
	if b == 'X' {
		return Exit
	}
	return Abyss
}

type GameState string

const (
	GameOn   GameState = "GAMEON"
	FellDown GameState = "FELLDOWN"
	Exited   GameState = "EXITED"
)

type PinmanGame struct {
	board      PinmanGameBoard
	man        Pinman
	steps      int
	game_state GameState
}

func (g *PinmanGame) Load(board_no int) {
	g.game_state = GameOn
	g.steps = 0
	g.board.load(board_no)
	r, c := g.board.get_pinman_position()
	g.man.row = r
	g.man.col = c
	g.man.orientation = Up
}

type PinmanMove string

const (
	MoveLeft  PinmanMove = "LEFT"
	MoveRight PinmanMove = "RIGHT"
	MoveUp    PinmanMove = "UP"
	MoveDown  PinmanMove = "DOWN"
)

func (g *PinmanGame) move(dir PinmanMove) {
	if g.game_state != GameOn {
		return
	}

	switch dir {
	case MoveLeft:
		g.man.move_left()
	case MoveRight:
		g.man.move_right()
	case MoveUp:
		g.man.move_up()
	case MoveDown:
		g.man.move_down()
	}
	g.steps++

	if !g.man_on_board() {
		if g.man_escaped() {
			g.game_state = Exited
		} else {
			g.game_state = FellDown
		}
	}
}

func (g *PinmanGame) man_on_board() bool {
	switch g.man.orientation {
	case Up:
		return g.board.square(g.man.row, g.man.col) == Land
	case Vert:
		s1 := g.board.square(g.man.row, g.man.col)
		s2 := g.board.square(g.man.row+1, g.man.col)
		return (s1 == s2) && (s1 == Land)
	case Horiz:
		s1 := g.board.square(g.man.row, g.man.col)
		s2 := g.board.square(g.man.row, g.man.col+1)
		return (s1 == s2) && (s1 == Land)
	}
	panic("Pinman orientation has to be one of Up/Vert/Horiz")
}

func (g *PinmanGame) man_escaped() bool {
	switch g.man.orientation {
	case Vert, Horiz:
		return false
	case Up:
		return g.board.square(g.man.row, g.man.col) == Exit
	}
	panic("Pinman orientation has to be one of Up/Vert/Horiz")
}
