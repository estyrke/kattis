package main

import (
	"fmt"
	"math"
)

func ExpandedLength(L, n, C float64) float64 {
	return (1 + n*C) * L
}

func Displacement(L, n, C float64) float64 {
	Le := ExpandedLength(L, n, C)
	angle := 0.0
	radius := 0.0

	//math.Sin(angle / 2) = (L / (2 * radius))
	//
	angle = math.Asin(L/(2*radius)) * 2
	angle = radius
	//Le = angle * radius
	//(1 + n*C) * L = angle * radius
	radius = Le / angle

	radius = (L*L + 4*h*h) / (8 * h)
	8 * h * radius = L*L + 4*h*h
	8*h*radius - 4*h*h = L * L
	4 * h * (2*radius - h) = L * L

	//Distance of the chord from the center is the b side of the right triangle with hypotenuse equal to the radius
	// and a side equal to half the original length of the rod
	chordDistance := math.Sqrt(radius*radius + (L/2)*(L/2))

	height := radius - math.Sqrt(4*radius*radius-L*L)/2
	height = radius - chordDistance
	return height
}

func main() {
	var L, n int
	var C float64
	for fmt.Scanln(&L, &n, &C); L != -1; fmt.Scanln(&L, &n, &C) {
		fmt.Println(Displacement(float64(L), float64(n), C))
	}
}
