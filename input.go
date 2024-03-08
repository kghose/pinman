package main

import "github.com/nsf/termbox-go"

func (g *PinmanGame) process_key(ev termbox.Event) {
	switch ev.Key {
	case termbox.KeyArrowLeft:
		g.man.move_left()
	case termbox.KeyArrowRight:
		g.man.move_right()
	case termbox.KeyArrowUp:
		g.man.move_up()
	case termbox.KeyArrowDown:
		g.man.move_down()
	}
}
