package oop

type TriangleCalculator struct {
	Base   float32
	Height float32
}

func (t TriangleCalculator) Area() float32 {
	var triangleConstant float32 = 0.5

	return triangleConstant * t.Base * t.Height
}
