package main

import (
	"fmt"
	"strings"
)

var chars = "abcdefghijklmnopqrstuvwxyz"

var b strings.Builder

func Word(num int) string {
	b.Reset()

	for {
		fmt.Fprintf(&b, "%c", chars[num%16])
		num >>= 4
		if num == 0 {
			break
		}
	}
	return b.String()
}

func main() {
	var a, b int

	fmt.Scanln(&a, &b)

	for word := 0; word < a || word < (b+1)/2; word++ {
		fmt.Print(Word(word), " ")
	}
	fmt.Println()
}
