package m2s

import (
	"errors"
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
	f := ComposeDecodeHookFunc(f1, f2)

	ret, err := DecodeHookExec(f, reflect.TypeOf(""), reflect.TypeOf([]byte("")), "")
	if err != nil {
		t.Fatalf("bad: %s", err)
	}
	if ret.(string) != "foobar" {
		t.Fatalf("bad: %#v", ret)
	}
}

func TestComposeDecodeHookFunc_err(t *testing.T) {
	f1 := func(reflect.Kind, reflect.Kind, interface{}) (interface{}, error) {
		return nil, errors.New("foo")
	}

	f2 := func(reflect.Kind, reflect.Kind, interface{}) (interface{}, error) {
		panic("NOPE")
	}

	f := ComposeDecodeHookFunc(f1, f2)

	_, err := DecodeHookExec(
		f, reflect.TypeOf(""), reflect.TypeOf([]byte("")), 42,
	)
	if err.Error() != "foo" {
		t.Fatalf("bad: %s", err)
	}
}

func TestComposeDecodeHookFunc_kinds(t *testing.T) {
	var f2From reflect.Kind

	f1 := func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{},
	) (interface{}, error) {
		return int(42), nil
	}

	f2 := func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{},
	) (interface{}, error) {
		f2From = f
		return data, nil
	}

	f := ComposeDecodeHookFunc(f1, f2)

	_, err := DecodeHookExec(
		f,
		reflect.TypeOf(""),
		reflect.TypeOf([]byte("")),
		"",
	)
	if err != nil {
		t.Fatalf("bad: %s", err)
	}
	if f2From != reflect.Int {
		t.Fatalf("bad: %#v", f2From)
	}
}
