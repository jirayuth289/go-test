package oop

import (
	"fmt"
)

func TestSingleResponsibility() {
	fmt.Println("--------------------------Single Responsibility--------------------------")
	var triangle TriangleCalculator = TriangleCalculator{Base: 2, Height: 2}
	area := triangle.Area()
	fmt.Printf("TriangleCalculator %f\n", area)
}
