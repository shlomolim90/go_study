
package simple_module

import (
    "fmt"
    "unsafe"
)



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


func TestType() {
    TestUintptr()
}
