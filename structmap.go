package structmap

import (
	"fmt"
	"github.com/koshqua/structmap/internal"
	"reflect"
)

var (
	EmptyValue reflect.Value
	ErrEmptyValue = fmt.Errorf("value is empty or nil, impossible to register")
)


func Register(obj interface{}) error{
	objValue := reflect.ValueOf(obj)
	if objValue == EmptyValue{
		return ErrEmptyValue
	}
	return internal.RegisterValue(objValue)
}

func Decode (from map[string]interface{}, to interface{}){

}