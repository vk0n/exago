package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Word struct {
	Value string
	Count int
	Index int
}

type WordList []Word

func (w WordList) Len() int { return len(w) }
func (w WordList) Less(i, j int) bool {
	return w[i].Count < w[j].Count || (w[i].Count == w[j].Count && w[i].Index > w[j].Index)
}
func (w WordList) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {
	var text string
	fmt.Println("Enter your text:")
	reader := bufio.NewReader(os.Stdin)

	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	words := regexp.MustCompile("[,.#$*&@+?!=(/);: -]").Split(text, -1)
	words = delete_empty(words)

	wl := make(WordList, len(words))
	i := 0
	for idx, word := range words {
		new := true
		for j := 0; j < wl.Len(); j++ {
			if word == wl[j].Value {
				wl[j] = Word{wl[j].Value, wl[j].Count + 1, wl[j].Index}
				new = false
			}
		}

		if new {
			wl[i] = Word{word, 1, idx}
			i++
		}
	}

	sort.Sort(sort.Reverse(wl))

	fmt.Println("Result:")
	for _, w := range wl {
		if w.Count > 0 {
			fmt.Print(w.Value + "(" + strconv.Itoa(w.Count) + ") ")
		}
	}
	fmt.Println()
}
