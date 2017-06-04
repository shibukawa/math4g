package math4g

import (
	"math"
	"math/rand"
	"testing"
)

func prepareTestScala() []Scala {
	tests := make([]Scala, 10000)
	for i := 0; i < 10000; i++ {
		tests[i] = Scala(rand.NormFloat64())
	}
	return tests
}

func prepareTestScala2() []Scala {
	tests := make([]Scala, 10000)
	for i := 0; i < 10000; i++ {
		tests[i] = Scala(rand.ExpFloat64())
	}
	return tests
}

func AbsGoMath(x Scala) Scala {
	return Scala(math.Abs(float64(x)))
}

func TestAbs(t *testing.T) {
	tests := prepareTestScala()
	tests = append(tests, 0.0, 1.0, -1.0)
	for _, v := range tests {
		v1 := Abs(v)
		v2 := AbsGoMath(v)
		if Abs(v1-v2) > 0.001 {
			t.Errorf("Error v=%v v1=%v v2=%v", v, v1, v2)
			return
		}
	}
}

func BenchmarkAbsGoMath(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AbsGoMath(tests[i%10000])
	}
}

func BenchmarkAbs(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Abs(tests[i%10000])
	}
}

func TestClamp(t *testing.T) {
	if Clamp(1.0, -20.0, 20.0) != 1.0 {
		t.Errorf("Clamp(%v, -20.0, 20.0)=%v (should be %v)", 1.0, Clamp(1.0, -20.0, 20.0), 1.0)
	}
	if Clamp(20.0, -20.0, 20.0) != 20.0 {
		t.Errorf("Clamp(%v, -20.0, 20.0)=%v (should be %v)", 1.0, Clamp(20.0, -20.0, 20.0), 20.0)
	}
	if Clamp(25.0, -20.0, 20.0) != 20.0 {
		t.Errorf("Clamp(%v, -20.0, 20.0)=%v (should be %v)", 1.0, Clamp(25.0, -20.0, 20.0), 20.0)
	}
	if Clamp(-20.0, -20.0, 20.0) != -20.0 {
		t.Errorf("Clamp(%v, -20.0, 20.0)=%v (should be %v)", 1.0, Clamp(-20.0, -20.0, 20.0), -20.0)
	}
	if Clamp(-30.0, -20.0, 20.0) != -20.0 {
		t.Errorf("Clamp(%v, -20.0, 20.0)=%v (should be %v)", 1.0, Clamp(-30.0, -20.0, 20.0), -20.0)
	}
}

func TestMin(t *testing.T) {
	if Min(1.0, 2.0) != 1.0 {
		t.Errorf("Min(1.0, 2.0)=%v (should be 1.0)", Min(1.0, 2.0))
	}
	if Min(2.0, 1.0) != 1.0 {
		t.Errorf("Min(2.0, 1.0)=%v (should be 1.0)", Min(2.0, 1.0))
	}
}

func TestMax(t *testing.T) {
	if Max(1.0, 2.0) != 2.0 {
		t.Errorf("Max(1.0, 2.0)=%v (should be 2.0)", Max(1.0, 2.0))
	}
	if Max(2.0, 1.0) != 2.0 {
		t.Errorf("Max(2.0, 1.0)=%v (should be 2.0)", Max(2.0, 1.0))
	}
}

func TestMinMax(t *testing.T) {
	x, y := MinMax(1.0, 2.0)
	if x != 1.0 || y != 2.0 {
		t.Errorf("Min(1.0, 2.0)=%v (should be 1.0)", x, y)
	}
	x, y = MinMax(2.0, 1.0)
	if x != 1.0 || y != 2.0 {
		t.Errorf("Min(2.0, 1.0)=%v (should be 1.0)", x, y)
	}
}

func TestIsNaN(t *testing.T) {
	if IsNaN(1.0) {
		t.Error("IsNaN(1.0) should be false")
	}
	a := 0.0
	b := 0.0
	c := Scala(a/b)
	if !IsNaN(c) {
		t.Error("IsNaN(0.0/0.0) should be true")
	}
}

func IsNaNGoMath(x Scala) bool {
	return math.IsNaN(float64(x))
}

func BenchmarkIsNaNGoMath(b *testing.B) {
	tests := prepareTestScala()
	x := 0.0
	y := 0.0
	tests = append(tests, Scala(x/y))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsNaNGoMath(tests[i%10001])
	}
}

func BenchmarkIsNaN(b *testing.B) {
	tests := prepareTestScala()
	x := 0.0
	y := 0.0
	tests = append(tests, Scala(x/y))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsNaN(tests[i%10001])
	}
}

func TestNaN(t *testing.T) {
	if !IsNaN(NaN()) {
		t.Error("NaN() should return NaN")
	}
}

func NaNGoMath() Scala {
	return Scala(math.NaN())
}

func BenchmarkNaNGoMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaNGoMath()
	}
}

func BenchmarkNaN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaN()
	}
}

func CbrtGoMath(x Scala) Scala {
	return Scala(math.Cbrt(float64(x)))
}

func BenchmarkCbrtGoMath(b *testing.B) {
	tests := prepareTestScala2()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CbrtGoMath(tests[i%10000])
	}
}

func BenchmarkCbrt(b *testing.B) {
	tests := prepareTestScala2()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cbrt(tests[i%10000])
	}
}

func SqrtGoMath(x Scala) Scala {
	return Scala(math.Sqrt(float64(x)))
}

func BenchmarkSqrtGoMath(b *testing.B) {
	tests := prepareTestScala2()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SqrtGoMath(tests[i%10000])
	}
}

func BenchmarkSqrt(b *testing.B) {
	tests := prepareTestScala2()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sqrt(tests[i%10000])
	}
}