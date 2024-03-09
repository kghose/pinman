package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type GridGlyph struct {
	glyph rune
	col   termbox.Attribute
}

var (
	void   = GridGlyph{glyph: ' ', col: termbox.ColorBlack}
	square = GridGlyph{glyph: 0x2b1c, col: termbox.ColorLightBlue}
	exit   = GridGlyph{glyph: 0x2b1b, col: termbox.ColorGreen}
	pinman = GridGlyph{glyph: 0x2b1b, col: termbox.ColorRed}
)

const (
	col_offset = 4
	row_offset = 2
)

func (g *PinmanGame) render() {
	termbox.Clear(void.col, void.col)
	g.board.render()
	g.man.render()
	g.render_status()
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func (g *PinmanGame) render_status() {

	message(g.board.board_height+row_offset+3, col_offset, "Q=quit,R=Restart")
	message(g.board.board_height+row_offset+1, col_offset, fmt.Sprintf("Steps: %d", g.steps))

	if g.game_state == Exited {
		message(0, col_offset, "You Escaped!! Press n for next game.")
	}
	if g.game_state == FellDown {
		message(0, col_offset, "INTO THE ABYSS!! Press r to restart.")
	}

}
func (g *PinmanGameBoard) render() {
	for row := 0; row <= g.board_height; row++ {
		for col := 0; col <= g.board_width; col++ {
			c := g.cell_glyph(row, col)
			termbox.SetCell(col_offset+col*2, row_offset+row, c.glyph, c.col, void.col)
		}
	}
}

func (g *PinmanGameBoard) cell_glyph(r int, c int) GridGlyph {
	switch g.square(r, c) {
	case Abyss:
		return void
	case Land:
		return square
	case Exit:
		return exit
	default:
		return void
	}
}

func (p *Pinman) render() {
	switch p.orientation {
	case Up:
		termbox.SetCell(col_offset+p.col*2, row_offset+p.row, pinman.glyph, pinman.col, pinman.col)
	case Vert:
		termbox.SetCell(col_offset+p.col*2, row_offset+p.row, pinman.glyph, pinman.col, pinman.col)
		termbox.SetCell(col_offset+p.col*2, row_offset+p.row+1, pinman.glyph, pinman.col, pinman.col)
	case Horiz:
		termbox.SetCell(col_offset+p.col*2, row_offset+p.row, pinman.glyph, pinman.col, pinman.col)
		termbox.SetCell(col_offset+p.col*2+2, row_offset+p.row, pinman.glyph, pinman.col, pinman.col)
	}
}

func message(row int, col int, m string) {
	for n, c := range m {
		termbox.SetCell(col+n, row, c, termbox.ColorBlack, termbox.ColorWhite)
	}
}
