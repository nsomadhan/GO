package main

import "fmt"

func main() {
	fmt.Println("const variable declares the variable as constant")
	fmt.Println("Which means unchangeable")

	const MyConstName string = "Rahim uddin"
	const PI float64 = 3.1416

	fmt.Println(MyConstName)
	fmt.Println(PI)

	const (

		A int = 1
		B = 2
		C string = "I am String"
	)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)



}