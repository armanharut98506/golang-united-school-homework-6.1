package golang_united_school_homework

import (
	"errors"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) + 1 > cap(b.shapes) {
		return errors.New("Exceeds capacity of the box")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		errors.New("Index out of range")
	}

	for index, _ := range b.shapes {
		if index == i {
			return b.shapes[index], nil
		}
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("Index out of range")
	}

	for index, shape := range b.shapes {
		if index == i {
			result := shape
			b.shapes = append(b.shapes[:index], b.shapes[index+1:]...)
			return result, nil
		}
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("Index out of range")
	}

	for index, shape := range b.shapes {
		if index == i {
			result := shape
			b.shapes[index] = shape
			return result, nil
		}
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	result := 0.0
	for index, shape := range b.shapes {
		result += shape.CalcPerimeter()
	}
	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	result := 0.0
	for index, shape := range b.shapes {
		result += shape.CalcArea()
	}
	return result
}

// RemoveAllCircles removes all circles in the list
/* whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	numCircles := 0
	for index, shape := range b.shapes {
		circle := shape.(Circle)
		shapes = append(b.shapes[:index], b.shapes[index+1:]...)
		umCircles++
	}
	if numCircles == 0 {
		return errors.New("No circles in the box")
	}
}
*/