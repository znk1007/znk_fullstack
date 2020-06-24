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
		var err error
		input, err = DecodeHookExec(
			d.config.DecodeHook,
			inputVal.Type(),
			outVal.Type(),
			input,
		)
		if err != nil {
			return fmt.Errorf("error decoding '%s': %s", name, err)
		}
	}
	// var err error
	outputKind := getKind(outVal)
	// addMetaKey := true
	switch outputKind {
	case reflect.Bool:
		// err = d.
	}
	return nil
}

func (d *Decoder) decodeMapFromMap(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
	valType := val.Type()
	valKeyType := valType.Key()
	valElemType := valType.Elem()

	//Accumulate errors
	errors := make([]string, 0)

	//If the input data is empty, then we just match what the input data is.
	if dataVal.Len() == 0 {
		if dataVal.IsNil() {
			if !val.IsNil() {
				val.Set(dataVal)
			}
		} else {
			//Set to empty allocated value
			val.Set(valMap)
		}
		return nil
	}

	for _, k := range dataVal.MapKeys() {
		fieldName := fmt.Sprintf("%s[%s]", name, k)

		//First decode the key into the proper type
		curKey := reflect.Indirect(reflect.New(valKeyType))
		if err := d.decode(fieldName, k.Interface(), curKey); err != nil {
			errors = appendErrors(errors, err)
			continue
		}

		//Next decode the data into the proper type
		v := dataVal.MapIndex(k).Interface()
		curVal := reflect.Indirect(reflect.New(valElemType))
		if err := d.decode(fieldName, v, curVal); err != nil {
			errors = appendErrors(errors, err)
			continue
		}
		valMap.SetMapIndex(curKey, curVal)
	}

	//Set the built up map to the value
	val.Set(valMap)

	//If we had errors, return those
	if len(errors) > 0 {
		return &Error{errors}
	}
	return nil
}

func (d *Decoder) decodeMapFromStruct(name string, dataVal reflect.Value, val reflect.Value, valMap reflect.Value) error {
	t := dataVal.Type()
	for i := 0; i < t.NumField(); i++ {
		//Get the StructField first since this is a cheap operation. If the
		//field is unexported, then ignore it.
		f := t.Field(i)
		if len(f.PkgPath) > 0 {
			continue
		}

		//Next get the actual value of this field and verify it is assignable
		//to the map value.
		v := dataVal.Field(i)
		if !v.Type().AssignableTo(valMap.Type().Elem()) {
			return fmt.Errorf("cannot assign type '%s' to map value field of type '%s'", v.Type(), valMap.Type().Elem())
		}

		tagValue := f.Tag.Get(d.config.TagName)
		keyName := f.Name

		//If Squash is set in the config, we squash the field down.
		squash := d.config.Squash && v.Kind() == reflect.Struct && f.Anonymous
		//Determine the name of the key in the map
		if index := strings.Index(tagValue, ","); index != -1 {
			if tagValue[:index] == "-" {
				continue
			}

			//If "omitempty" is specified in the tag, it ignores empty values.
			if strings.Index(tagValue[index+1:], "omitempty") != -1 && isEmptyValue(v) {
				continue
			}

			//If "squash" is specified in the tag, we squash the field down.
			squash = !squash && strings.Index(tagValue[index+1:], "omitempty") != -1
			if squash && v.Kind() != reflect.Struct {
				return fmt.Errorf("cannot squash non-struct type '%s'", v.Type())
			}
			keyName = tagValue[:index]
		} else if len(tagValue) > 0 {
			if tagValue == "-" {
				continue
			}
			keyName = tagValue
		}

		switch v.Kind() {
		case reflect.Struct: //this is embeded struct, so handle it differently
			x := reflect.New(v.Type())
			x.Elem().Set(v)

			vType := valMap.Type()
			vKeyType := vType.Key()
			vElemType := vType.Elem()
			mType := reflect.MapOf(vKeyType, vElemType)
			vMap := reflect.MakeMap(mType)

			err := d.decode(keyName, x.Interface(), vMap)
			if err != nil {
				return err
			}

			if squash {
				for _, k := range vMap.MapKeys() {
					valMap.SetMapIndex(k, vMap.MapIndex(k))
				}
			} else {
				valMap.SetMapIndex(reflect.ValueOf(keyName), vMap)
			}
		default:
			valMap.SetMapIndex(reflect.ValueOf(keyName), v)
		}
	}
	if val.CanAddr() {
		val.Set(valMap)
	}
	return nil
}

func (d *Decoder) decodePtr(name string, data interface{}, val reflect.Value) (bool, error) {
	//If the input data is nil, then we want to just set the output
	//pointer to be nil as well.
	isNil := data == nil
	if !isNil {
		switch v := reflect.Indirect(reflect.ValueOf(data)); v.Kind() {
		case reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Ptr,
			reflect.Slice:
			isNil = v.IsNil()
		}
	}
	if isNil {
		if !val.IsNil() && val.CanSet() {
			nilValue := reflect.New(val.Type()).Elem()
			val.Set(nilValue)
		}
		return true, nil
	}

	//Create an element of the concrete (non pointer) type and decode
	//into that. Then set the value of the pointer to this type.
	valType := val.Type()
	valElemType := valType.Elem()
	if val.CanSet() {
		realVal := val
		if realVal.IsNil() || d.config.ZeroFields {
			realVal = reflect.New(valElemType)
		}

		if err := d.decode(name, data, reflect.Indirect(realVal)); err != nil {
			return false, err
		}

		val.Set(realVal)
	} else {
		if err := d.decode(name, data, reflect.Indirect(val)); err != nil {
			return false, err
		}
	}
	return false, nil
}

func (d *Decoder) decodeFunc(name string, data interface{}, val reflect.Value) error {
	//Create an element of the concrete (non pointer) type and decode
	//into that. Then set the value of the pointer to this type.
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	if val.Type() != dataVal.Type() {
		return fmt.Errorf("'%s' expected type '%s', got unconvertable type '%s'", name, val.Type(), dataVal.Type())
	}
	val.Set(dataVal)
	return nil
}

func (d *Decoder) decodeSlice(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	dataValKind := dataVal.Kind()
	valType := val.Type()
	valElemType := valType.Elem()
	sliceType := reflect.SliceOf(valElemType)

	//If we have a non array/slice type then we first attempt to convert.
	if dataValKind != reflect.Array && dataValKind != reflect.Slice {
		if d.config.WeaklyTypedInput {
			switch {
			//Slice and array we use the normal logic
			case dataValKind == reflect.Slice, dataValKind == reflect.Array:
				break
			//Empty maps turn into empty slices
			case dataValKind == reflect.Map:
				if dataVal.Len() == 0 {
					val.Set(reflect.MakeSlice(sliceType, 0, 0))
					return nil
				}
				//Create slice of maps of other sizes
				return d.decodeSlice(name, []interface{}{data}, val)
			case dataValKind == reflect.String && valElemType.Kind() == reflect.Uint8:
				return d.decodeSlice(name, []byte(dataVal.String()), val)

			//All other types we try to convert to the slice type
			//and "lift" it into it. i.e. a string becomes a string slice.
			default:
				//Just re-try this function with data as a slice.
				return d.decodeSlice(name, []interface{}{data}, val)
			}
		}
		return fmt.Errorf("'%s': source data must be an array or slice, got %s", name, dataValKind)
	}

	//If the input vale is nil, then don't allocate since empty != nil
	if dataVal.IsNil() {
		return nil
	}

	valSlice := val
	if valSlice.IsNil() || d.config.ZeroFields {
		//Make a new slice to hold our result, same size as the original data.
		valSlice = reflect.MakeSlice(sliceType, dataVal.Len(), dataVal.Len())
	}

	//Accumulate any errors
	errors := make([]string, 0)

	for i := 0; i < dataVal.Len(); i++ {
		curData := dataVal.Index(i).Interface()
		for valSlice.Len() <= i {
			valSlice = reflect.Append(valSlice, reflect.Zero(valElemType))
		}
		curField := valSlice.Index(i)

		fieldName := fmt.Sprintf("%s[%d]", name, i)
		if err := d.decode(fieldName, curData, curField); err != nil {
			errors = appendErrors(errors, err)
		}
	}

	//Finally, set the value to the slice we built up
	val.Set(valSlice)

	//If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}
	return nil
}

func (d *Decoder) decodeArray(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))
	dataValKind := dataVal.Kind()
	valType := val.Type()
	valElemType := valType.Elem()
	arrayType := reflect.ArrayOf(valType.Len(), valElemType)

	valArray := val

	if valArray.Interface() == reflect.Zero(valArray.Type()).Interface() || d.config.ZeroFields {
		//Check input type
		if dataValKind != reflect.Array && dataValKind != reflect.Slice {
			if d.config.WeaklyTypedInput {
				switch {
				//Empty maps turn into empty arrays
				case dataValKind == reflect.Map:
					if dataVal.Len() == 0 {
						val.Set(reflect.Zero(arrayType))
						return nil
					}
				//All other types we try to convert to the array type
				//and "lift" it into it. i.e. a string becomes a string array.
				default:
					//Just re-try this function with data as a slice.
					return d.decodeArray(name, []interface{}{data}, val)
				}
			}

			return fmt.Errorf("'%s': source data must be an array or slice, got %s", name, dataValKind)
		}
		if dataVal.Len() > arrayType.Len() {
			return fmt.Errorf("'%s': expected source data to have length less or equal to %d, got %d", name, arrayType.Len(), dataVal.Len())
		}

		//Make a new array to hold our result, same size as the original data
		valArray = reflect.New(arrayType).Elem()
	}
	//Accumulate any erros
	errors := make([]string, 0)

	for i := 0; i < dataVal.Len(); i++ {
		curData := dataVal.Index(i).Interface()
		curField := valArray.Index(i)

		fieldName := fmt.Sprintf("%s[%d]", name, i)
		if err := d.decode(fieldName, curData, curField); err != nil {
			errors = appendErrors(errors, err)
		}
	}

	//Finally, set the value to the array we built up
	val.Set(valArray)

	//If there were errors, we return those
	if len(errors) > 0 {
		return &Error{errors}
	}
	return nil
}

func (d *Decoder) decodeStruct(name string, data interface{}, val reflect.Value) error {
	dataVal := reflect.Indirect(reflect.ValueOf(data))

	//If the type of the value to write to and the data match directly,
	//then we just set it directly instread of recursing into the struct.
	if dataVal.Type() == val.Type() {
		val.Set(dataVal)
		return nil
	}

	dataValKind := dataVal.Kind()
	switch dataValKind {
	case reflect.Map:
		return d.decodeStructFromMap(name, dataVal, val)
	case reflect.Struct:
		//Not the most efficient way to do this but we can optimize later if
		//we want to. To convert from struct to struct we go to map first
		//as an intermediary.
		m := make(map[string]interface{})
		mval := reflect.Indirect(reflect.ValueOf(&m))
		if err := d.decodeMapFromStruct(name, dataVal, mval, mval); err != nil {
			return err
		}
		result := d.decodeStructFromMap(name, mval, val)
		return result
	default:
		return fmt.Errorf("'%s' expected a map , got '%s'", name, dataVal.Kind())
	}
}

func (d *Decoder) decodeStructFromMap(name string, dataVal, val reflect.Value) error {
	dataValType := dataVal.Type()
	if kind := dataValType.Key().Kind(); kind != reflect.String && kind != reflect.Interface {
		return fmt.Errorf("'%s' needs a map with string keys, has '%s' keys", name, dataValType.Key().Kind())
	}
	dataValKeys := make(map[reflect.Value]struct{})
	dataValKeysUnused := make(map[interface{}]struct{})
	for _, dataValKey := range dataVal.MapKeys() {
		dataValKeys[dataValKey] = struct{}{}
		dataValKeysUnused[dataValKey.Interface()] = struct{}{}
	}

	errors := make([]string, 0)

	//This slice will keep track of all the struts we'll be decoding.
	//There can be more than one struct if there are embedded structs
	//that are squashed.
	structs := make([]reflect.Value, 1, 5)
	structs[0] = val

	//Compile the list of all the fields that we'er going to be decoding
	//from all the structs.
	type field struct {
		field reflect.StructField
		val   reflect.Value
	}

	//remainField is set to a valid field set with the "remain" tag is
	//we are keeping track of remaining values.
	var remainField *field

	fields := []field{}
	for len(structs) > 0 {
		structVal := structs[0]
		structs = structs[1:]

		structType := structVal.Type()

		for i := 0; i < structType.NumField(); i++ {
			fieldType := structType.Field(i)
			fieldKind := fieldType.Type.Kind()

			//If "squash" is specified in the tag, we squash the field down.
			squash := d.config.Squash && fieldKind == reflect.Struct && fieldType.Anonymous
			remain := false

			//We always parse the tags cause we're looking for other tags too
			tagParts := strings.Split(fieldType.Tag.Get(d.config.TagName), ",")
			for _, tag := range tagParts[1:] {
				if tag == "squash" {
					squash = true
					break
				}
				if tag == "remain" {
					remain = true
					break
				}
			}

			if squash {
				if fieldKind != reflect.Struct {
					errors = appendErrors(errors, fmt.Errorf("%s: unsupported type for squash: %s", fieldType.Name, fieldKind))
				} else {
					structs = append(structs, structVal.FieldByName(fieldType.Name))
				}
				continue
			}

			//Build our field
			if remain {
				remainField = &field{fieldType, structVal.Field(i)}
			} else {
				//Normal struct field, store it away
				fields = append(fields, field{fieldType, structVal.Field(i)})
			}
		}
	}

	//for fieldType, field := range fields {
	for _, f := range fields {
		field, fieldValue := f.field, f.val
		fieldName := field.Name

		tagValue := field.Tag.Get(d.config.TagName)
		tagValue = strings.SplitN(tagValue, ",", 2)[0]
		if tagValue != "" {
			fieldName = tagValue
		}

		rawMapKey := reflect.ValueOf(fieldName)
		rawMapVal := dataVal.MapIndex(rawMapKey)
		if !rawMapVal.IsValid() {
			//Do a slower search by iterating over each key and
			//doing case-insensitive search.
			for dataValKey := range dataValKeys {
				mK, ok := dataValKey.Interface().(string)
				if !ok {
					//Not a string key
					continue
				}
				if strings.EqualFold(mK, fieldName) {
					rawMapKey = dataValKey
					rawMapVal = dataVal.MapIndex(dataValKey)
					break
				}
			}
			if !rawMapVal.IsValid() {
				//There was no matching key in the map for the value
				//in the struct. Just ignore
				continue
			}
		}

		if !fieldValue.IsValid() {
			//This should never happen
			panic("field is not valid")
		}

		//If we can't set the field, then it is unexported or something,
		//and we just continue onwards.
		if !fieldValue.CanSet() {
			continue
		}

		//Delete the key we're using from the unused map so we stop tracking
		delete(dataValKeysUnused, rawMapKey.Interface())

		//If the name is empty string, then we're at the root, and we
		//don't dot-join the fields.
		if len(name) > 0 {
			fieldName = fmt.Sprintf("%s.%s", name, fieldName)
		}

		if err := d.decode(fieldName, rawMapVal.Interface(), fieldValue); err != nil {
			errors = appendErrors(errors, err)
		}
	}

	//If we have a "remain"-tagged field and we have unused keys then
	//we put the unused keys directly into the remain field.
	if remainField != nil && len(dataValKeysUnused) > 0 {
		//Build a map of only the unused values
		remain := map[interface{}]interface{}{}
		for key := range dataValKeysUnused {
			remain[key] = dataVal.MapIndex(reflect.ValueOf(key)).Interface()
		}

		//Set the map to nil so we have none so that the next check will
		//not error (ErrorUnused)
		dataValKeysUnused = nil
	}

	if d.config.ErrorUnuserd && len(dataValKeysUnused) > 0 {
		keys := make([]string, 0, len(dataValKeysUnused))
		for rawKey := range dataValKeysUnused {
			keys = append(keys, rawKey.(string))
		}
		sort.Strings(keys)

		err := fmt.Errorf("'%s' has invalid keys: %s", name, strings.Join(keys, ", "))
		errors = appendErrors(errors, err)
	}

	if len(errors) > 0 {
		return &Error{errors}
	}

	//Add the unused keys to the list of unused keys if we're tracking metadata
	if d.config.Metadata != nil {
		for rawKey := range dataValKeysUnused {
			key := rawKey.(string)
			if len(name) > 0 {
				key = fmt.Sprintf("%s.%s", name, key)
			}
			d.config.Metadata.Unused = append(d.config.Metadata.Unused, key)
		}
	}

	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch getKind(v) {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}
