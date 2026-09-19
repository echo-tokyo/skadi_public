package slices

import (
	goslices "slices"
	"testing"
)

func TestDelDupl(t *testing.T) {
	raw := []int{4, 4, 2, 8, 19, 1, 45, 8}
	expect := []int{1, 2, 4, 8, 19, 45}

	res := DelDupls(raw)
	if goslices.Equal(res, expect) {
		t.Log("OK")
	} else {
		t.Errorf("ERROR. Got slice: %+v", res)
	}
}

func TestDelDupl_NilSlice(t *testing.T) {
	var raw, expect []int

	res := DelDupls(raw)
	if goslices.Equal(res, expect) {
		t.Log("OK")
	} else {
		t.Errorf("ERROR. Got slice: %+v", res)
	}
}
