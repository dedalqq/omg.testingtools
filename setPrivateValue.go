package testingtools

import (
	"fmt"
	"reflect"
	"unsafe"
)

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

	var offset uintptr

	for i := 0; i < reflectValue.NumField(); i++ {
		if reflectType.Field(i).Name == fieldName {
			currentFieldType := reflectType.Field(i).Type.String()

			if newReflectValueType.String() != currentFieldType {
				panic(fmt.Sprintf("incorrect type: must be %s", currentFieldType))
			}

			fieldPtr := unsafe.Pointer(reflectValue.UnsafeAddr() + offset)
			reflect.NewAt(newReflectValueType, fieldPtr).Elem().Set(reflect.ValueOf(value))

			return
		}

		offset += reflectValue.Field(i).Type().Size()
	}

	panic(fmt.Sprintf("Field [%s] not found in struct", fieldName))
}
