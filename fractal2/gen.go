package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := 200
	fmt.Println(c)
	for i := 0; i < c; i++ {
		PrintCase()
	}
}

func PrintCase() {
	n := 100
	fmt.Println(n)

	for i := 0; i < n; i++ {
		x := int(rand.Float64()*2001 - 1000)
		y := int(rand.Float64()*2001 - 1000)
		fmt.Println(x, y)
	}
	depth := 10
	fraction := rand.Float64()
	fmt.Println(depth)
	fmt.Println(fraction)
}
