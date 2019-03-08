package main

import "fmt"

func main() {
	point := &Point{1.5, 5}
	elongate(point, 3)
	fmt.Println("New point", point.X, point.Y)
}

type Point struct {
	X float32
	Y float32
}

func elongate(point *Point, factor float32) {
	point.X = point.X * factor
	point.Y = point.Y * factor
}
