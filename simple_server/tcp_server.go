package simple_server


import (
    "fmt"
    "net"
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


func TcpStart(port string) {
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
func TestServer() {
    TcpStart(":8081")
}
