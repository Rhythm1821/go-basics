package main
import "fmt"

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
	fmt.Println(student1, student2, a, b, true1, false2, a1,b1,c1,x,y,z,a2,b2,c2)
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