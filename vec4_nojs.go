// +build 386 amd64 amd64p32

package math4g

// Vec4 is a vector with 4 elements.
// You can use it for color or something.
type Vec4 [4]Scala

// Add adds two vectors
func (l Vec4) Add(r Vec4) Vec4 {
	addSIMD(&l, &r)
	return l
}

func addSIMD(lhs, rhs *Vec4)

func (l Vec4) Multiply(r Vec4) Vec4 {
	mulSIMD(&l, &r)
	return l
}

func mulSIMD(lhs, rhs *Vec4)
