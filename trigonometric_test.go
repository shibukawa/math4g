package math4g

import (
	"math"
	"testing"
)

func SinGoMath(x Scala) Scala {
	return Scala(math.Sin(float64(x)))
}

func BenchmarkSinGoMath(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SinGoMath(tests[i%10000])
	}
}

func BenchmarkSin(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sin(tests[i%10000])
	}
}

func CosGoMath(x Scala) Scala {
	return Scala(math.Cos(float64(x)))
}

func BenchmarkCosGoMath(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CosGoMath(tests[i%10000])
	}
}

func BenchmarkCos(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cos(tests[i%10000])
	}
}

func TanGoMath(x Scala) Scala {
	return Scala(math.Tan(float64(x)))
}

func BenchmarkTanGoMath(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TanGoMath(tests[i%10000])
	}
}

func BenchmarkTan(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Tan(tests[i%10000])
	}
}

func SincosGoMath(x Scala) (Scala, Scala) {
	sin, cos := math.Sincos(float64(x))
	return Scala(sin), Scala(cos)
}

func BenchmarkSincosGoMath(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SincosGoMath(tests[i%10000])
	}
}

func BenchmarkSincos(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sincos(tests[i%10000])
	}
}
