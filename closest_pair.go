package closest_pair

import "sort"
import "math"
import "strconv"

func ClosestPair(points []Point) ([2]Point, error) {

	if points == nil || len(points) < 2 {
		return [2]Point{}, &Error{code: 1, desc: "At least two Points are required to find the closest pair"}
	}
	sort.Sort(xSortable(points))
	return closest_pair(points), nil
}

func closest_pair(points []Point) [2]Point {
	var size = len(points)
	if size < 3 {
		return bruteforce_closest_pair(points)
	}

	var left = closest_pair(points[0:size/2])
	var right = closest_pair(points[size/2:size])

	var closest = closest(left, right)
	var min_dist = closest[0].Distance(closest[1])
	var center_x = (points[size/2].x + points[size/2-1].x) / 2

	var candidates []Point = get_close_enough_points(center_x, min_dist, points)

	var isEnoughCandidatesToOverrideClosestPair = len(candidates) > 1
	if !isEnoughCandidatesToOverrideClosestPair {
		return closest
	}

	sort.Sort(ySortable(candidates))
	var closest_candidates = bruteforce_closest_pair(candidates)

	if closest_candidates[0].Distance(closest_candidates[1]) < min_dist {
		closest = closest_candidates
		min_dist = closest_candidates[0].Distance(closest_candidates[1])
	}
	return closest
}

func get_close_enough_points(center_x int, dist float64, points []Point) []Point {

	var candidates []Point

	for _, point := range points {
		if math.Abs(float64(point.x-center_x)) < dist {
			candidates = append(candidates, point)
		}
	}
	return candidates
}

func bruteforce_closest_pair(points []Point) [2]Point {
	var closest [2]Point
	var min_dist = math.Inf(1)
	var i int
	var j int
	var size = len(points)

	for i = 0; i < size; i++ {
		for j = i + 1; j < size; j++ {
			var dist = points[i].Distance(points[j])
			if dist < min_dist {
				min_dist = dist
				closest = [2]Point{points[i], points[j]}
			}
		}
	}
	return closest
}

func closest(left [2]Point, right [2]Point) [2]Point {
	var left_dist = left[0].Distance(left[1])
	var right_dist = right[0].Distance(right[1])
	if left_dist < right_dist {
		return left
	}
	return right
}

type xSortable []Point

func (p xSortable) Len() int {
	return len(p)
}
func (p xSortable) Less(i, j int) bool {
	return p[i].x < p[j].x
}
func (p xSortable) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ySortable []Point

func (p ySortable) Len() int {
	return len(p)
}
func (p ySortable) Less(i, j int) bool {
	return p[i].y < p[j].y
}
func (p ySortable) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Error struct {
	code int
	desc string
}

func (e *Error) Error() string {
	return strconv.Itoa(e.code) + ": " + e.desc
}
