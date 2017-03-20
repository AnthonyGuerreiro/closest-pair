package closest_pair

import "math"

type Point struct {
	x int
	y int
}

type Interface interface {
	Distance(p Point) float64
}

func (p1 Point) Distance(p2 Point) float64 {
	var squaredXDiff = math.Pow(float64(p2.x-p1.x), 2)
	var squaredYDiff = math.Pow(float64(p2.y-p1.y), 2)
	return math.Sqrt(squaredXDiff + squaredYDiff)
}
