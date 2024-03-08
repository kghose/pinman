package main

import (
	"github.com/nsf/termbox-go"
)

func main() {

	game := PinmanGame{}
	game.load(1)

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	game.render()

	event_q := make(chan termbox.Event)
	go func() {
		for {
			event_q <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-event_q:
			if ev.Type == termbox.EventKey {
				if quit_key(ev) {
					return
				}
				game.process_key(ev)
				game.render()
			}
		}
	}
}

func quit_key(ev termbox.Event) bool {
	return ev.Ch == 'q' ||
		ev.Key == termbox.KeyEsc ||
		ev.Key == termbox.KeyCtrlC ||
		ev.Key == termbox.KeyCtrlD
}
