package main

import (
	"container/list"
	"fmt"
	"os"
)

type Location struct {
	r, c int
}

type Direction struct {
	dr, dc int
}

var NORTH = Direction{-1, 0}
var EAST = Direction{0, 1}
var WEST = Direction{0, -1}
var SOUTH = Direction{1, 0}

var r, c int
var queue = list.New()
var visited [][]bool
var openSouth [][]bool
var openEast [][]bool

func Visit(loc Location) {
	queue.PushBack(loc)
	visited[loc.r][loc.c] = true
}

func Visited(loc Location) bool {
	return visited[loc.r][loc.c]
}

func OpenNorth(loc1, loc2 Location) {
	openSouth[loc2.r][loc2.c] = true
}

func OpenSouth(loc1, loc2 Location) {
	openSouth[loc1.r][loc1.c] = true
}

func OpenEast(loc1, loc2 Location) {
	openEast[loc1.r][loc1.c] = true
}

func OpenWest(loc1, loc2 Location) {
	openEast[loc2.r][loc2.c] = true
}

func Move(loc Location, d Direction) Location {
	return Location{loc.r + d.dr, loc.c + d.dc}
}

func Explore(loc Location) {
	if new := Move(loc, NORTH); loc.r > 0 && !Visited(new) {
		OpenNorth(loc, new)
		Visit(new)
	}
	if new := Move(loc, EAST); loc.c < c-1 && !Visited(new) {
		OpenEast(loc, new)
		Visit(new)
	}
	if new := Move(loc, SOUTH); loc.r < r-1 && !Visited(new) {
		OpenSouth(loc, new)
		Visit(new)
	}
	if new := Move(loc, WEST); loc.c > 0 && !Visited(new) {
		OpenWest(loc, new)
		Visit(new)
	}

}

func PrintLayout() {
	for i := 0; i < c; i++ {
		os.Stdout.Write([]byte(" _"))
	}
	os.Stdout.Write([]byte("\n"))

	for y := 0; y < r; y++ {
		buf := make([]byte, c*2+2)
		buf[0] = '|'
		for x := 0; x < c; x++ {
			if openSouth[y][x] {
				buf[1+x*2] = ' '
			} else {
				buf[1+x*2] = '_'
			}
			if openEast[y][x] {
				buf[2+x*2] = ' '
			} else {
				buf[2+x*2] = '|'
			}
		}
		buf[2*c+1] = '\n'
		os.Stdout.Write(buf)
	}
}
func main() {
	fmt.Scanln(&r, &c)

	openSouth = make([][]bool, r)
	openEast = make([][]bool, r)
	visited = make([][]bool, r)
	for i := 0; i < r; i++ {
		openSouth[i] = make([]bool, c)
		openEast[i] = make([]bool, c)
		visited[i] = make([]bool, c)
	}

	if r%2 == 0 && c%2 == 0 {
		// four center rooms, just pick one
		Visit(Location{r / 2, c / 2})
	} else if r%2 == 0 {
		// Two vertically centered rooms
		Visit(Location{r / 2, c / 2})
		Visit(Location{r/2 - 1, c / 2})
		OpenNorth(Location{r / 2, c / 2}, Location{r/2 - 1, c / 2})
	} else if c%2 == 0 {
		// Two horizontally centered rooms
		Visit(Location{r / 2, c / 2})
		Visit(Location{r / 2, c/2 - 1})
		OpenWest(Location{r / 2, c / 2}, Location{r / 2, c/2 - 1})
	} else {
		// A single center room
		Visit(Location{r / 2, c / 2})
	}

	for queue.Len() > 0 {
		loc := queue.Front()
		Explore(loc.Value.(Location))
		queue.Remove(loc)
	}

	PrintLayout()
}
