package main

import "github.com/nsf/termbox-go"

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
	err := termbox.Flush()
	if err != nil {
		panic(err)
	}
}

func (g *PinmanGameBoard) render() {
	for row := 0; row <= g.board_height; row++ {
		for col := 0; col <= g.board_width; col++ {
			c := g.cell(row, col)
			termbox.SetCell(col_offset+col*2, row_offset+row, c.glyph, c.col, void.col)
		}
	}
}

func (g *PinmanGameBoard) cell(r int, c int) GridGlyph {
	if r < 0 || r >= g.board_height {
		return void
	}
	if c < 0 || c >= len(g.board[r]) {
		return void
	}
	b := g.board[r][c]
	if b == '.' || b == 'P' {
		return square
	}
	if b == 'X' {
		return exit
	}
	return void
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
