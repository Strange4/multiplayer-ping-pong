package main

import (
	"github.com/gdamore/tcell/v2"
)

func main() {
	s, err := tcell.NewScreen()
	s.Colors()
	if err != nil {
		panic(err)
	}
	if err = s.Init(); err != nil {
		panic(err)
	}
	// s.Clear()
	defStyle := tcell.StyleDefault
	s.SetContent(0, 0, 'H', nil, defStyle)
	s.SetContent(1, 0, 'i', nil, defStyle)
	s.SetContent(2, 0, '!', nil, defStyle)
	s.Show()
}
