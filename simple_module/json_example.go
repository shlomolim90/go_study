package simple_module


import (
    "fmt"

    "encoding/json"
)


// interface{} on the position of value indicates that zero methods.
// An empty interface may hold values of any type. (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type.
type Dictionary map[string]interface{}


func TestJson() {
    fmt.Printf("----------- TestJson() start ----------\n")
    data := []Dictionary {
        {
            "metrics": []Dictionary {
                {"hello": "good"},
                {"bad": "is yours"},
            },
        },
    }

    var list =  []string {
        "apple",
        "peach",
        "pear",
    }

    ret1, _ := json.Marshal(list)
    fmt.Println(string(ret1))
    ret2, _ := json.Marshal(data)
    fmt.Println(string(ret2))
    fmt.Printf("----------- TestJson() Done ----------\n")
}
