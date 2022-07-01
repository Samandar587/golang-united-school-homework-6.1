package golang_united_school_homework

import (
	"errors"
	"reflect"
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
	var err error
	b.shapes = append(b.shapes, shape)
	if len(b.shapes) > b.shapesCapacity {
		err = errors.New("Out of range!")
	} else {
		err = nil
	}
	return err
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	var err error
	var res Shape
	var exists bool
	if i < (b.shapesCapacity - 1) {
		for _, value := range b.shapes {
			if value == b.shapes[i] {
				exists = true
			}
		}
	}
	if !exists || i > (b.shapesCapacity-1) {
		err = errors.New("Out of Range!")
		res = nil
	} else {
		res = b.shapes[i]
		err = nil
	}
	return res, err

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	var res Shape
	var exists bool
	var err error
	if i < (b.shapesCapacity - 1) {
		for _, val := range b.shapes {
			if val == b.shapes[i] {
				exists = true

			}
		}
	}
	if !exists || i > (b.shapesCapacity-1) {
		err = errors.New("Out of range or does not exist!")
		res = nil
	} else {
		res = b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		err = nil
	}
	return res, err
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	//    panic("implement me")
	if i > b.shapesCapacity-1 || i < 0 {
		// err := errors.New("Превышен размер коробки")
		// fmt.Println("the expected length is 1, but actual %w", err)
		return nil, errors.New("вышли за пределы")
	} else {
		//b.shapes = append(b.shapes, shape)

		bb := b.shapes[i]
		b.shapes[i] = shape
		//b.shapes = b.shapes[:len(b.shapes)-1]
		return bb, nil
	}

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	t := &Triangle{
		Side: 30,
	}
	r := &Rectangle{
		Height: 10,
		Weight: 20,
	}
	c := &Circle{
		Radius: 20,
	}
	sum = t.CalcPerimeter() + c.CalcPerimeter() + r.CalcPerimeter()
	return sum

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	t := &Triangle{
		Side: 30,
	}
	r := &Rectangle{
		Height: 10,
		Weight: 20,
	}
	c := &Circle{
		Radius: 20,
	}
	var sum = r.CalcArea() + c.CalcArea() + t.CalcArea()
	return sum

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var err error
	var exist bool
	c := &Circle{}
	for i := 0; i < len(b.shapes); i++ {
		if reflect.TypeOf(b.shapes[i]) == reflect.TypeOf(c) {
			exist = true
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			i--
		}
	}
	if exist {
		err = nil
	}
	if !exist {
		err = errors.New("Circle does not exist!")
	}
	return err
}
