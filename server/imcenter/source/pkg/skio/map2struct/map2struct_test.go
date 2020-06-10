package map2struct

import (
	"reflect"
	"testing"
)

func TestComposeDecodeHookFunc(t *testing.T) {
	f1 := func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{},
	) (interface{}, error) {
		return data.(string) + "foo", nil
	}

	f2 := func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{},
	) (interface{}, error) {
		return data.(string) + "bar", nil
	}

	f := TestComposeDecodeHookFunc(f1, f2)

	ret, err := DecodeHookExec(
		f, reflect.TypeOf(""),reflect.TypeOf([]byte("")), ""
	)
	if err != nil {
		t.Fatalf("bad: %s", err)
	}
	if ret.(string) != "foobar" {
		
	}
}
