// +build js

package math4g

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	sin = js.Global.Get("Math").Get("sin")
	cos = js.Global.Get("Math").Get("cos")
	tan = js.Global.Get("Math").Get("tan")
)

func Sin(x Scala) Scala {
	return Scala(sin.Invoke(x).Float())
}

func Cos(x Scala) Scala {
	return Scala(cos.Invoke(x).Float())
}

func Tan(x Scala) Scala {
	return Scala(tan.Invoke(x).Float())
}

func Sincos(x Scala) (s, c Scala) {
	/* slow implementation: 1
	sin, cos := math.Sincos(float64(x))
	return Scala(sin), Scala(cos)
	*/
	return Scala(sin.Invoke(x).Float()), Scala(cos.Invoke(x).Float())
}