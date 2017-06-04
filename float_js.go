// +build js

package math4g

import (
	"math"
	"github.com/gopherjs/gopherjs/js"
)

var (
	nan = Scala(js.Global.Get("NaN").Float())
)

func NaN() Scala {
	return nan
}

func IsNaN(x Scala) bool {
	/* slow implementation: 1
	return js.Global.Get("isNaN").Invoke(x).Bool()
	 */
	return math.IsNaN(float64(x))
}

func Cbrt(x Scala) Scala {
	y := Scala(math.Pow(float64(Abs(x)), 1.0/3.0))
	if x < 0 {
		return -y
	} else {
		return y
	}
}