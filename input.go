package main

import "github.com/nsf/termbox-go"

func (g *PinmanGame) Start() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	event_q := make(chan termbox.Event)
	go g.event_poll_loop(event_q)
	g.main_loop(event_q)
}

func (g *PinmanGame) event_poll_loop(event_q chan termbox.Event) {
	for {
		event_q <- termbox.PollEvent()
	}
}

func (g *PinmanGame) main_loop(event_q chan termbox.Event) {
	g.render()
	game_on := true
	for game_on {
		select {
		case ev := <-event_q:
			if ev.Type == termbox.EventKey {
				key_type := g.process_key(ev)

				switch key_type {
				case Quit:
					game_on = false
				case Restart:
					g.Load(g.board.board_no)
				case NextGame:
					g.Load(g.board.board_no + 1)
				}

				g.render()
			}
		}
	}
}

type KeyType string

const (
	Move     KeyType = "MOVE"
	Quit     KeyType = "QUIT"
	Restart  KeyType = "RESTART"
	NextGame KeyType = "NEXTGAME"
)

func (g *PinmanGame) process_key(ev termbox.Event) KeyType {
	if quit_key(ev) {
		return Quit
	}
	if ev.Ch == 'r' {
		return Restart
	}
	if ev.Ch == 'n' && g.game_state == Exited {
		return NextGame
	}
	switch ev.Key {
	case termbox.KeyArrowLeft:
		g.move(MoveLeft)
	case termbox.KeyArrowRight:
		g.move(MoveRight)
	case termbox.KeyArrowUp:
		g.move(MoveUp)
	case termbox.KeyArrowDown:
		g.move(MoveDown)
	}
	return Move
}
func quit_key(ev termbox.Event) bool {
	return ev.Ch == 'q' ||
		ev.Key == termbox.KeyEsc ||
		ev.Key == termbox.KeyCtrlC ||
		ev.Key == termbox.KeyCtrlD
}
