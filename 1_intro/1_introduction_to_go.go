package main

import (
	"fmt"
)

func main() {
	fmt.Println("Q: can i")

	var intVar int = 42
	fmt.Println(intVar)

	var strVar string = "Go is fun!"
	fmt.Println(strVar)

	var boolVar bool = true
	fmt.Println(boolVar)

	var byteVar byte = 'A'
	fmt.Println(byteVar)

	var runeVar rune = '世'
	fmt.Println(runeVar)

	var complexVar complex128 = 1 + 2i
	fmt.Println(complexVar)

	var uintVar uint = 100
	fmt.Println(uintVar)

	var int8Var int8 = -128
	fmt.Println(int8Var)

	var int16Var int16 = 32767
	fmt.Println(int16Var)

	var int32Var int32 = 2147483647
	fmt.Println(int32Var)

	var int64Var int64 = 9223372036854775807
	fmt.Println(int64Var)

	var uint8Var uint8 = 255
	fmt.Println(uint8Var)

	var uint16Var uint16 = 65535
	fmt.Println(uint16Var)

	var uint32Var uint32 = 4294967295
	fmt.Println(uint32Var)

	var uint64Var uint64 = 18446744073709551615
	fmt.Println(uint64Var)

	var float32Var float32 = 3.14
	fmt.Println(float32Var)
	
	var float64Var float64 = 3.141592653589793
	fmt.Println(float64Var)

	var ptrVar *int = &intVar
	fmt.Println(ptrVar)
	fmt.Println(*ptrVar)
	
	var arrVar [3]int = [3]int{1, 2, 3}
	fmt.Println(arrVar)

	var sliceVar []int = []int{4, 5, 6}
	fmt.Println(sliceVar)

	var mapVar map[string]int = map[string]int{"one": 1, "two": 2}
	fmt.Println(mapVar)

	var structVar struct {
		name string
		age  int
	} = struct {
		name string
		age  int
	}{"Alice", 30}
	fmt.Println(structVar)

	var chanVar chan int = make(chan int)
	go func() {
		chanVar <- 7
	}()
	fmt.Println(<-chanVar)

	var funcVar func(int) int = func(x int) int { return x * x }
	fmt.Println(funcVar(5))
	
	var interfaceVar interface{} = "I can hold any type"
	fmt.Println(interfaceVar)

	var nilVar interface{} = nil
	fmt.Println(nilVar)

	const constVar int = 10
	fmt.Println(constVar)
	
	const iotaVar1 = iota + 2
	const iotaVar2 = iota + 56
	fmt.Println(iotaVar1, iotaVar2)
}