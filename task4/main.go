package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cipher(text string, shft int, direction int) string {
	shift, offset := rune(shft), rune(26)
	runes := []rune(text)

	for index, char := range runes {
		switch direction {
		case -1: // encoding
			if char >= 'a' && char+shift <= 'z' ||
				char >= 'A' && char+shift <= 'Z' {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		case +1: // decoding
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		}

		runes[index] = char
	}

	return string(runes)
}

func encode(text string, shift int) string { return cipher(text, shift, -1) }
func decode(text string, shift int) string { return cipher(text, shift, +1) }

func main() {
	var text string
	var words string
	var words_list []string

	fmt.Println("Enter your text:")
	reader := bufio.NewReader(os.Stdin)

	text, _ = reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	fmt.Println("Enter 5 source words (comma-separated):")
	words, _ = reader.ReadString('\n')
	words = strings.Replace(words, "\n", "", -1)
	words_list = strings.Split(words, ",")

	shift_found := false
	shift := 0
	for shift < 26 && !shift_found {
		shift++
		shift_found = true
		for _, w := range words_list {
			w = strings.TrimSpace(w)
			if !strings.Contains(text, encode(w, shift)) {
				shift_found = false
			}
		}
	}

	if shift_found {
		fmt.Println("Shift:", shift)
		fmt.Println(decode(text, shift))
	} else {
		fmt.Println("Decoding failed!")
	}
}
