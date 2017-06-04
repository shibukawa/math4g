package math4g

import (
	"testing"
)

func TestVec4Add(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := Vec4{1.0, 2.0, 3.0, 4.0}
	v3 := v1.Add(v2)

	if v3[0] != 2.0 || v3[1] != 4.0 || v3[2] != 6.0 || v3[3] != 8.0 {
		t.Errorf("Vec4.Add error %v", v3)
	}
}

func TestVec4Multiply(t *testing.T) {
	v1 := Vec4{1.0, 2.0, 3.0, 4.0}
	v2 := Vec4{1.0, 2.0, 3.0, 4.0}
	v3 := v1.Multiply(v2)

	if v3[0] != 1.0 || v3[1] != 4.0 || v3[2] != 9.0 || v3[3] != 16.0 {
		t.Errorf("Vec4.Multiply error %v", v3)
	}
}
