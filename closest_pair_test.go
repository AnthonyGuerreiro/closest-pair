package closest_pair

import "testing"
import "fmt"

func TestNil(t *testing.T) {
	var _, err = ClosestPair(nil)
	if err == nil {
		fmt.Println("Expected error when executing ClosestPair with nil args")
		t.FailNow()
	}
}

func TestEmpty(t *testing.T) {
	var _, err = ClosestPair([]Point{})
	if err == nil {
		fmt.Println("Expected error when executing ClosestPair with empty args")
		t.FailNow()
	}
}

func TestOnePoint(t *testing.T) {
	var _, err = ClosestPair([]Point{Point{1, 1}})
	if err == nil {
		fmt.Println("Expected error when executing ClosestPair with one point")
		t.FailNow()
	}
}

func TestTwoPoints(t *testing.T) {
	var p1 = Point{1, 1}
	var p2 = Point{2, 2}
	var closest, err = ClosestPair([]Point{p1, p2})
	if err != nil {
		fmt.Println("Unexpected error", err.Error())
		t.FailNow()
	}

	if closest[0] != p1 {
		fmt.Println("First point in result should be p1")
		t.FailNow()
	}
	if closest[1] != p2 {
		fmt.Println("Second point in result should be p2")
		t.FailNow()
	}
}

func TestMultiplePoints(t *testing.T) {
	var p1 = Point{1, 1}
	var p2 = Point{2, 1}
	var points []Point
	points = append(points, Point{3, 4})
	points = append(points, Point{-3, 3})
	points = append(points, p2)
	points = append(points, Point{-3, -3})
	points = append(points, Point{0, 0})
	points = append(points, Point{1, 3})
	points = append(points, p1)
	points = append(points, Point{3, 2})

	var closest, err = ClosestPair(points)
	if err != nil {
		fmt.Println("Unexpected error ", err.Error())
		t.FailNow()
	}

	if closest[0] != p1 && closest[1] != p1 {
		fmt.Println("closest should contain", p1, ", but is", closest)
		t.Fail()
	}
	if closest[0] != p2 && closest[1] != p2 {
		fmt.Println("closest should contain", p2, ", but is", closest)
		t.Fail()
	}
}

func TestMultiplePoinstHavingSolutionInDifferentHalves(t *testing.T) {

	var p1 = Point{1, 1}
	var p2 = Point{2, 1}
	var points []Point
	points = append(points, Point{0, 0})
	points = append(points, Point{2, 3})
	points = append(points, p1)
	points = append(points, p2)

	var closest, err = ClosestPair(points)
	if err != nil {
		fmt.Println("Unexpected error ", err.Error())
		t.FailNow()
	}

	if closest[0] != p1 && closest[1] != p1 {
		fmt.Println("closest should contain", p1, ", but is", closest)
		t.Fail()
	}
	if closest[0] != p2 && closest[1] != p2 {
		fmt.Println("closest should contain", p2, ", but is", closest)
		t.Fail()
	}

}
