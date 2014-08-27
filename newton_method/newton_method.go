package main

import (
    "fmt"
)

func Sqrt(x float64, t int) float64 {
    z := 1.0
    for i:=0;i<t;i++ {
    	z = z - (z*z-x)/2
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2, 10))
    fmt.Println(Sqrt(2, 100))
    fmt.Println(Sqrt(2, 1000))
    fmt.Println(Sqrt(3, 10))
    fmt.Println(Sqrt(3, 100))
    fmt.Println(Sqrt(3, 1000))
}