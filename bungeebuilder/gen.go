package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	nPtr := flag.Int("n", 1000000, "number of nodes")
	flag.Parse()

	n := *nPtr
	fmt.Println(n)

	for i := 0; i < n-1; i++ {
		var h = rand.Int()
		fmt.Print(h, " ")
	}
}
