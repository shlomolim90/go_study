package main

import (

    "github.com/go_study/simple_module"
    "github.com/go_study/simple_server"
)


func main() {
    //simple_module.Echo("start!")
    //simple_module.TestWaiter()
    simple_module.TestNode()
    simple_module.TestJson()

    simple_server.TestServer()

    //simple_module.TestGoroutine()
    //simple_module.TestGoChannel()
    //simple_module.TestAtMostWork()
    //simple_module.TestCountDown()
    //simple_module.TestWaiter()
}
