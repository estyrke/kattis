package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

type Point struct {
	x, y float64
}

type Line struct {
	p1, p2             Point
	dist1, dist2       float64
	length             float64
	scale              float64
	angle              float64
	cosAngle, sinAngle float64
}

func Dist(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}

func PointAdd(p1, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y}
}

func PointDiff(p1, p2 Point) Point {
	return Point{p1.x - p2.x, p1.y - p2.y}
}

func PointRotate(p Point, cosAngle, sinAngle float64) Point {
	return Point{p.x*cosAngle - p.y*sinAngle, p.x*sinAngle + p.y*cosAngle}
}

func PointScale(p Point, scale float64) Point {
	return Point{p.x * scale, p.y * scale}
}

func TargetIndexForFraction(lines []Line, fraction float64) (index int, remainder float64) {
	for i, l := range lines {
		if fraction <= l.dist2 {
			remainder := (fraction - l.dist1) / (l.dist2 - l.dist1)
			return i, remainder
		}
	}
	return -1, -1
}

func MidPoint(l Line, fraction float64) Point {
	return Point{l.p1.x + (l.p2.x-l.p1.x)*fraction, l.p1.y + (l.p2.y-l.p1.y)*fraction}
}

func TransformPointToLine(point Point, line Line) Point {
	return PointAdd(line.p1, PointScale(PointRotate(point, line.cosAngle, line.sinAngle), line.scale))
}

func Solve(lines []Line, fraction float64, depth int) Point {
	i, remainder := TargetIndexForFraction(lines, fraction)
	if depth == 1 {
		return MidPoint(lines[i], remainder)
	} else {
		p := Solve(lines, remainder, depth-1)
		return TransformPointToLine(p, lines[i])
	}
}

func RunCase(r io.Reader) {
	var n int
	fmt.Fscanln(r, &n)

	points := make([]Point, n)
	lines := make([]Line, n-1)

	dist := 0.0

	var offset Point
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscanln(r, &x, &y)
		points[i] = Point{float64(x), float64(y)}
		if i == 0 {
			offset = points[i]
		} else {
			lines[i-1].p1 = PointDiff(points[i-1], offset)
			lines[i-1].p2 = PointDiff(points[i], offset)
			length := Dist(points[i-1], points[i])
			lines[i-1].length = length
			lines[i-1].dist1 = dist
			dist += length
			lines[i-1].dist2 = dist
			lines[i-1].angle = math.Atan2(points[i].y-points[i-1].y, points[i].x-points[i-1].x)
		}
	}
	totalAngle := math.Atan2(points[n-1].y-points[0].y, points[n-1].x-points[0].x)
	lengthScale := 1 / lines[n-2].dist2
	distScale := 1 / Dist(points[0], points[n-1])
	for i := 0; i < n-1; i++ {
		lines[i].dist1 *= lengthScale
		lines[i].dist2 *= lengthScale
		lines[i].scale = lines[i].length * distScale
		angle := lines[i].angle - totalAngle
		lines[i].cosAngle = math.Cos(angle)
		lines[i].sinAngle = math.Sin(angle)
	}

	var depth int
	var fraction float64

	fmt.Fscanln(r, &depth)
	fmt.Fscanln(r, &fraction)

	targetPoint := PointAdd(Solve(lines, fraction, depth), offset)
	fmt.Println(targetPoint.x, targetPoint.y)
}

func main() {
	// f, err := os.Create("fractal2.prof")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()
	var cases int
	reader := io.Reader(os.Stdin)

	fmt.Fscanln(reader, &cases)

	for c := 0; c < cases; c++ {
		RunCase(reader)
	}
}
