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

	var badStruct = BadStruct{}
	var goodStruct = GoodStruct{}

	fmt.Printf("good struct Type: %T Size:%d\n", goodStruct, unsafe.Sizeof(goodStruct))
	fmt.Printf("bad struct Type: %T Size:%d\n", badStruct, unsafe.Sizeof(badStruct))
}
