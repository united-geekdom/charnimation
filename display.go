package main

import (
	"github.com/gdamore/tcell"
	"github.com/mattn/go-runewidth"
)

func putframe(s tcell.Screen, frame []string) {
	x, y := 0, 0
	s.Clear()
	for _, i := range frame {
		for _, c := range i {
			var comb []rune
			w := runewidth.RuneWidth(c)
			if w == 0 {
				comb = []rune{c}
				c = ' '
				w = 1
			}
			s.SetContent(x, y, c, comb, tcell.StyleDefault)
			x += w
		}
		x = 0
		y++
	}
	s.Show()
}

func tochunks(s string, w int) []string {
	var chunks []string = make([]string, 0, (len(s)-1)/w+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == w {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	return chunks
}
