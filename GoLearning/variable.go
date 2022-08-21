package main

import ("fmt")

//varibale := ":= variabel work inside function"

func main() {

	variable := " := this variable work inside functioni"
	fmt.Println(variable)

	var student1 string = "john" // type is string 
	var student2 = "rahim" // type is inferred
	x := 2 // type is inferred
	//	x := // will not work := will not work without assing variable. 

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

}
