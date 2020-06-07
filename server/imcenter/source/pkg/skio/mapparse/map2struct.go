package m2p

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

//Error implements the error interface and can represents multiple
//errors that occur in the course of a single decode.
type Error struct {
	Errors []string
}

func (e *Error) Error() string {
	ps := make([]string, len(e.Errors))
	for i, err := range e.Errors {
		ps[i] = fmt.Sprintf("* %s", err)
	}
	sort.Strings(ps)
	return fmt.Sprintf("%d error(s) decoding:\n\n%s", len(e.Errors), strings.Join(ps, "\n"))
}

//WrappedErrors implements the errwrap.Wrapper interface to make this
//return value more useful with the errwrap and go-multierror libraries.
func (e *Error) WrappedErrors() []error {
	if e == nil {
		return nil
	}
	result := make([]error, len(e.Errors))
	for i, e := range e.Errors {
		result[i] = errors.New(e)
	}
	return result
}

func appendErrors(errors []string, err error) []string {
	switch e := err.(type) {
	case *Error:
		return append(errors, e.Errors...)
	default:
		return append(errors, e.Error())
	}
}

func typedDecodeHook(h DecodeHookFunc)

//DecodeHookFunc is the callback function that can be used for
//data transformations. See "DecodeHook" in the DecoderConfig
//struct.
//
//The type should be DecodeHookFuncType or DecodeHookFuncKind.
//Either is accepted. Types are a superset of Kinds (Types can return
//Kinds) and are generally a richer thing to use, but Kinds are simpler
//if you only need those.
//
//The reason DecodeHookFunc is multi-typed is for backwards compatibility:
//we started with Kinds and then realized Types were the better solution,
//but have a promise to not break backwards compat so we now support both.
type DecodeHookFunc interface{}

//DecodeHookFuncType is a DecodeHookFunc which has complete information
//about the source and target types.
type DecodeHookFuncType func(reflect.Type, reflect.Type, interface{}) (interface{}, error)

//DecodeHookFuncKind is a DecodeHookFunc which knows only the Kinds of the
//source and target types
type DecodeHookFuncKind func(reflect.Kind, reflect.Kind, interface{}) (interface{}, error)

//DecoderConfig is the configuration that is used to create a new decoder
//and allows customization of various aspects of decoding.
type DecoderConfig struct {
	//DecodeHook, if set, will be called before any decoding and any
	//type conversion (if WeaklyTypedInput is on). This lets you modify
	//the values before they're set down onto the resulting struct.
	//
	//If an error is returned, the entire decode will fail with that
	//error
	DecodeHook DecodeHookFunc
	//ErrorUnused if is true, the it is an error for there to exist
	//keys in the original map that were unused in the decoding process
	//(extra keys)
	ErrorUnuserd bool
	//ZeroFields, if set to true, will zero fields before writing them.
	//For example, a map will be emptyed before decoded values are put
	//in it. If this is false, a map will be merged.
	ZeroFields bool

	//WeaklyTypedInput, if it is true, the decoder will make the following
	//"weak" conversions:
	//
	//- bools to string (true = "1", false = "0")
	//- numbers to string (base 10)
	WeaklyTypedInput bool
}
