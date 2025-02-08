package main

import (
	"fmt"
)

func main()  {
	// All printing ways

	// Prints with a newline
	fmt.Println("Hello World")  

	// Prints without a newline
	// fmt.Print("Hello World")	

	// Prints with a format
	name:= "John"
	age:= 20
	fmt.Printf("Name: %s, Age: %d\n",name,age) 	

	// Sprintf
	msg:= fmt.Sprintf("Hello, %s","Bob")
	fmt.Println(msg)

	// Go Variable Types
	var student1 string = "John"
	var student2 = "Jane"  //type is inferred
	var a int = 20
	var b = 30
	var true1 bool = true
	var false2 = false

	var a1,b1,c1 int = 1,2,3
	x,y,z:=11,22,33

	var (
		a2 int = 1
		b2 int = 2
		c2 int = 3
	)

	// constants 

	const Pi = 3.14159
	const (
		Pi1 = 3.14159
		Pi2 = 3.14159
	)
	const Pi3 int = 3
	fmt.Println(student1, student2, a, b, true1, false2, a1,b1,c1,x,y,z,a2,b2,c2,Pi,Pi1,Pi2,Pi3)

	// Array Syntax - var <name> = [<length]type{<values>} | <name> := [<length]type{<values>}
	var arr = [4]int{1,2,3,4}
	var arr2 = [...]int{2,3,1,6}

	arr[0] = 10
	arr3 := [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(arr, arr2, arr3)
	fmt.Println(arr[0:2])

	// Array initialization
	arr4 := []int{1,2,3,4} // Initialize with values
	arr5 := [6]int{} // Initialize all elements with 0

	arr6 := [5]int{3: 10,1: 20} // Initialize only specified elements

	fmt.Println(arr4, arr5, arr6)
	fmt.Println(len(arr4), cap(arr5)) // 1. Length, 2. Capacity: length of the underlying array

	// Slices syntax - slice_name := []datatype{values}
	slice1 := []int{1,2,3,4}
	slice2 := []string{"Hello", "World"}
	fmt.Println(slice1, slice2)

	// Append operation
	slice1 = append(slice1, 5)
	fmt.Println(slice1)
}

/*
Go has three basic data types:

bool: represents a boolean value and is either true or false
Numeric: represents integer types, floating point values, and complex types
string: represents a string value

// Go Variable Types

int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
string - stores text, such as "Hello World". String values are surrounded by double quotes
bool- stores values with two states: true or false
*/