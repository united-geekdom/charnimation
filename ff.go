/*
this file contains function that deal with handling the file format
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/stoicperlman/fls"
)

func frameNOf(f *fls.File, n int64, w int) ([]string, error) {
	pos, _ := f.SeekLine(n, io.SeekStart)
	f.Seek(pos, io.SeekStart)
	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	if err != nil {
		return []string{}, errors.New("EOF")
	}
	return tochunks(string(line), w), nil
}

func dataOf(f *fls.File) (int, int, int) { //width height fps
	pos, _ := f.SeekLine(0, io.SeekStart)
	f.Seek(pos, io.SeekStart)
	reader := bufio.NewReader(f)
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	tentative := strings.Split(string(line), " ")
	fmt.Println(tentative)
	if len(tentative) < 3 {
		log.Fatal("Invalid metadata descriptor")
	}
	keys := [3]int{}

	for a := 0; a < 3; a++ {
		keys[a], err = strconv.Atoi(tentative[a])
		if err != nil {
			log.Fatal("Invalid file descriptor")
		}
	}
	fmt.Println(keys)
	return keys[0], keys[1], keys[2]
}
