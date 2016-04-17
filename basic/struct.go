package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("<Point: %v, %v>", p.X, p.Y)
}

func (p *Point) AddVector(p2 *Point) {
	p.X += p2.X
	p.Y += p2.Y
}

type Circle struct {
	Center Point
	Radius float64
}

func (c *Circle) Surface(phi float64) *Point {
	x := c.Center.X + (c.Radius * math.Cos(phi))
	y := c.Center.Y + (c.Radius * math.Sin(phi))
	return &Point{x, y}
}

func (c *Circle) Expand(scale float64) {
	c.Radius *= scale
}

func (c *Circle) Translate(offset *Point) {
	c.Center.AddVector(offset)
}

func main() {
	center := Point{4.0, 5.0}
	// why doesn't this work with the stinger interface?
	fmt.Println(center)
	fmt.Println(&center)

	circle := Circle{Center: center, Radius: 5.0}
	right := circle.Surface(0.0)
	fmt.Println(right)

	bottomLeft := circle.Surface(math.Pi * 1.25)
	fmt.Println(bottomLeft)

	down := circle.Surface(math.Pi * 1.5)
	fmt.Println(down)

	fmt.Println("   moving...")
	// now, radius 10, centered at (0, 10)
	circle.Expand(2.0)
	circle.Translate(&Point{-4.0, 5.0})
	newDown := circle.Surface(math.Pi * 1.5)
	// should be 0, 0
	fmt.Println(newDown)

	// should be 10, 10
	newRight := circle.Surface(0.0)
	fmt.Println(newRight)
}
