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

	print(a.frame())
	a.cpos(2)
	print(a.frame())

	refresh()
	for {
		print(a.frame())
		time.Sleep(41 * time.Millisecond)
		a.cpos(1)
		refresh()

	}

}

func (f *film) cpos(n int64) {
	f.framecount += n
}

func (f *film) setdimensions(w int, h int) {
	f.dimension.width = w
	f.dimension.height = h
}

type film struct {
	file       *fls.File
	framecount int64
	buf        []string
	dimension  struct {
		width  int
		height int
	}
}

/*func (f *film) frame() {
	if len(f.seekframe(f.framecount))%f.dimension.height != 0 {
		os.Exit(1)
	}
	fmt.Println(split(f.seekframe(f.framecount), f.dimension.width))

}*/

//https://stackoverflow.com/a/61469854
func (f *film) frame() string {

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
