package main

import (
	"fmt"
	"strconv"
)

func compress(text string) string {
	fmt.Println("Compressing...")

	var sl_text []rune = []rune(text)
	var current rune = sl_text[0]
	var count int = 1
	var encoding []rune

	for _, char := range sl_text[1:] {
		if char == current {
			count++
		} else {
			if count > 4 {
				rune_count := []rune(strconv.FormatInt(int64(count), 10))
				encoding = append(encoding, '#')
				encoding = append(encoding, rune_count...)
				encoding = append(encoding, '#', current)
			} else {
				for i := 0; i < count; i++ {
					encoding = append(encoding, current)
				}
			}
			count = 1
			current = char
		}
	}

	if count > 4 {
		rune_count := []rune(strconv.FormatInt(int64(count), 10))
		encoding = append(encoding, '#')
		encoding = append(encoding, rune_count...)
		encoding = append(encoding, '#', current)
	} else {
		for i := 0; i < count; i++ {
			encoding = append(encoding, current)
		}
	}

	return string(encoding)
}

func decompress(text string) string {
	fmt.Println("Decompressing...")

	var sl_text []rune = []rune(text)
	var current rune
	var sl_count []rune
	var decoding []rune
	var i int64

	for idx := 0; idx < len(sl_text); {
		current = sl_text[idx]
		count, err := strconv.ParseInt(string(sl_count), 10, 64)
		if err != nil {
			count = 1
		}
		if current != '#' {
			for i = 0; i < count; i++ {
				decoding = append(decoding, current)
			}
			sl_count = sl_count[:0]
			idx++
		} else {
			c := idx + 1
			for sl_text[c] != '#' {
				sl_count = append(sl_count, sl_text[c])
				c++
			}
			idx = c + 1
		}
	}

	return string(decoding)
}

func main() {
	fmt.Println("Enter your text:")
	var text string
	fmt.Scanln(&text)
	fmt.Print("Enter 1 to compress, 2 to decompress: ")
	var action int
	fmt.Scanln(&action)

	switch action {
	case 1:
		fmt.Println(compress(text))
	case 2:
		fmt.Println(decompress(text))
	default:
		fmt.Println("Wrong action code!")
	}
}
