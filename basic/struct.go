package main

import (
    "fmt"
    "math"
)


type Point struct {
    X, Y float64
}

func (p *Point) Display() {
    fmt.Println("x = ", p.X, " y = ", p.Y)
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
    center.Display()

    circle := Circle{Center: center, Radius: 5.0}
    right := circle.Surface(0.0)
    right.Display()

    bottomLeft := circle.Surface(math.Pi * 1.25)
    bottomLeft.Display()

    down := circle.Surface(math.Pi * 1.5)
    down.Display()

    fmt.Println("   moving...")
    // now, radius 10, centered at (0, 10)
    circle.Expand(2.0)
    circle.Translate(&Point{-4.0, 5.0})
    newDown := circle.Surface(math.Pi * 1.5)
    // should be 0, 0
    newDown.Display()

    // should be 10, 10
    newRight := circle.Surface(0.0)
    newRight.Display()
}
