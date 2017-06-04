package math4g

// Mat32 is a 3x2 matrix is represented as float[6].
type Mat32 [6]Scala

// NewMat32 creates Mat32 instance
func NewMat32(a, b, c, d, e, f Scala) Mat32 {
	return Mat32{a, b, c, d, e, f}
}

// IdentityMatrix makes the transform to identity matrix.
func IdentityMat32() Mat32 {
	return NewMat32(1.0, 0.0, 0.0, 1.0, 0.0, 0.0)
}

// TranslateMatrix makes the transform to translation matrix matrix.
func TranslateMat32(tx, ty Scala) Mat32 {
	return NewMat32(1.0, 0.0, 0.0, 1.0, tx, ty)
}

// ScaleMatrix makes the transform to scale matrix.
func ScaleMat32(sx, sy Scala) Mat32 {
	return NewMat32(sx, 0.0, 0.0, sy, 0.0, 0.0)
}

// RotateMatrix makes the transform to rotate matrix. Angle is specified in radians.
func RotateMat32(a Scala) Mat32 {
	sin, cos := Sincos(a)
	return NewMat32(Scala(cos), Scala(sin), Scala(-sin), Scala(cos), 0.0, 0.0)
}

// SkewXMatrix makes the transform to skew-x matrix. Angle is specified in radians.
func SkewXMat32(a Scala) Mat32 {
	return NewMat32(1.0, 0.0, Tan(a), 1.0, 0.0, 0.0)
}

// SkewYMatrix makes the transform to skew-y matrix. Angle is specified in radians.
func SkewYMat32(a Scala) Mat32 {
	return NewMat32(1.0, Tan(a), 0.0, 1.0, 0.0, 0.0)
}

// Multiply makes the transform to the result of multiplication of two transforms, of A = A*B.
func (t Mat32) Multiply(s Mat32) Mat32 {
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

// PreMultiply makes the transform to the result of multiplication of two transforms, of A = B*A.
func (t Mat32) PreMultiply(s Mat32) Mat32 {
	return s.Multiply(t)
}

// InlinePreMultiply makes the transform to the result of multiplication of two transforms, of A = B*A.
func (s *Mat32) InlinePreMultiply(t *Mat32) {
	s0 := t[0]*s[0] + t[1]*s[2]
	s2 := t[2]*s[0] + t[3]*s[2]
	s4 := t[4]*s[0] + t[5]*s[2] + s[4]
	s[1] = t[0]*s[1] + t[1]*s[3]
	s[3] = t[2]*s[1] + t[3]*s[3]
	s[5] = t[4]*s[1] + t[5]*s[3] + s[5]
	s[0] = s0
	s[2] = s2
	s[4] = s4
}

// Inverse makes the destination to inverse of specified transform.
// Returns 1 if the inverse could be calculated, else 0.
func (t Mat32) Inverse() Mat32 {
	det := t[0]*t[3] - t[2]*t[1]
	if det > -1e-6 && det < 1e-6 {
		return IdentityMat32()
	}
	invdet := 1.0 / det
	return Mat32{
		t[3] * invdet,
		-t[1] * invdet,
		-t[2] * invdet,
		t[0] * invdet,
		(t[2]*t[5] - t[3]*t[4]) * invdet,
		(t[1]*t[4] - t[0]*t[5]) * invdet,
	}
}

// TransformPoint transforms a point by given Transform.
func (t Mat32) TransformPoint(sx, sy Scala) (dx, dy Scala) {
	dx = sx*t[0] + sy*t[2] + t[4]
	dy = sx*t[1] + sy*t[3] + t[5]
	return
}

// Transform transforms a pointer that is represented by vector by given Matrix.
func (t *Mat32) Transform(v Vec2) Vec2 {
	return Vec2{
		v[0]*t[0] + v[1]*t[2] + t[4],
		v[0]*t[1] + v[1]*t[3] + t[5],
	}
}

// ToMat3x4 makes 3x4 matrix.
func (t Mat32) ToMat34() []Scala {
	return []Scala{
		t[0], t[1], 0.0, 0.0,
		t[2], t[3], 0.0, 0.0,
		t[4], t[5], 1.0, 0.0,
	}
}

func (t Mat32) getAverageScale() Scala {
	sx := Sqrt(t[0]*t[0] + t[2]*t[2])
	sy := Sqrt(t[1]*t[1] + t[3]*t[3])
	return (sx + sy) * 0.5
}

