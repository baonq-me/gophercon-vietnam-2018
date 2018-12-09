package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var isGolangFun bool = true
	//isGolangFun := true
	fmt.Println(isGolangFun)

	var myInteger int = 10
	fmt.Println(&myInteger)

	//myString := "We are gophers"
	fmt.Println(unsafe.Sizeof(complex128(2+3i)))
	fmt.Println(unsafe.Sizeof(isGolangFun))
	fmt.Println(unsafe.Sizeof(myInteger))

	/*
	 * Basic golang types:
	 * 		bool
	 * 		string
	 * 		int  int8  int16  int32  int64
	 * 		uint uint8 uint16 uint32 uint64 uintptr
	 * 		byte
	 * 		rune
	 * 		float32 float64
	 * 		complex64 complex128
	 */
}