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

	banner_width := 2*g.board.board_width + 2*col_offset
	banner(
		g.board.board_height+row_offset+3,
		0,
		banner_width,
		"q=quit,r=restart",
		true,
		termbox.ColorBlack,
		termbox.ColorGreen,
	)
	banner(
		g.board.board_height+row_offset+2,
		0,
		banner_width,
		fmt.Sprintf("Steps: %d", g.steps),
		true,
		termbox.ColorLightYellow,
		termbox.ColorLightGray,
	)

	if g.game_state == Exited {
		banner(
			0,
			0,
			banner_width,
			"You Escaped!! n=next game.",
			true,
			termbox.ColorYellow,
			termbox.ColorWhite,
		)
	}
	if g.game_state == FellDown {
		banner(
			0,
			0,
			banner_width,
			"You FELL!!",
			true,
			termbox.ColorWhite,
			termbox.ColorRed,
		)
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
		termbox.SetCell(
			col_offset+p.col*2,
			row_offset+p.row+1,
			pinman.glyph,
			pinman.col,
			pinman.col,
		)
	case Horiz:
		termbox.SetCell(col_offset+p.col*2, row_offset+p.row, pinman.glyph, pinman.col, pinman.col)
		termbox.SetCell(
			col_offset+p.col*2+2,
			row_offset+p.row,
			pinman.glyph,
			pinman.col,
			pinman.col,
		)
	}
}

func banner(
	row int,
	col int,
	width int,
	msg string,
	center bool,
	fg termbox.Attribute,
	bg termbox.Attribute,
) {
	if center {
		col = max(0, width/2-len(msg)/2)
	}
	for n := 0; n < col; n++ {
		termbox.SetCell(n, row, ' ', bg, bg)
	}
	for n, c := range msg {
		termbox.SetCell(col+n, row, c, fg, bg)
	}
	for n := col + len(msg); n < width; n++ {
		termbox.SetCell(n, row, ' ', bg, bg)
	}

}
