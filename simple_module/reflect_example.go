/*
  go reflect

  Package reflect allowes a program to manipulate objects with arbitrary types.
  Typical use is to take a value with static type interface{} and extract its dynamic type information by calling
  TypeOf
*/


package simple_module

import (
    "fmt"
    "reflect"
)


func TestValueOf() {
    dst := []int{
        1,
        2,
        3,
    }

    fmt.Printf("%p %v\n", &dst, dst)
    // Allocate New object from dst.
    dst2 := reflect.ValueOf(dst)
    fmt.Printf("%p %v\n", &dst2, dst2)
}

func TestTypeOf() {
    /*
    아래는 Type 에 대한 Interface.
    TypeOf 의 결과값은 Type Interface 이며, 아래 함수를 사용할 수 있다.

    type Type interface {
	    // Align returns the alignment in bytes of a value of
	    // this type when allocated in memory.
	    Align() int

	    // FieldAlign returns the alignment in bytes of a value of
	    // this type when used as a field in a struct.
	    FieldAlign() int

	    // Method returns the i'th method in the type's method set.
	    // It panics if i is not in the range [0, NumMethod()).
	    //
	    // For a non-interface type T or *T, the returned Method's Type and Func
	    // fields describe a function whose first argument is the receiver,
	    // and only exported methods are accessible.
	    //
	    // For an interface type, the returned Method's Type field gives the
	    // method signature, without a receiver, and the Func field is nil.
	    //
	    // Methods are sorted in lexicographic order.
	    Method(int) Method

	    // MethodByName returns the method with that name in the type's
	    // method set and a boolean indicating if the method was found.
	    //
	    // For a non-interface type T or *T, the returned Method's Type and Func
	    // fields describe a function whose first argument is the receiver.
	    //
	    // For an interface type, the returned Method's Type field gives the
	    // method signature, without a receiver, and the Func field is nil.
	    MethodByName(string) (Method, bool)

	    // NumMethod returns the number of methods accessible using Method.
	    //
	    // For a non-interface type, it returns the number of exported methods.
	    //
	    // For an interface type, it returns the number of exported and unexported methods.
	    NumMethod() int

	    // Name returns the type's name within its package for a defined type.
	    // For other (non-defined) types it returns the empty string.
	    Name() string

	    // PkgPath returns a defined type's package path, that is, the import path
	    // that uniquely identifies the package, such as "encoding/base64".
	    // If the type was predeclared (string, error) or not defined (*T, struct{},
	    // []int, or A where A is an alias for a non-defined type), the package path
	    // will be the empty string.
	    PkgPath() string

	    // Size returns the number of bytes needed to store
	    // a value of the given type; it is analogous to unsafe.Sizeof.
	    Size() uintptr

	    // String returns a string representation of the type.
	    // The string representation may use shortened package names
	    // (e.g., base64 instead of "encoding/base64") and is not
	    // guaranteed to be unique among types. To test for type identity,
	    // compare the Types directly.
	    String() string

	    // Kind returns the specific kind of this type.
	    Kind() Kind

	    // Implements reports whether the type implements the interface type u.
	    Implements(u Type) bool

	    // AssignableTo reports whether a value of the type is assignable to type u.
	    AssignableTo(u Type) bool

	    // ConvertibleTo reports whether a value of the type is convertible to type u.
	    // Even if ConvertibleTo returns true, the conversion may still panic.
	    // For example, a slice of type []T is convertible to *[N]T,
	    // but the conversion will panic if its length is less than N.
	    ConvertibleTo(u Type) bool

	    // Comparable reports whether values of this type are comparable.
	    // Even if Comparable returns true, the comparison may still panic.
	    // For example, values of interface type are comparable,
	    // but the comparison will panic if their dynamic type is not comparable.
	    Comparable() bool

	    // Bits returns the size of the type in bits.
	    // It panics if the type's Kind is not one of the
	    // sized or unsized Int, Uint, Float, or Complex kinds.
	    Bits() int

	    // ChanDir returns a channel type's direction.
	    // It panics if the type's Kind is not Chan.
	    ChanDir() ChanDir

	    // IsVariadic reports whether a function type's final input parameter
	    // is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
	    // implicit actual type []T.
	    //
	    // For concreteness, if t represents func(x int, y ... float64), then
	    //
	    //	t.NumIn() == 2
	    //	t.In(0) is the reflect.Type for "int"
	    //	t.In(1) is the reflect.Type for "[]float64"
	    //	t.IsVariadic() == true
	    //
	    // IsVariadic panics if the type's Kind is not Func.
	    IsVariadic() bool

	    // Elem returns a type's element type.
	    // It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
	    Elem() Type

	    // Field returns a struct type's i'th field.
	    // It panics if the type's Kind is not Struct.
	    // It panics if i is not in the range [0, NumField()).
	    Field(i int) StructField

	    // FieldByIndex returns the nested field corresponding
	    // to the index sequence. It is equivalent to calling Field
	    // successively for each index i.
	    // It panics if the type's Kind is not Struct.
	    FieldByIndex(index []int) StructField

	    // FieldByName returns the struct field with the given name
	    // and a boolean indicating if the field was found.
	    // If the returned field is promoted from an embedded struct,
	    // then Offset in the returned StructField is the offset in
	    // the embedded struct.
	    FieldByName(name string) (StructField, bool)

	    // FieldByNameFunc returns the struct field with a name
	    // that satisfies the match function and a boolean indicating if
	    // the field was found.
	    //
	    // FieldByNameFunc considers the fields in the struct itself
	    // and then the fields in any embedded structs, in breadth first order,
	    // stopping at the shallowest nesting depth containing one or more
	    // fields satisfying the match function. If multiple fields at that depth
	    // satisfy the match function, they cancel each other
	    // and FieldByNameFunc returns no match.
	    // This behavior mirrors Go's handling of name lookup in
	    // structs containing embedded fields.
	    //
	    // If the returned field is promoted from an embedded struct,
	    // then Offset in the returned StructField is the offset in
	    // the embedded struct.
	    FieldByNameFunc(match func(string) bool) (StructField, bool)

	    // In returns the type of a function type's i'th input parameter.
	    // It panics if the type's Kind is not Func.
	    // It panics if i is not in the range [0, NumIn()).
	    In(i int) Type

	    // Key returns a map type's key type.
	    // It panics if the type's Kind is not Map.
	    Key() Type

	    // Len returns an array type's length.
	    // It panics if the type's Kind is not Array.
	    Len() int

	    // NumField returns a struct type's field count.
	    // It panics if the type's Kind is not Struct.
	    NumField() int

	    // NumIn returns a function type's input parameter count.
	    // It panics if the type's Kind is not Func.
	    NumIn() int

	    // NumOut returns a function type's output parameter count.
	    // It panics if the type's Kind is not Func.
	    NumOut() int

	    // Out returns the type of a function type's i'th output parameter.
	    // It panics if the type's Kind is not Func.
	    // It panics if i is not in the range [0, NumOut()).
	    Out(i int) Type

	    // OverflowComplex reports whether the complex128 x cannot be represented by type t.
	    // It panics if t's Kind is not Complex64 or Complex128.
	    OverflowComplex(x complex128) bool

	    // OverflowFloat reports whether the float64 x cannot be represented by type t.
	    // It panics if t's Kind is not Float32 or Float64.
	    OverflowFloat(x float64) bool

	    // OverflowInt reports whether the int64 x cannot be represented by type t.
	    // It panics if t's Kind is not Int, Int8, Int16, Int32, or Int64.
	    OverflowInt(x int64) bool

	    // OverflowUint reports whether the uint64 x cannot be represented by type t.
	    // It panics if t's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.
	    OverflowUint(x uint64) bool

	    // CanSeq reports whether a [Value] with this type can be iterated over using [Value.Seq].
	    CanSeq() bool

	    // CanSeq2 reports whether a [Value] with this type can be iterated over using [Value.Seq2].
	    CanSeq2() bool
	    // contains filtered or unexported methods
        }
    */

    t := reflect.TypeOf(int(40))

    // = 8bytes (=64bits)
    fmt.Printf("Align(): %d\n", t.Align())
    // = 0
    fmt.Printf("NumMethod(): %d\n", t.NumMethod())
    // = int
    fmt.Printf("Name(): %s\n", t.Name())
    // = ""
    fmt.Printf("PkgPath(): %s\n", t.PkgPath())
    // Its return type is uintptr.
    // = 8bytes (=64bits)
    fmt.Printf("Size(): %x\n", t.Size())
}


func TestDeepEqual() {
    // Two values of identical type are deeply equal if one of following cases applies.
    // 1. Array values are deeply equal when their corresponding elements are deeply equal.
    // 2. Struct values are deeply equal if their corresponding fields, both exported and unexported? are deeply equal.
    // 3. Func values are deeply euqal if both are nil;
    // 4. Interface values are deeply equal if they hold deeply equal concrete values.
    // 5. Map values are deeply equal when both nil, both non-nil, they have same length and either they are the
    //    same map object or their corresponding keys map to equal values 
    // 6. Pointer values are deeply equal if they are equal using Go's == operator or if they point to deeply equal values.
    // <Need to write more detail.>

    // Creating a structure
    type Author struct {
        name      string
        branch    string
        language  string
        Particles int
    }
    // Creating variables
    // of Author structure
    a1 := Author{
        name:      "Moana",
        branch:    "CSE",
        language:  "Python",
        Particles: 38,
    }
 
    a2 := Author{
        name:      "Moana",
        branch:    "CSE",
        language:  "Python",
        Particles: 38,
    }
 
    // = true
    fmt.Printf("%t\n", reflect.DeepEqual(&a1, &a2))
}


func TestReflect() {
    TestValueOf()
    TestTypeOf()
    TestDeepEqual()
}
