package math4g

// Vec2 is a 2 elements vector
type Vec2 [2]Scala

// Length returns length of vector.
func (v Vec2) Length() Scala {
	return Sqrt(v[0] * v[0] + v[1] * v[1])
}

// Normalize changes vector length to 1
func (v *Vec2) Normalize() Vec2 {
	l := v.Length()
	if l != 0 {
		v[0] /= l
		v[1] /= l
	}
	return *v
}

