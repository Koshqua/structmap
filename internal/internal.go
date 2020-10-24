package internal

import (
	"fmt"
	"reflect"
	"sync"
)



var (
	EmptyValue reflect.Value
	ErrValueNilPointer = fmt.Errorf("value is an empty pointer")
	ErrValueNotStruct = fmt.Errorf("value is not a struct")
	StructMapping sync.Map
)

func RegisterValue(val reflect.Value) error{
	if val.Type().Kind() == reflect.Ptr{
		val = reflect.Indirect(val)
	}
	if val == EmptyValue{
		return ErrValueNilPointer
	}
	if val.Type().Kind() != reflect.Struct{
		return ErrValueNotStruct
	}
	valType := val.Type()

	typeName := valType.Name()
	for i := 0; i < valType.NumField(); i ++{
		field := valType.Field(i)
		tag := field.Tag.Get("structmap")
		if tag == ""{
			continue
		}
		keyName := fmt.Sprintf("%v_%v", typeName, field.Name)
		StructMapping.Store(keyName, tag)
	}
	return nil
}