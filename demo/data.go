/*
Package demo contains a model of charged particles and algorithms to manipilate them
*/
package demo

import "fmt"

// Vector is a simple 2d vector
type Vector struct {
	X, Y float64
}

// Particle contains the location, velocity, and charge of a particle
type Particle struct {
	Pos, Vel Vector
	Charge   float64
}

// World is a collection of Particles with various operations upon them
type World []Particle

// Add one vector to the given vector: in-place
func (v *Vector) Add(offset *Vector) {
	v.X += offset.X
	v.Y += offset.Y
}

// Mult creates a new vector as a multiple of the original
func (v *Vector) Mult(mult float64) *Vector {
	return &Vector{mult * v.X, mult * v.Y}
}

// Equals returns true if two vectors have the same contents
func (v *Vector) Equals(cmp *Vector) bool {
	return fEqual(v.X, cmp.X) && fEqual(v.Y, cmp.Y)
}

func (v *Vector) String() string {
	return fmt.Sprintf("<%f, %f>", v.X, v.Y)
}

// Move applies the current velocity for a given time to update the Particle's position
func (p *Particle) Move(step float64) {
	p.Pos.Add(p.Vel.Mult(step))
}
