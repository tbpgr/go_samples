package main

import (
	"fmt"
)

type Person struct {
	FirstName, LastName string
}

func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	p := &Person{"kazuo", "tanaka"}
	fmt.Println(p.FullName())
}
