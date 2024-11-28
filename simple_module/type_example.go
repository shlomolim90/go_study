package simple_module

import (
    "fmt"
    "unsafe"
    "strconv"
)

// 본 파일에서는 Type Casting, Type Assertion 에 대해서 배운다.

func TestUintptr() {
    var x   int     = 10
    // Called Pointer
    var p   *int    = &x
    // Called uintptr
    var u   uintptr = uintptr(unsafe.Pointer(p))

    // uintptr type
    // uintptr is large enough to hold the bit pattern of any pointer without loss of information.
    // uintptr does not reference the value at the address it represents
    //
    // uintptr 은 메모리주소를 단순 숫자형식으로 표현한 것.
    // uintptr 은
    //   1. Type Safety: 어떠한 타입의 포인터 변수를 가질 수 있음. Go 는 꽤나 type 에 제한적이기 때문.
    //   2. Intented Use: 이 주소는 특히 메모리 변수에 대한 조작을 위해 사용됨.
    //   3. GC: uintptr 로 선언된 메모리는 더이상 GC 에서 Tracking 하지 않음으로 메모리 해제가 자동적으로
    //          되지 않음. uintptr can lead to memory management issues if not handled carefully.

    // 추가:
    // uintptr 은 size 계산시에도 주로 사용. 이유: It can hold Any kind of object size if treated as unsigned.
    // size 는 항상 양수.(unsigned)
    fmt.Printf("uintptr u: %x\n", u)
}

type test interface {
    GetName()   string
}

type Test struct {
    Name    string
    Age     int
}

func (t *Test) GetName() string {
    return t.Name
}

func TestTypeCasting() {
    var i   test
    var i2  interface{}
    var t   Test
    t.Name = "LIM"
    t.Age = 3

    i2 = &t
    i = &t

    fmt.Printf("%s\n", i.GetName())
    fmt.Printf("%s\n", i2.(test).GetName())
}

func TestTypeCasting2() {
    var i   Test
    i.Name = "LEE"
    i.Age = 10

    var i2  interface{}

    // Test 구조체를 가리키는 포인터변수를 interface{} 타입으로 캐스팅 = <type>()
    // 참고: &i 을 interface 로 캐스팅하는 이유는 아래 GetName 함수의 receiver 가 pointer 형식이기 때문.
    // 참고: Go 용어로 Casting 이라고 함.
    i2 = interface{}(&i)

    // interface{} 타입을 test interface 로 캐스팅 = <from_interface>.(<to_interface>)
    // 참고: Go 용어로 "Type Assertion" 이라고 함.
    fmt.Printf("%s\n", i2.(test).GetName())
}

func TestStrConversion() {
    // strconv 는 문자열 전환에 특화된 패키지.
    i, err := strconv.Atoi("100")
    fmt.Printf("i:%d, err:%v\n", i, err)

    i, err = strconv.Atoi("s100")
    fmt.Printf("i:%d, err:%v\n", i, err)

    str := strconv.Itoa(231)
    fmt.Printf("str:%s\n", str)
}

func TestType() {
    TestUintptr()
    TestTypeCasting()
    TestTypeCasting2()
    TestStrConversion()
}
