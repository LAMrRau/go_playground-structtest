package main

import (
	"fmt"
	"unsafe"
)

type BadStruct struct {
	age         uint8
	passportNum uint64
	siblings    uint16
}

type GoodStruct struct {
	age         uint8
	siblings    uint16
	passportNum uint64
}

func main() {
	fmt.Println("Hello World")

	var i int
	var u uint
	var up uintptr

	fmt.Printf("i Type:%T Size:%d\n", i, unsafe.Sizeof(i))
	fmt.Printf("u Type:%T Size:%d\n", u, unsafe.Sizeof(u))
	fmt.Printf("up Type:%T Size:%d\n", up, unsafe.Sizeof(up))

	var badStruct = BadStruct{}
	var goodStruct = GoodStruct{}

	fmt.Printf("good struct Type: %T Size:%d\n", goodStruct, unsafe.Sizeof(goodStruct))
	fmt.Printf("bad struct Type: %T Size:%d\n", badStruct, unsafe.Sizeof(badStruct))
}
