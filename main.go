/*
what it does for now:
Theres a class called "film" which serves as a container to manipulate the file
a separate function to seek frame n
and one that returns the correct frame, formatted properly
*/
package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/stoicperlman/fls"
)

func main() {

	f, _ := os.OpenFile("example.txt", os.O_RDONLY, 0600)
	defer f.Close()

	a := film{}
	a.file = fls.LineFile(f)
	a.setdimensions(3, 3)
	a.setFPS(60)
	print(a.cframe())
	a.cpos(2)
	print(a.cframe())

	refresh()
	for {
		print(a.cframe())
		time.Sleep(time.Duration(a.delay))
		a.cpos(1)
		refresh()

	}

}

type film struct {
	file       *fls.File
	framecount int64
	delay      uint
	dimension  struct {
		width  int
		height int
	}
}

func (f *film) cpos(n int64) {
	f.framecount += n
}

func (f *film) setdimensions(w int, h int) {
	f.dimension.width = w
	f.dimension.height = h
}

func (f *film) setFPS(fps int) {
	f.delay = uint(time.Second) / uint(fps)
}

//https://stackoverflow.com/a/61469854
func (f *film) cframe() string {

	s := f.seekframe(f.framecount)
	if len(s) == 0 {
		return "nil"
	}
	if f.dimension.width >= len(s) {
		return ""
	}
	if len(f.seekframe(f.framecount))%f.dimension.height != 0 {
		os.Exit(1)
	}
	var chunks []string = make([]string, 0, (len(s)-1)/f.dimension.width+1)
	currentLen := 0
	currentStart := 0
	for i := range s {
		if currentLen == f.dimension.width {
			chunks = append(chunks, s[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, s[currentStart:])
	c := strings.Join(chunks, "\n") + "\n"
	return c
}

//line numbers start from 0
func (f *film) seekframe(n int64) string {

	pos, _ := f.file.SeekLine(n, io.SeekStart)
	f.file.Seek(pos, io.SeekStart)
	reader := bufio.NewReader(f.file)
	line, _, err := reader.ReadLine()
	if err != nil {
		f.framecount = 0
		return f.seekframe(f.framecount)
	}
	return string(line)
}

func refresh() {
	print("\033[H\033[J")
}
