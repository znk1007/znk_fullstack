//Package map2struct exposes functionality to convert one arbitrary(任意)
//Go type into another, typically to convert a map[string]interface{}
//into a native Go struct
//
//The Go struct can be arbitrarily complex, containing slices, other structs,
//etc. And the decoder will properly decode nested maps and so on into the
//proper structs in the native Go struct.
//
//See the examples to see what the decoder is capable of.
//The simplest function to start with is Decode.
//
//Field Tags
//
//When decoding to a struct, map2struct will use the field name by default
//to perform the mapping. For example, if a struct has a field "Username"
//then map2struct will look for a key in the source value of "username" (case
//insensitive).
//
//	type User struct {
//		Username string
//	}
//
//You can change the behavior of map2struct by using struct tags.
//The default struct tag that map2struct looks for is "map2struct", but you
//can customize it using DecodeConfig.
//
//Renaming Fields
//
//To rename the key that map3struct looks for, use the "map2struct" tag and
//set a value directly. For example, to change the "username" example
//above to "user":
//
//	type Use struct {
//		Username string `map2struct:"user"`
//	}
//
//Embedded Structs and Squashing
//
//Embedded structs are treated as if they're another field with that name.
//By default, the two structs below are equivalent when decoding with map2struct:
//
//	type Person struct {
//		Name string
//	}
//
//	type Friend struct {
//		Person
//	}
//
//	type Friend struct {
//		Person Person
//	}
//
//This would require an input that looks like below:
//
//	map[string]interface{}{
//		"person":map[string]interface{}{"name":"alice"},
//	}
//
//If your "person" value is NOT nested, then you can append ",squash" to
//your tag value and map2struct will treat(认定) it as if the embedded(嵌套)
//struct were part of the struct directly.
//
//Example:
//
// 	type Friend struct {
//		Person `map2struct:",squash"`
// 	}
//
//Now the following input whould be accepted:
//
//	map[string]interface{}{
//		"name":"alice",
//	}
//
//DecodeConfig has a field that changes the behavior of map2struct
//to always squash embedded structs.
//
//Remainder Values
//
//If there are any unmapped keys in the source value, map2struct by default
//will silently ignore them. You can error by setting ErrorUnused in
//DecoderConfig. If you're using Metadata you can also maintain a slice of
//the unused keys.
//
//You can also use the ",remain" suffix on your tag to collect all unused
//values in a map. The field with this tag MUST be a map type and should
//probably be a "map[string]interface{}" or "map[interface{}]interface{}"
//See example below:
//
//	type Friend struct {
//		Name string
//		Other map[string]
//	}
//
//Given the input below, Other would be populated with the other values
//that weren't used (everything but "name"):
//
//	map[string]interface{}{
//		"name":"bob",
//		"address":"123 Maple St.",
//	}
//
//Omit Empty Values
//
//When decoding from a struct to any other value, you may use the ",omitempty"
//suffix on your tag to omit that value it it equates to the zero value.
//The zero value of all types is specified in the Go specification.
//
//For example, the zero type of a numeric type is zero ("0"). If the struct
//field value is zero and a numeric type, the field is empty, and it won't
//be encoded into the destination type.
//
//	type Source {
//		Age int `map2struc:",omitempty"`
//	}
//
//Unexported fields
//
//Since unexported (private) struct fields cannot be set outside the package
//where they are defined, the decoder will simply skip them.
//
//For this output type definition:
//
//	type Exported struct {
//		private string // this unexported field will be skipped
// 		Public	string
//	}
//
//Using this map as input:
//
//	map[string]interface{}{
//		"private": "I will be ignored",
//		"Public": "I made it through!",
//	}
//
//The following struct will be decoded:
//
//	type Exported struct {
//		private :""//field is left with an empty string (zero value)
//		Public: "I made it through!"
//	}
//
//Other Configuration
//
//map2struct is highly configurable. See the DecoderConfig struct for
//other features and options that are supported.
package m2s
