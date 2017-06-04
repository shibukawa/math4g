// +build !js

package math4g

import (
	"math"
)

const (
	uvnan uint32 = 0x7F800001
)

func NaN() Scala {
	return Scala(math.Float32frombits(uvnan))
}

func IsNaN(x Scala) bool {
	return x != x
}

func Cbrt(x Scala) Scala {
	return Scala(math.Cbrt(float64(x)))
}
