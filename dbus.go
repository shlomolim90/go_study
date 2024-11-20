package main

import (
    "fmt"
    "os"

    "github.com/godbus/dbus/v5"
)



// Notice: To build this source, Pass env variable "CGO_ENABLED=0"
func main() {

    // DBus Configuration
    dest := "-------------------------"
    path := dbus.ObjectPath("/")
    method := "------------------------------------"


    conn, err := dbus.ConnectSystemBus()
    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to connect to session bus:", err)
        os.Exit(1)
    }

    defer conn.Close()

    var s string

    obj := conn.Object(dest, path)
    err = obj.Call(method, 0, "").Store(&s)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Failed to call: ", err)
        os.Exit(1)
    }

    fmt.Println(s)
}
