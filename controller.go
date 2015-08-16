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
	InitView(m)

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
		case <-time.After(5 * time.Second):
			break loop
		}
	}
	close(event)
	termbox.Close()

}
