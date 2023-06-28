package main

import "fmt"

var aMila = 100

func main() {
	var aString string = "Hello, Go"
	fmt.Println(aString)
	fmt.Printf("The variable aString is of type %T\n", aString)

	var anInt int = 42
	fmt.Println(anInt)
	fmt.Printf("The variable anInt is of type %T\n", anInt)

	var aFloat float32 = 243.146
	fmt.Println(aFloat)
	fmt.Printf("The variable aFloat is of type %T\n", aFloat)
	
	var aBool bool = true
	fmt.Println(aBool)
	fmt.Printf("The variable aBool is of type %T  \n", aBool) 
	
	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("The variable defaultInt is of type %T\n", defaultInt)

	fmt.Println(aMila)
	fmt.Printf("The variable aMila is of type %T\n", aMila)
}