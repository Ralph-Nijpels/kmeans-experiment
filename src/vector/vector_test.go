package vector

import (
	"fmt"
	"testing"
)

func TestMake3D(t *testing.T) {
	fmt.Printf("Testing Make3D..\n")

	v := make3D([]float64{0.0, 0.0, 0.0})

	r := v.(vector3d)
	if r[0] != 0.0 || r[1] != 0.0 || r[2] != 0.0 {
		t.Errorf("Fail Make3D(0.0, 0.0, 0.0")
	}

	v = make3D([]float64{1.0, 1.0, 1.0})

	r = v.(vector3d)
	if r[0] != 1.0 || r[1] != 1.0 || r[2] != 1.0 {
		t.Errorf("Fail Make3D(1.0, 1.0, 1.0")
	}

	v = make3D([]float64{-1.0, -1.0, -1.0})

	r = v.(vector3d)
	if r[0] != -1.0 || r[1] != -1.0 || r[2] != -1.0 {
		t.Errorf("Fail Make3D(-1.0, -1.0, -1.0")
	}
}

func TestZero3D(t *testing.T) {
	fmt.Printf("Testing Zero3D..\n")

	v := zero3D()

	r := v.(vector3d)
	if r[0] != 0.0 || r[1] != 0.0 || r[2] != 0.0 {
		t.Errorf("Fail Zero3D(0.0, 0.0, 0.0")
	}

}

func TestLen3d(t *testing.T) {
	fmt.Printf("Testing Len3D..\n")

	v := zero3D()
	if v.Len() != 3 {
		t.Errorf("Fail Len3D<Vector>")
	}

	r := v.(vector3d)
	if r.Len() != 3 {
		t.Errorf("Fail Len3D<Vector3D>")
	}
}

func TestRaw3D(t *testing.T) {
	fmt.Printf("Testing Raw3D..\n")

	v := make3D([]float64{1.0, 2.0, 3.0})

	a := v.Raw()
	if a[0] != 1.0 || a[1] != 2.0 || a[2] != 3.0 {
		t.Errorf("Fail Raw3D<Vector>(1.0, 2.0, 3.0)")
	}

	r := v.(vector3d)
	a = r.Raw()
	if a[0] != 1.0 || a[1] != 2.0 || a[2] != 3.0 {
		t.Errorf("Fail Raw3D<Vector3D>(1.0, 2.0, 3.0)")
	}

}

func TestGet3D(t *testing.T) {
	fmt.Printf("Testing Get3D..\n")

	v := make3D([]float64{1.0, 2.0, 3.0})
	if v.Get(0) != 1.0 || v.Get(1) != 2.0 || v.Get(2) != 3.0 {
		t.Errorf("Fail Get3D<Vector>(1.0, 2.0, 3.0)")
	}

	r := v.(vector3d)
	if r.Get(0) != 1.0 || r.Get(1) != 2.0 || r.Get(2) != 3.0 {
		t.Errorf("Fail Get3D<Vector3D>(1.0, 2.0, 3.0)")
	}
}

func TestSet3D(t *testing.T) {
	fmt.Printf("Testing Set3D..\n")

	v := zero3D()
	v = v.Set(0, 1.0)
	v = v.Set(1, 2.0)
	v = v.Set(2, 3.0)

	a := v.Raw()
	if a[0] != 1.0 || a[1] != 2.0 || a[2] != 3.0 {
		t.Errorf("Fail Set3D<Vector>(1.0, 2.0, 3.0): %v", a)
	}

	v = zero3D()
	r := v.(vector3d)
	r = r.Set(0, 1.0).(vector3d)
	r = r.Set(1, 2.0).(vector3d)
	r = r.Set(2, 3.0).(vector3d)

	a = r.Raw()
	if a[0] != 1.0 || a[1] != 2.0 || a[2] != 3.0 {
		t.Errorf("Fail Set3D<Vector3D>(1.0, 2.0, 3.0): %v", a)
	}
}
