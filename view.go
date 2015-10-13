package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"strings"
)

type View struct {
	height int
	width  int
	ptr    int
	quit   bool
}

func InitView(m *Model) *View {
	w, h := termbox.Size()
	v := &View{
		height: h,
		width:  w,
		ptr:    0,
		quit:   false,
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	v.printHeader()
	for y, c := range m.contents {
		if y == 0 {
			v.println(y+1, c, true)
		} else {
			v.println(y+1, c, false)
		}
	}
	return v
}

func (v *View) Exit() {
	v.quit = true
}

func (v *View) LineUp(m *Model) {
	if v.ptr == 1 {
		return
	}
	v.println(v.ptr, m.contents[v.ptr-1], false)
	v.ptr--
	v.println(v.ptr, m.contents[v.ptr-1], true)
}

func (v *View) LineDown(m *Model) {
	if v.ptr == len(m.contents) {
		return
	}
	v.println(v.ptr, m.contents[v.ptr-1], false)
	v.ptr++
	v.println(v.ptr, m.contents[v.ptr-1], true)
}

func (v *View) printHeader() {
	hedderLeft := "Command"
	hedderRight := "Memo"
	v.printTwoColumn(0, hedderLeft, hedderRight, false)
}

func (v *View) printTwoColumn(y int, left, right string, highlight bool) {
	leftLimit := v.width / 2
	rightLimit := (v.width - 1) / 2
	leftTruncated := runewidth.Truncate(left, leftLimit, "")
	rightTruncated := runewidth.Truncate(right, rightLimit, "")
	leftFilled := leftTruncated + strings.Repeat(" ", leftLimit-len(leftTruncated))
	rowMsg := leftFilled + "|" + rightTruncated
	v.println(y, rowMsg, highlight)
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
		for x := len(msg); x < v.width; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
		}
	}
	termbox.Flush()
}
