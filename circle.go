package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c *Circle) CalcArea() {
	return math.Pi * c.Radius ** 2.
}

func (c *Circle) CalcPerimeter() {
	return 2. * math.Pi * c.Radius
}
