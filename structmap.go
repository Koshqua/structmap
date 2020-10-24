package structmap

import (
	"fmt"
	"github.com/koshqua/structmap/internal"
	"reflect"
)

var (
	EmptyValue reflect.Value
	ErrEmptyValue = fmt.Errorf("value is empty or nil")
	ErrNotPointer = fmt.Errorf("non pointer value passed to a function")
)


func Register(obj interface{}) error{
	objValue := reflect.ValueOf(obj)
	if objValue == EmptyValue{
		return ErrEmptyValue
	}
	return internal.RegisterValue(objValue)
}

func Decode (from map[string]interface{}, to interface{}) error{
	val := reflect.ValueOf(to)
	if val == EmptyValue{
		return ErrEmptyValue
	}
	if val.Type().Kind() != reflect.Ptr{
		return ErrNotPointer
	}
	ptr := val
	val = reflect.Indirect(ptr)
	typeOF := val.Type()
	if typeOF.Kind() != reflect.Struct{
		return internal.ErrValueNotStruct
	}
	typeName := typeOF.Name()
	for i := 0; i < typeOF.NumField(); i++{
		fmt.Printf("\niteration %v", i)
		typeField := fmt.Sprintf("%v_%v", typeName, typeOF.Field(i).Name)
		value, ok := internal.GetValueFromMap(typeField)
		if !ok {
			continue
		}
		tagStr, ok := value.(string)
		if !ok{
			fmt.Printf("\nvalue in map is not string")
			continue
		}
		realValue, ok := from[tagStr]
		if !ok {
			fmt.Printf("\nValue is not in a map")
			continue
		}
		fmt.Printf("\ngot value %#+v", realValue)
		typeOfValue := reflect.TypeOf(realValue)
		typeOfField := typeOF.Field(i).Type
		if typeOfValue.AssignableTo(typeOfField){
			fmt.Printf("\n value is assignable to struct")
			val.Field(i).Set(reflect.ValueOf(realValue))
		}
	}


	return nil
}