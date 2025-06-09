package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	i := 20
	var f float64 = float64(i)
	fmt.Println("Exercise 1:")
	fmt.Println(i)
	fmt.Println(f)
}

func ex2() {
	const value = 10
	var i int = value
	var f float64 = value
	fmt.Println("Exercise 2:")
	fmt.Println(i)
	fmt.Println(f)
}

func ex3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI int64 = 9223372036854775807
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println("Exercise 3:")
	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)
}
