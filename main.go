/*
what it does for now:
Theres a class called "film" which serves as a container to manipulate the file
a separate function to seek frame n
and one that returns the correct frame, formatted properly
*/
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/stoicperlman/fls"
)

var file string
var wi uint
var he uint

func init() {
	flag.StringVar(&file, "f", "", ".CHNM file to read")

}

func main() {
	flag.Parse()
	s, _ := tcell.NewScreen()
	s.Init()
	s.Show()

	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					s.Fini()
					os.Exit(1)
					return
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	film := fls.LineFile(f)
	wi, _, fps := dataOf(film)
	fmt.Println(fps)
	fcount := 1
	delay := uint(time.Second) / uint(fps)
	for {
		frame, err := frameNOf(film, int64(fcount), wi)
		if err != nil {
			fcount = 1
			frame, err = frameNOf(film, int64(fcount), wi)
		}
		putframe(s, frame)
		time.Sleep(time.Duration(delay))
		fcount++
	}

}
