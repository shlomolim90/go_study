package simple_module


import (
    "fmt"

    "encoding/json"
)


// Make a dictionary for json. Key should be string type but, Value can be any type.
// interface{} on the position of value indicates that zero methods.
// An empty interface may hold values of any type. (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type.
type dict map[string]interface{}

// It is an instance to be converted from json bytes.
type JsonStrictStruct struct {
    // The backquote block is struct tag.
    // Using backquote means that the string between backquotes treated as raw string.
    // json tags is used for JSON data transformation.
    // For example, json:"userid" means the value of "userid" key in JSON bytes can be
    // decoded to "UserID" in JsonStrictSturct when using Unmarshal().
    UserID      int                     `json:"userid"`
    ID          int                     `json:"id"`
    Title       string                  `json:"title"`
    Complete    bool                    `json:"complete"`
    // Rest of the fields should go here.
    Rest        map[string]interface{}  `json:"-"`
}

type JsonGeneralStruct struct {
    UserID      interface{}             `json:"userid"`
    ID          interface{}             `json:"id"`
    Title       interface{}             `json:"title"`
    Complete    interface{}             `json:complete"`
    // Rest of the fields should go here.
    Rest        map[string]interface{}  `json:"-"`
}


// Entry Point.
func TestJson() {
    var strict_to JsonStrictStruct
    var general_to JsonGeneralStruct


    fmt.Printf("----------- TestJson() start ----------\n")
    from := dict {
        "userid": 6,
        "id": 6,
        "title": "What a nice day",
        "complete": true,
        "Unknown": "hello",
    }

    from_bytes, _ := json.Marshal(from)

    err := json.Unmarshal(from_bytes, &strict_to)
    if err != nil {
        fmt.Println(err)
    }

    err = json.Unmarshal(from_bytes, &general_to)
    fmt.Printf("UserID:%d, ID:%d, Title:%s, Complete:%t\n",
                strict_to.UserID, strict_to.ID, strict_to.Title, strict_to.Complete)
    fmt.Printf("UserID:%d, ID:%d, Title:%s, Complete:%t\n",
                general_to.UserID, general_to.ID, general_to.Title, general_to.Complete)
    fmt.Printf("----------- TestJson() Done ----------\n")
}
