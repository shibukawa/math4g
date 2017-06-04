// +build !js

package math4g

import (
	"github.com/rkusa/gm/math32"
)

func Sin(x Scala) Scala {
	return Scala(math32.Sin(float32(x)))
}

func Cos(x Scala) Scala {
	return Scala(math32.Cos(float32(x)))
}

func Tan(v Scala) Scala {
	return Scala(math32.Tan(float32(v)))
}

func Sincos(v Scala) (Scala, Scala) {
	sin, cos := math32.Sincos(float32(v))
	return Scala(sin), Scala(cos)
}
