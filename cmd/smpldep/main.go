package main

import (
    "fmt"
)

func main() {
    fmt.Println("hello from sample dependency")
}

func getSampleDependency() {
    fmt.Println("hello from getSampleDependency() method")
}
