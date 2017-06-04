package math4g

import (
	"math"
)

func Abs(x Scala) Scala {
	if x > 0 {
		return x
	}
	return -x
}

func Clamp(value, min, max Scala) Scala {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func Min(x, y Scala) Scala {
	if x < y {
		return x
	}
	return y
}

func Max(x, y Scala) Scala {
	if x < y {
		return y
	}
	return x
}

func MinMax(x, y Scala) (Scala, Scala) {
	if x < y {
		return x, y
	}
	return y, x
}

func Sqrt(x Scala) Scala {
	/* slow implementation: 1
	import "github.com/rkusa/gm/math32"
	return Scala(math32.Sqrt(float32(x)))
	 */
	/* slow implementation: 2
	import "github.com/chewxy/math32"
	return Scala(math32.Sqrt(float32(x)))
	 */
	return Scala(math.Sqrt(float64(x)))
}