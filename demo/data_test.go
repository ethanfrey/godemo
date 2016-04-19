/*
Test code for data.go
*/
package demo

import "testing"

func TestAddVector(t *testing.T) {
	cases := [][]Vector{
		[]Vector{Vector{1, 2}, Vector{2, 3}, Vector{3, 5}},
		[]Vector{Vector{1.2, 2.1}, Vector{2.5, 3.2}, Vector{3.7, 5.3}},
	}
	for _, args := range cases {
		a := args[0]
		a.Add(&args[1])
		if !a.Equals(&args[2]) {
			t.Error("Unexpected vector add:", &a, &args[2])
		}
	}
}

func TestMultVector(t *testing.T) {
	type testData struct {
		input    Vector
		mult     float64
		expected Vector
	}

	// test case here
	doTest := func(args testData) {
		res := args.input.Mult(args.mult)
		if !res.Equals(&args.expected) {
			t.Error("Unexpected vector add:", res, &args.expected, res.Y-args.expected.Y)
		}
	}

	// test data here
	cases := []testData{
		testData{Vector{1, 2}, 0.1, Vector{0.1, 0.2}},
		testData{Vector{-7.3, 13.2}, 1.2, Vector{-8.76, 15.84}},
	}

	// and now check all cases...
	for _, args := range cases {
		doTest(args)
	}
}

func TestParticleMove(t *testing.T) {
	type testData struct {
		pos, vel, expected Vector
		steps              int
		delta              float64
	}

	doTest := func(args testData) {
		p := Particle{args.pos, args.vel, 0}
		for i := 0; i < args.steps; i++ {
			p.Move(args.delta)
		}
		if !p.Pos.Equals(&args.expected) {
			t.Error("Unexpected position:", &p.Pos, &args.expected)
		}
	}

	// test data here
	cases := []testData{
		testData{Vector{1, 2}, Vector{0.1, 0.2}, Vector{1.2, 2.4}, 20, 0.1},
		testData{Vector{-7.3, 13.2}, Vector{1.2, -3.4}, Vector{-0.1, -7.2}, 120, 0.05},
	}

	// and now check all cases...
	for _, args := range cases {
		doTest(args)
	}
}
