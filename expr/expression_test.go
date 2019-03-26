package expr

import (
	"reflect"
	"testing"
)

func TestIsArithmetic(t *testing.T) {
	for _, test := range []struct {
		kind     reflect.Kind
		expected bool
	}{
		{reflect.Int, true},
		{reflect.Uint, true},
		{reflect.Int8, true},
		{reflect.Uint8, true},
		{reflect.Int16, true},
		{reflect.Uint16, true},
		{reflect.Int32, true},
		{reflect.Uint32, true},
		{reflect.Int64, true},
		{reflect.Uint64, true},
		{reflect.Float32, true},
		{reflect.Float64, true},
		{reflect.Complex64, false},
		{reflect.Complex128, false},
		{reflect.Array, false},
		{reflect.Chan, false},
		{reflect.Func, false},
		{reflect.Interface, false},
		{reflect.Map, false},
		{reflect.Ptr, false},
		{reflect.Slice, false},
		{reflect.String, false},
		{reflect.Struct, false},
		{reflect.UnsafePointer, false},
	} {
		if actual := IsArithmetic(test.kind); actual != test.expected {
			t.Fatalf("expected %v but got %v", test.expected, actual)
		}
	}
}
