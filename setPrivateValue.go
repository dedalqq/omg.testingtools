package testingtools

import (
	"fmt"
	"reflect"
	"unsafe"
)

// SetPrivateValue sets new a value in to the struct. Use it if you need to get access to private
// field of the struct from imported package. important: It function must be using only for the tests.
func SetPrivateValue(obj interface{}, fieldName string, value interface{}) {
	reflectValue := reflect.ValueOf(obj)

	if reflectValue.Kind() != reflect.Ptr {
		panic("object is not a pointer")
	}

	reflectValue = reflectValue.Elem()
	reflectType := reflect.TypeOf(obj).Elem()

	if reflectValue.Kind() != reflect.Struct {
		panic("object is not a pointer to struct")
	}

	newReflectValueType := reflect.TypeOf(value)

	if fieldType, ok := reflectType.FieldByName(fieldName); ok {
		currentFieldType := fieldType.Type.String()

		if newReflectValueType.String() != currentFieldType {
			panic(fmt.Sprintf("incorrect type: must be %s", currentFieldType))
		}

		fieldPtr := unsafe.Pointer(reflectValue.UnsafeAddr() + fieldType.Offset)
		reflect.NewAt(newReflectValueType, fieldPtr).Elem().Set(reflect.ValueOf(value))

		return
	}

	panic(fmt.Sprintf("Field [%s] not found in struct", fieldName))
}
