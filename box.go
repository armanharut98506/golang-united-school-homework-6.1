package golang_united_school_homework

import (
	"errors"
	"fmt"
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
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("Exceeds capacity of the box")
	} else {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("Index out of range")
	}

	for index, _ := range b.shapes {
		if index == i {
			return b.shapes[index], nil
		}
	}
	return nil, nil
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
	return nil, nil
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
	return nil, nil 
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	result := 0.0
	for _, shape := range b.shapes {
		result += shape.CalcPerimeter()
	}
	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	result := 0.0
	for _, shape := range b.shapes {
		result += shape.CalcArea()
	}
	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	numCircles := 0
	for index, shape := range b.shapes {
		switch shape.(type) {
		case Circle:
			fmt.Printf("circle %T", shape)
			b.shapes = append(b.shapes[:index], b.shapes[index+1:]...)
			numCircles++
		default:
			fmt.Printf("not circle %T", shape)
			continue
		}
	}
	fmt.Println(numCircles)
	if numCircles == 0 {
		return errors.New("No circles in the box")
	}
	return nil
}
