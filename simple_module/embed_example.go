package simple_module

// Null container
import _ "embed"


//go:embed hello.txt
var s string

func TestEmbed() {
    print(s)
}
