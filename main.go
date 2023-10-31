package main

import (
    "fmt"
    "singleton"
)

func main() {
    instance1 := singleton.getInstance()
    instance2 := singleton.getInstance()

    fmt.Println(instance1 == instance2)
}
