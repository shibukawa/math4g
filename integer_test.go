package math4g

import (
	"math"
	"testing"
)

func FloorGoMath(x Scala) Scala {
	return Scala(math.Floor(float64(x)))
}

func TestFloor(t *testing.T) {
	tests := prepareTestScala()
	for _, v := range tests {
		v1 := Floor(v)
		v2 := FloorGoMath(v)
		if Abs(v1-v2) > 0.001 {
			t.Errorf("Error v=%v v1=%v v2=%v", v, v1, v2)
			return
		}
	}
}

func BenchmarkFloorGoMath(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FloorGoMath(tests[i%10000])
	}
}

func BenchmarkFloor(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Floor(tests[i%10000])
	}
}

func CeilGoMath(x Scala) Scala {
	return Scala(math.Ceil(float64(x)))
}

func TestCeil(t *testing.T) {
	tests := prepareTestScala()
	tests = append(tests, 0.0, 1.0, -1.0)
	for _, v := range tests {
		v1 := Ceil(v)
		v2 := CeilGoMath(v)
		if Abs(v1-v2) > 0.001 {
			t.Errorf("Error v=%v v1=%v v2=%v", v, v1, v2)
			return
		}
	}
}

func BenchmarkCeilGoMath(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CeilGoMath(tests[i%10000])
	}
}

func BenchmarkCeil(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Ceil(tests[i%10000])
	}
}

func TruncGoMath(x Scala) Scala {
	return Scala(math.Trunc(float64(x)))
}

func TestTrunc(t *testing.T) {
	tests := prepareTestScala()
	tests = append(tests, 0.0, 1.0, -1.0)
	for _, v := range tests {
		v1 := Trunc(v)
		v2 := TruncGoMath(v)
		if Abs(v1-v2) > 0.001 {
			t.Errorf("Error v=%v v1=%v v2=%v", v, v1, v2)
			return
		}
	}
}

func BenchmarkTruncGoMath(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TruncGoMath(tests[i%10000])
	}
}

func BenchmarkTrunc(b *testing.B) {
	tests := prepareTestScala()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Trunc(tests[i%10000])
	}
}

func TestNorm(t *testing.T) {
	if Norm(-20.0) != -1.0 {
		t.Errorf("Error v=%v r=%v", -20.0, Norm(-20.0))
	}
	if Norm(20.0) != 1.0 {
		t.Errorf("Error v=%v r=%v", 20.0, Norm(20.0))
	}
	if Norm(0.0) != 0.0 {
		t.Errorf("Error v=%v r=%v", 0.0, Norm(0.0))
	}
}

