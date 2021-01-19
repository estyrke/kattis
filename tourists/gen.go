package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	nPtr := flag.Int("n", 200000, "number of nodes")
	starPtr := flag.Bool("star", false, "use a star shape")
	linePtr := flag.Bool("line", false, "use a line shape")
	flag.Parse()

	n := *nPtr
	fmt.Println(n)

	for i := 0; i < n-1; i++ {
		var v1, v2 int
		switch {
		case *starPtr:
			v1, v2 = i+1, 1
		case *linePtr:
			v1, v2 = i+1, i+2
		default:
			fmt.Fprintln(os.Stderr, "No shape specified!")
			return
		}
		fmt.Println(v1, v2)
	}
}
