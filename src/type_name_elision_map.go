package main

type Point2D struct {
	x, y float64
}

// START OMIT
var square = map[Point2D]struct{}{
	{0, 0}: {}, // instead of Point2D{0, 0}: struct{}{}
	{1, 0}: {},
	{1, 1}: {},
	{0, 1}: {},
}

// END OMIT

func main() {}
