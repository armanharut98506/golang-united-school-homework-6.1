package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func (r *Rectangle) CalcArea() {
	return r.Height * R.Weight
}

func (r *Rectangle) CalcPerimeter() {
	return 2. * (r.Height + r.Weight)
}

