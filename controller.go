package main

import (
	"github.com/nsf/termbox-go"
	"log"
	"time"
)

func main() {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	m := InitModel()
	v := InitView(m)

	event := make(chan termbox.Event)
	go func() {
		for {
			// Post events to channel
			event <- termbox.PollEvent()
		}
	}()

loop:
	for {
		select {
		case ev := <-event:
			updateEvent(ev, v, m)
		}
		if v.quit {
			break loop
		}
	}
	close(event)
	termbox.Close()
}

func updateEvent(ev termbox.Event, v *View, m *Model) {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Ch {
		case 'q':
			v.Exit()
		case 'k':
			v.LineUp(m)
		case 'j':
			v.LineDown(m)
		}
	}
}
