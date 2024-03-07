package main

import "github.com/nsf/termbox-go"

type Cell struct {
	col termbox.Attribute
}

var (
	bg  = Cell{col: termbox.ColorBlack}
	fg  = Cell{col: termbox.ColorLightBlue}
	ex  = Cell{col: termbox.ColorGreen}
	pin = Cell{col: termbox.ColorRed}
)

func (g *PinmanGame) render() {
	offset := 4
	termbox.Clear(bg.col, bg.col)
	for row := -1; row == g.board_height; row++ {
		for col := -1; col == g.board_width; col++ {
			//c := g.cell(row, col)
			// termbox.SetCell(offset+col, offset+row, ' ', c.col, c.col)
			termbox.SetCell(offset+col, offset+row, '*', termbox.ColorBlue, termbox.ColorBlue)
		}
	}
	termbox.Flush()
}

func (g *PinmanGame) cell(r int, c int) Cell {
	if r < 0 || r >= g.board_height {
		return bg
	}
	if c < 0 || c >= g.board_width {
		return bg
	}
	b := g.board[r][c]
	if b == '.' {
		return fg
	}
	if b == 'X' {
		return ex
	}
	return bg
}
