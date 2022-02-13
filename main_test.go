package main

import (
	"testing"
)

func BenchmarkCreateGoodStructArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mockGoodStructArray()
	}
}

func BenchmarkCreateBadStructArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = mockBadStructArray()
	}
}

func BenchmarkTraverseGoodStruct(b *testing.B) {
	testArray := mockGoodStructArray()
	for n := 0; n < b.N; n++ {
		traverseGoodStruct(testArray)
	}
}

func BenchmarkTraverseBadStruct(b *testing.B) {
	testArray := mockBadStructArray()
	for n := 0; n < b.N; n++ {
		traverseBadStruct(testArray)
	}
}

func traverseGoodStruct(testArray []GoodStruct) uint16 {
	var arbitraryNum uint16
  
	for _, goodStruct := range testArray {
		arbitraryNum += goodStruct.siblings
	}
  
	return arbitraryNum
}

func traverseBadStruct(testArray []BadStruct) uint16 {
	var arbitraryNum uint16
  
	for _, badStruct := range testArray {
		arbitraryNum += badStruct.siblings
	}
  
	return arbitraryNum
}

func mockGoodStructArray() []GoodStruct {
	goodStructArray := []GoodStruct{}

	var index uint16
	for index = 0;  index < 10000; index++ {
		goodStructArray = append(goodStructArray, GoodStruct{
			age: 5,
			passportNum: 10,
			siblings: index,
		})
	}

	return goodStructArray
}

func mockBadStructArray() []BadStruct {
	goodStructArray := []BadStruct{}

	var index uint16
	for index = 0;  index < 10000; index++ {
		goodStructArray = append(goodStructArray, BadStruct{
			age: 5,
			passportNum: 10,
			siblings: index,
		})
	}

	return goodStructArray
}
