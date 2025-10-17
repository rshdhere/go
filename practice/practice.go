package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c *Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := Circle{radius: 5.0}
	result := c.Circumference()
	answer := c.Area()
	fmt.Println(result)
	fmt.Println(answer)
}
