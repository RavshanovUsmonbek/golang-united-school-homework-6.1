package golang_united_school_homework

import "errors"

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
	if b.shapesCapacity < len(b.shapes)+1 {
		return errors.New("Not enough capacity to store")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i > (len(b.shapes) - 1) {
		return nil, errors.New("Element doesn't exist in that index")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index out of boundries")
	}
	shape := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if len(b.shapes) <= i {
		return nil, errors.New("index out of boundries")
	}
	removingShape := b.shapes[i]
	b.shapes[i] = shape
	return removingShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	oldLength := len(b.shapes)
	for index, shape := range b.shapes {
		if _, ok := shape.(Circle); ok {
			b.shapes = append(b.shapes[:index], b.shapes[index+1:]...)
		}
	}
	if len(b.shapes) == oldLength {
		return errors.New("No circles in the list")
	}
	return nil
}
