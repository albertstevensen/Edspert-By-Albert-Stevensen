package main

import (
	"fmt"
)

type groupChars []rune

func iniVokal(c rune) bool {
	vokal := groupChars{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, value := range vokal {
		if value == c {
			return true
		}
	}
	return false
}

func main() {
	myString := "ALBERT STEVENSEN HANYA ADA 1"
	t := 0
	for _, value := range myString {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			t++
		}
	}
	fmt.Printf("Huruf Vokal yang ditemukan ada: %d", t)

}
