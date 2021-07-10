package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func getDistance(p1, p2 Point) float64 {
	d1 := math.Pow(p1.x - p2.x, 2)
	d2 := math.Pow(p1.y - p2.y, 2)
	return math.Sqrt(d1 + d2)
}

func getShortTestLine(a []Point) (Point, Point) {
	len := len(a)
	min := 0.0
	var cord1, cord2 Point

	for i := 0; i < len; i++ {
		for j := i+1; j < len; j++ {
			d := getDistance(a[i], a[j])
			if d < min || min == 0 {
				min = d
				cord1 = a[i]
				cord2 = a[j]
			}
		}
	}
	return cord1, cord2
}

func main() {
	array := []Point{{1,1},{-1,-1},{3,4},{6,1},{-1,-6},{-4,-3}}
	fmt.Println(getShortTestLine(array))
}