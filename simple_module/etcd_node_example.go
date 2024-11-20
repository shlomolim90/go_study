package simple_module


import (
    "fmt"
    "time"
)


type RawData struct {
    id   int
    name string
    age  int
}

type node struct {
    stop chan struct{}
    done chan struct{}
    data *RawData
}

// go routine
func (n *node) run() {
    fmt.Printf("Node %s (%d) is running\n", n.data.name, n.data.age)

    loop:
    for {
        select {
        case <-n.stop:
            fmt.Printf("Node %s (%d) is stopping\n", n.data.name, n.data.age)
            time.Sleep(2 * time.Second)
            defer close(n.done)
            break loop
        }
    }
}

type Node interface {
    GetName() string
    Stop()
    WaitDone()
}

func (n *node) GetName() string {
    return n.data.name
}

func (n *node) Stop() {
    select {
    case n.stop <- struct{}{}:
    }
}

func (n *node) WaitDone() {
    select {
    case <- n.done:
    }
}

// Accessible Functions.
func TestNode() {
    fmt.Printf("----------- TestNode() start ----------\n")
    n := StartNode("Lim", 35)

    time.Sleep(4 * time.Second)
    fmt.Printf("Main Sleep done\n")
    n.Stop()
    fmt.Printf("Main Send stop\n")
    n.WaitDone()
    fmt.Printf("Main WaitDone Node\n")
    fmt.Printf("----------- TestNode() done ----------\n")
}

func StartNode(name string, age int) Node {
    var rd RawData
    rd.name = name
    rd.age = age

    n := NewNode(&rd)
    go n.run()

    // Because "func (n *node) Stop()"
    return &n
}

func NewNode(r *RawData) node {
    return node{
        // Without initialization, Default "nil".
        stop: make(chan struct{}),
        done: make(chan struct{}),
        data: r,
    }
}
