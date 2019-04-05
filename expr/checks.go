package expr

import "reflect"

func IsArithmetic(t reflect.Kind) bool {
	switch t {
	case reflect.Int:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint64:
		fallthrough
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		return true
	default:
		return false
	}
}

func IsUnsigned(t reflect.Kind) bool {
	switch t {
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	}
	return false
}
