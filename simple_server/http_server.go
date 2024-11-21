package simple_server

/*

import (
    "fmt"
    "net/http"
)


func handleConnection(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 1024)
    rlen, err := conn.Read(buf)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Received len:%d buf:%s", rlen, buf)
}


func HttpStart(port string) {
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))



    ln, err := net.Listen("tcp", port)
    // no set is nil
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }

        go handleConnection(conn)
    }
}

// Entry Point.
func TestHttpServer() {
    HttpStart(":8081")
}
*/
