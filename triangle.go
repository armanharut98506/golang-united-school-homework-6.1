package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t *Triangle) CalcArea() {
	return math.Sqrt(3.)/4. * (r.Side ** 2.)
}

func (t *Triangle) CalcPerimeter() {
	return 3. * t.Side
}
