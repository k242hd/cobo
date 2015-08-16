package main

import "github.com/nsf/termbox-go"

type View struct {
	height int
	width  int
	ptr    int
}

func InitView(m *Model) *View {
	w, h := termbox.Size()
	v := &View{
		height: h,
		width:  w,
		ptr:    0,
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for y, c := range m.contents {
		if y == 0 {
			v.println(y, c, true)
		} else {
			v.println(y, c, false)
		}
	}
	return v
}

func (v *View) println(y int, msg string, highlight bool) {
	if highlight {
		for x, c := range msg {
			termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorMagenta)
		}
		for x := len(msg); x < v.width; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorMagenta)
		}
	} else {
		for x, c := range msg {
			termbox.SetCell(x, y, c, termbox.ColorWhite, termbox.ColorDefault)
		}
	}
	termbox.Flush()
}
