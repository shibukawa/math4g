package math4g

import (
	"math"
)

func Acos(v Scala) Scala {
	return Scala(math.Acos(float64(v)))
}

func Atan2(x, y Scala) Scala {
	return Scala(math.Atan2(float64(x), float64(y)))
}

func CosDeg (degrees Scala) Scala {
	return Scala(math.Cos(float64(degrees * AngleToRadian)))
}

func SinDeg (degrees Scala) Scala {
	return Scala(math.Sin(float64(degrees * AngleToRadian)))
}

func SincosDeg (degrees Scala) (s, c Scala) {
	sin, cos := math.Sincos(float64(degrees * AngleToRadian))
	return Scala(sin), Scala(cos)
}

