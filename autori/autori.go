package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		c, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		if c >= 'A' && c <= 'Z' {
			fmt.Print(string(c))
		} else if c == '\n' {
			fmt.Println()
			break
		}
	}
}
