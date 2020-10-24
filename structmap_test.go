package structmap

import (
	"github.com/koshqua/structmap/internal"
	"testing"
)

type TestStruct struct{
	SomeStr string
	SomeInt int64 `structmap:"some_int"`
}
func TestRegister(t *testing.T) {
	t.Run("should return error if obj is nil pointer", func(t *testing.T){
		err := Register(nil)
		if err == nil {
			t.Errorf("expected to get error %v, got nil ", ErrEmptyValue)

		}
		t.Logf("error is %v",err)
	})
	t.Run("should return error for empty interface", func(t *testing.T){
		var val interface{}
		err := Register(val)
		if err == nil {
			t.Errorf("expected to get error %v, got nil", ErrEmptyValue)
		}
		t.Logf("error is %v", err)
	})
	t.Run("should return error if value is not struct", func(t *testing.T){
		var val = []int{1,2,3,4}
		err := Register(val)
		if err != internal.ErrValueNotStruct{
			t.Logf("expected to get error %v, got %v", internal.ErrValueNotStruct, err)
		}
		t.Logf("error is %v", err)
	})
	t.Run("should register values to map if there are tags", func(t *testing.T){
		err := Register(TestStruct{})
		if err != nil {
			t.Fatalf("didn't expect to get error, got %v", err)
		}
		val, ok := internal.StructMapping.Load("TestStruct_SomeInt")
		if !ok {
			t.Fatal("expected to get ok for map lookup")
		}
		valStr, ok  := val.(string)
		if !ok {
			t.Fatal("expected value to be string")
		}
		if valStr != "some_int"{
			t.Fatalf("expected value to be %v, got %v", "some_int", valStr)
		}
	})
}
