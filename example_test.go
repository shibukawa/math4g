// +build !js

package math4g_test

import (
	"fmt"
	"github.com/shibukawa/math4g"
	"math"
)

func ExampleTranslateMatrix() {
	// move 10, 20
	translate := math4g.TranslateMat32(10, 20)
	x, y := translate.TransformPoint(5, 5)
	fmt.Println(x, y)
	// Output: 15 25
}

func ExampleRotateMatrix() {
	// rotate 90 degree
	rotate := math4g.RotateMat32(math4g.Pi * 0.5)
	x, y := rotate.TransformPoint(10, 10)
	fmt.Println(x, y)
	// Output: -10 10
}

func ExampleScaleMatrix() {
	// scale x1.5 for x-axis, x2.5 for y-axis
	scale := math4g.ScaleMat32(1.5, 2.5)
	x, y := scale.TransformPoint(10, 10)
	fmt.Println(x, y)
	// Output: 15 25
}

func ExampleMatrix() {
	translate := math4g.TranslateMat32(10, 20)
	rotate := math4g.RotateMat32(math4g.Pi * 0.5)
	scale := math4g.ScaleMat32(1.5, 2.5)
	x, y := translate.Multiply(rotate).Multiply(scale).TransformPoint(10, 10)
	fmt.Println(math.Floor(float64(x) + 0.5), math.Floor(float64(y) + .5))
	// Output: -45 50
}