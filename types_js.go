// +build js

package math4g

import (
	"math"
)

// Scala is a type of element of vector and matrix.
// Scala is a float32 on regular environment, and float64 for Gopher.js
// (https://github.com/gopherjs/gopherjs#performance-tips).
type Scala float64

const (
	// Pi is a constant of Pi value in Scala type.
	Pi            Scala = Scala(math.Pi)
	Pi_D                = Pi * 2.0
	Pi_H                = Pi * 0.5
	Pi_Q                = Pi * 0.25
	AngleToRadian       = Pi / 180.0
	RadianToAngle       = 180 / Pi
	MaxValue            = Scala(math.MaxFloat32)

	IsJS = true
)

var (
	PositiveInfinity = Scala(math.Inf(1))
	NegativeInfinity = Scala(math.Inf(-1))
)
