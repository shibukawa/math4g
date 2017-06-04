package math4g

import (
	"testing"
	"math/rand"
)

func TestVec2(t *testing.T) {
	v := Vec2{}
	if v.Length() > 0.0 {
		t.Errorf("Vec2 Length Error expected: 0 actual: %f", v.Length())
	}
	v = Vec2{3.0, 4.0}
	if v.Length() < 4.9 || v.Length() > 5.1 {
		t.Errorf("Vec2 Length Error expected: 5 actual: %f", v.Length())
	}
	v.Normalize()
	if v.Length() < 0.9 || v.Length() > 1.1 {
		t.Errorf("Vec2 Length Error expected: 1 actual: %f", v.Length())
	}
}

func TestMat32(t *testing.T) {
	a := NewMat32(1, 2, 3, 4, 5, 6)
	b := a.ToMat34()
	if b[0] != 1 || b[1] != 2 || b[4] != 3 || b[5] != 4 || b[8] != 5 || b[9] != 6 {
		t.Errorf("NewMat32 creates wrong matrix: %v", a)
	}
}

func TestTranslateMatrix(t *testing.T) {
	// move 10, 20
	translate := TranslateMat32(10, 20)
	x, y := translate.TransformPoint(5, 5)
	if int(x) != 15 || int(y) != 25 {
		t.Errorf("Wrong Transform: %v %f %f", translate, x, y)
	}
}

func TestRotateMatrix(t *testing.T) {
	// rotate 90 degree
	rotate := RotateMat32(Pi * 0.5)
	x, y := rotate.TransformPoint(10, 10)
	if int(x) != -10 || int(y) != 10 {
		t.Errorf("Wrong Rotation: %v %f %f", rotate, x, y)
	}
}

func TestScaleMatrix(t *testing.T) {
	// scale x1.5 for x-axis, x2.5 for y-axis
	scale := ScaleMat32(1.5, 2.5)
	x, y := scale.TransformPoint(10, 10)
	if int(x) != 15 || int(y) != 25 {
		t.Errorf("Wrong Scale: %v %f %f", scale, x, y)
	}
}

func TestMatrix(t *testing.T) {
	translate := TranslateMat32(10, 20)
	rotate := RotateMat32(Pi * 0.5)
	scale := ScaleMat32(1.5, 2.5)
	x, y := translate.Multiply(rotate).Multiply(scale).TransformPoint(10, 10)
	if int(x) != -45 || int(y + 0.5) != 50 {
		t.Errorf("Wrong Affine Transformation: %f %f", x, y)
	}
}

//
func BenchmarkMat32_Multiply(b *testing.B) {
	matrices := make([]Mat32, 2000)
	for i := 0; i < 2000; i++ {
		matrices[i] = Mat32{Scala(rand.Float32()), Scala(rand.Float32()), Scala(rand.Float32()), Scala(rand.Float32()), Scala(rand.Float32()), Scala(rand.Float32())}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrices[i % 1000].Multiply(matrices[i % 1000 + 1000])
	}
}

// Benchmark mutable/immutable for design decision

type Mat32_32 [6]float32

func NewMat32_32() *Mat32_32 {
	result := &Mat32_32{}
	for i := 0; i < 6; i++ {
		result[i] = rand.Float32()
	}
	return result
}

func (t Mat32_32) CopyMultiply(s Mat32_32) Mat32_32 {
	t0 := t[0]*s[0] + t[1]*s[2]
	t2 := t[2]*s[0] + t[3]*s[2]
	t4 := t[4]*s[0] + t[5]*s[2] + s[4]
	t[1] = t[0]*s[1] + t[1]*s[3]
	t[3] = t[2]*s[1] + t[3]*s[3]
	t[5] = t[4]*s[1] + t[5]*s[3] + s[5]
	t[0] = t0
	t[2] = t2
	t[4] = t4
	return t
}

func (t *Mat32_32) InlineMultiply(s *Mat32_32) *Mat32_32 {
	t0 := t[0]*s[0] + t[1]*s[2]
	t2 := t[2]*s[0] + t[3]*s[2]
	t4 := t[4]*s[0] + t[5]*s[2] + s[4]
	t[1] = t[0]*s[1] + t[1]*s[3]
	t[3] = t[2]*s[1] + t[3]*s[3]
	t[5] = t[4]*s[1] + t[5]*s[3] + s[5]
	t[0] = t0
	t[2] = t2
	t[4] = t4
	return t
}

type Mat32_64 [6]float64

func NewMat32_64() *Mat32_64 {
	result := &Mat32_64{}
	for i := 0; i < 6; i++ {
		result[i] = rand.Float64()
	}
	return result
}

func (t Mat32_64) CopyMultiply(s Mat32_64) Mat32_64 {
	t0 := t[0]*s[0] + t[1]*s[2]
	t2 := t[2]*s[0] + t[3]*s[2]
	t4 := t[4]*s[0] + t[5]*s[2] + s[4]
	t[1] = t[0]*s[1] + t[1]*s[3]
	t[3] = t[2]*s[1] + t[3]*s[3]
	t[5] = t[4]*s[1] + t[5]*s[3] + s[5]
	t[0] = t0
	t[2] = t2
	t[4] = t4
	return t
}

func (t *Mat32_64) InlineMultiply(s *Mat32_64) *Mat32_64 {
	t0 := t[0]*s[0] + t[1]*s[2]
	t2 := t[2]*s[0] + t[3]*s[2]
	t4 := t[4]*s[0] + t[5]*s[2] + s[4]
	t[1] = t[0]*s[1] + t[1]*s[3]
	t[3] = t[2]*s[1] + t[3]*s[3]
	t[5] = t[4]*s[1] + t[5]*s[3] + s[5]
	t[0] = t0
	t[2] = t2
	t[4] = t4
	return t
}


func BenchmarkMat32_32_InlineMultiply(b *testing.B) {
	matrices := make([]*Mat32_32, 2000)
	for i := 0; i < 2000; i++ {
		matrices[i] = NewMat32_32()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrices[i % 1000].InlineMultiply(matrices[i % 1000 + 1000])
	}
}

func BenchmarkMat32_32_CopyMultiply(b *testing.B) {
	matrices := make([]Mat32_32, 2000)
	for i := 0; i < 2000; i++ {
		matrices[i] = *NewMat32_32()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrices[i % 1000].CopyMultiply(matrices[i % 1000 + 1000])
	}
}

func BenchmarkMat32_64_InlineMultiply(b *testing.B) {
	matrices := make([]*Mat32_64, 2000)
	for i := 0; i < 2000; i++ {
		matrices[i] = NewMat32_64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrices[i % 1000].InlineMultiply(matrices[i % 1000 + 1000])
	}
}


func BenchmarkMat32_64_CopyMultiply(b *testing.B) {
	matrices := make([]Mat32_64, 2000)
	for i := 0; i < 2000; i++ {
		matrices[i] = *NewMat32_64()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrices[i % 1000].CopyMultiply(matrices[i % 1000 + 1000])
	}
}

