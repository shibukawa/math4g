// +build !386,!amd64,!amd64p32

package math4g

type Vec4 [4]Scala

func (l Vec4) Add(r Vec4) Vec4 {
	l[0] += r[0]
	l[1] += r[1]
	l[2] += r[2]
	l[3] += r[3]
	return l
}

func (l Vec4) Multiply(r Vec4) Vec4 {
	l[0] *= r[0]
	l[1] *= r[1]
	l[2] *= r[2]
	l[3] *= r[3]
	return l
}
