package simple_module

import (
    "fmt"
)

func Echo(config string) int {
    fmt.Println("Load config done\n")
    fmt.Println("Initialize Server\n")
    fmt.Println("Listen IP/Port done\n")
    fmt.Println("Wait events\n")
    fmt.Println("Close Server\n")
    return 0
}
