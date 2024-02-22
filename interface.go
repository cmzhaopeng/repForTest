package main

import (
	"fmt"
)

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("No name")
		return
	}
	fmt.Println(d.name, "barks")
}

// func (d Dog) Speak() {
//	fmt.Println(d.name,"barks")
//}

func main() {
	var d Dog = Dog{name: "Spot"}
	//var d *Dog
	var s Speaker
	//s = d
	s = &d
	s.Speak()
}
