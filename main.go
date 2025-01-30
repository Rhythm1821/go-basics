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
}