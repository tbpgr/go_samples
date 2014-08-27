package main

import "fmt"

type Person struct {
    Name string
    Age int
}


func main() {
    tanaka := Person{"tanaka", 34}
    fmt.Println(tanaka)
    fmt.Println(tanaka.Name)
    fmt.Println(tanaka.Age)
}