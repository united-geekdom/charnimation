//skimmer - proof of concept implementation of the charnimation file format
//incredibly hacky, but gets the job done
package main

import (
	"os"
	"strings"
	"time"
)

func main() {
	frames := []string{"o##\n###\n###", "o##\no##\n###", "o##\no##\no##", "o##\no##\noo#", "o##\no##\nooo", "o##\no#o\nooo", "o#o\no#o\nooo", "ooo\no#o\nooo", "#oo\no#o\nooo", "#oo\n##o\nooo", "#oo\n##o\n#oo", "#oo\n##o\n##o", "#oo\n##o\n###", "#oo\n###\n###", "#o#\n###\n###", "###\n###\n###"}
	//frames := getframes("example.txt")
	refresh()
	i := 0
	for {
		print(frames[i])
		//print(i)
		time.Sleep(82 * time.Millisecond)
		i++
		if i == len(frames) {
			i = 0
		}
		refresh()

	}

}
func refresh() {
	print("\033[H\033[J")
}

func getframes(name string) []string {
	a, _ := os.ReadFile(name)
	b := strings.Split(string(a), "\n")
	//fmt.Println(b)
	return b
}
