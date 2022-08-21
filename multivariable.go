package main

import "fmt"


func main() {
  
 fmt.Println ("Go variable name are case sensetive");

 var a, b, c int = 1, 2, 3;
 fmt.Printf("T",a);
 fmt.Println(b);
 fmt.Println(c);



 // Another way 

 var x, y, z = 6, "string", 1.5
 m, n := 7, "String"

 fmt.Println(x)
 fmt.Printf("T", x)
 fmt.Println(y)
 fmt.Println(z)
 fmt.Printf("T", z)
 fmt.Println(m)
 fmt.Printf("T", m)
 fmt.Println(n)



//Camel Case Variable
var myVariableName = "My name"
fmt.Println(myVariableName)
fmt.Printf("T", myVariableName)

// Pascal Case Variable
MyVaribaleName := "My name"
fmt.Println(MyVaribaleName)
fmt.Printf("T", MyVaribaleName)

//Snake Case
var my_variable_name string = "My name"
fmt.Println(my_variable_name)
fmt.Printf("T", my_variable_name)

}; 