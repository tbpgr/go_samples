package main

import (
	"fmt"
)

type Aisatsuer interface {
	Hello() string
}

func main() {
	aisatuers := []Aisatsuer{Human{},Dog{}}

	for _, a := range aisatuers {
		fmt.Println(a.Hello())
	}
}

type Human struct {
}

func (h Human) Hello() string {
	return "hello"
}

type Dog struct {
}

func (d Dog) Hello() string {
	return "bow-wow "
}

