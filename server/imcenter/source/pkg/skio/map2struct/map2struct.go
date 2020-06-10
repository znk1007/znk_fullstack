package m2s

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
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
	//- bools to int/uint (true = 1, false = 0)
	//- strings to int/uint (base implies by prefix)
	//-	int to bool (true is value != 0)
	//- string to bool (accepts: 1, t, T, TRUE, true, True, 0, f, F,
	//	FALSE, false, False. Anything else is an error)
	//-	empty array = empty map and vice versa
	//-	negative numbers to overflowed uint values (base 10)
	//- slice of maps to a merged map
	//-	single values are converted to slices if required. Each
	//	element is weakly decoded. For example: "4" can become []int{4}
	//	if the target type is an int slice.
	WeaklyTypedInput bool

	//Squash will squash embedded structs. A squash(压缩映射) tag may also
	//be added to an individual struct field using a tag. For example:
	//
	// type Parent struct {
	//		Child `m2s:",squash"`
	//}
	Squash bool

	//Metadata is the struct that will contain extra metadata about
	//the decoding. If this nil, then no metadata will be tracked.
	Metadata *Metadata

	// Result is a pointer to the struct that will contain the decoded value.
	Result interface{}

	//The tag name that map2struct reads for field names.
	//This defaults to "map2struct"
	TagName string
}

//Decoder takes a raw interface value and turns ti into structured data,
//keeping track of rich error information along the way in case anything
//goes wrong. Unlike the basic top-level Decode method, you can more finely
//control how the Decoder behaves using the DecoderConfig struct.
//The top-level Decode method is just a convenience that sets up the most
//basic Decoder.
type Decoder struct {
	config *DecoderConfig
}

//Metadata contains information about decoding a structure that
//is tedious or difficult to get otherwise.
type Metadata struct {

	//Keys are the keys of the struct which were successfully decoded
	Keys []string

	//Unused is a slice of keys that were found in the raw value but
	//weren't decoded since there was no matching field in the result
	//interface
	Unused []string
}

//typedDecodeHook takes a raw DecodeHookFunc (an interface{}) and turns
//it into the proper DecodeHookFunc type, such as DecodeHookFuncType.
func typedDecodeHook(h DecodeHookFunc) DecodeHookFunc {
	//Create variables here so we can reference them with the reflect pkg
	var f1 DecodeHookFuncType
	var f2 DecodeHookFuncKind

	//Fill in the variables into this interface and the rest is done
	//automatically using the reflect package.
	potential := []interface{}{f1, f2}

	v := reflect.ValueOf(h)
	vt := v.Type()
	for _, raw := range potential {
		pt := reflect.ValueOf(raw).Type()
		if vt.ConvertibleTo(pt) {
			return v.Convert(pt).Interface()
		}
	}
	return nil
}

//DecodeHookExec executes the given decode hook. This should be used since
//it'll naturally degrade to the older backwards compatible DecodeHookFunc
//that took reflect.Kind instead of reflect.Type.
func DecodeHookExec(
	raw DecodeHookFunc,
	from, to reflect.Type,
	data interface{}) (interface{}, error) {

	switch f := typedDecodeHook(raw).(type) {
	case DecodeHookFuncType:
		return f(from, to, data)
	case DecodeHookFuncKind:
		return f(from.Kind(), to.Kind(), data)
	default:
		return nil, errors.New("invalid decode hook signature")
	}
}

//ComposeDecodeHookFunc creates a single DecodeHookFunc that
//automatically composes multiple DecodeHookFuncs.
//
//The composed funcs are called in order, with the result of
//the previous transformation.
func ComposeDecodeHookFunc(fs ...DecodeHookFunc) DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		var err error
		for _, f1 := range fs {
			data, err = DecodeHookExec(f1, f, t, data)
			if err != nil {
				return nil, err
			}

			//Modify the from kind to be correct with the new data
			f = nil
			if val := reflect.ValueOf(data); val.IsValid() {
				f = val.Type()
			}
		}
		return data, nil
	}
}

//StrToSliceHookFunc returns a DecodeHookFunc that converts
//string to []string by splitting on the given sep.
func StrToSliceHookFunc(sep string) DecodeHookFunc {
	return func(
		f reflect.Kind,
		t reflect.Kind,
		data interface{},
	) (interface{}, error) {
		if f != reflect.String || t != reflect.Slice {
			return data, nil
		}

		raw := data.(string)
		if len(raw) == 0 {
			return []string{}, nil
		}
		return strings.Split(raw, sep), nil
	}
}

//StrToTimeDurationHookFunc returns a DecodeHookFunc that converts
//strings to time.Duration.
func StrToTimeDurationHookFunc() DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(time.Duration(5)) {
			return data, nil
		}

		//Convert it by parsing
		return time.ParseDuration(data.(string))
	}
}

//StrToIPHookFunc returns a DecodeHookFunc that converts
//strings to net.IP
func StrToIPHookFunc() DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(net.IP{}) {
			return data, nil
		}

		//Convert it by parsing
		ip := net.ParseIP(data.(string))
		if ip == nil {
			return net.IP{}, fmt.Errorf("failed parsing ip %v", data)
		}
		return ip, nil
	}
}

//StrToIPNetHookFunc returns a DecodeHookFunc that converts
//strings to net.IPNet
func StrToIPNetHookFunc() DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(net.IPNet{}) {
			return data, nil
		}
		//Convert it by parsing
		_, net, err := net.ParseCIDR(data.(string))
		return net, err
	}
}

//StrToTimeHookFunc returns a DecodeHookFunc that converts
//strings to time.Time.
func StrToTimeHookFunc(layout string) DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(time.Time{}) {
			return data, nil
		}
		//Convert it by parsing
		return time.Parse(layout, data.(string))
	}
}

//WeaklyTypedHook is a DecodeHookFunc which adds support for weak typing to
//
func WeaklyTypedHook(
	f reflect.Kind,
	t reflect.Kind,
	data interface{},
) (interface{}, error) {
	dataVal := reflect.ValueOf(data)
	switch t {
	case reflect.String:
		switch f {
		case reflect.Bool:
			if dataVal.Bool() {
				return "1", nil
			}
			return "0", nil
		case reflect.Float32:
			return strconv.FormatFloat(dataVal.Float(), 'f', -1, 64), nil
		case reflect.Int:
			return strconv.FormatInt(dataVal.Int(), 10), nil
		case reflect.Slice:
			dataType := dataVal.Type()
			elemKind := dataType.Elem().Kind()
			if elemKind == reflect.Uint8 {
				return string(dataVal.Interface().([]uint8)), nil
			}
		case reflect.Uint:
			return strconv.FormatUint(dataVal.Uint(), 10), nil
		}
	}
	return data, nil
}

//Decode takes an input struct and uses reflection to translate it to
//the output struct. output must be a pointer to a map or struct.
func Decode(input interface{}, output interface{}) error {
	// config := &DecoderConfig{
	// 	Metadata: nil,
	// 	Result:   output,
	// }
	// decoder, err :=
	return nil
}

//NewDecoder returns a new decoder for the given configuration.
//Once a decoder has been returned, the same configuration must not
//be used again.
func NewDecoder(config *DecoderConfig) (*Decoder, error) {
	val := reflect.ValueOf(config.Result)
	if val.Kind() != reflect.Ptr {
		return nil, errors.New("result must be a pointer")
	}

	val = val.Elem()
	if !val.CanAddr() {
		return nil, errors.New("result must be addressable (a pointer)")
	}

	if config.Metadata != nil {
		if config.Metadata.Keys == nil {
			config.Metadata.Keys = make([]string, 0)
		}
		if config.Metadata.Unused == nil {
			config.Metadata.Unused = make([]string, 0)
		}
	}

	if len(config.TagName) == 0 {
		config.TagName = "m2s"
	}

	result := &Decoder{
		config: config,
	}
	return result, nil
}

//decode deco
func (d *Decoder) decode(name string, input interface{}, outVal reflect.Value) error {
	var inputVal reflect.Value
	if input != nil {
		inputVal = reflect.ValueOf(input)
		//We need to check here if input is a typed nil.
		//Typed nils won't match the "input == nil" below so we check
		//that here.
		if inputVal.Kind() == reflect.Ptr && inputVal.IsNil() {
			input = nil
		}
	}

	if input == nil {
		//If the data is nil, then we don't set anything,
		//unless ZeroFields is set to true.
		if d.config.ZeroFields {
			outVal.Set(reflect.Zero(outVal.Type()))

			if d.config.Metadata != nil && len(name) > 0 {
				d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
			}
		}
		return nil
	}

	if !inputVal.IsValid() {
		//If the input value is invalid, then we just set the value
		//to be the zero value.
		outVal.Set(reflect.Zero(outVal.Type()))
		if d.config.Metadata != nil && len(name) > 0 {
			d.config.Metadata.Keys = append(d.config.Metadata.Keys, name)
		}
		return nil
	}

	if d.config.DecodeHook != nil {
		//We have a DecodeHook, so let's pre-process the input.
		// var err error
		// input, err =
	}
	return nil
}
