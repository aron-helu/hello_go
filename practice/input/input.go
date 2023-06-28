package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
) 

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your full name: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Enter name: ", input)

	fmt.Println("Enter your age: ")
	age, _:= reader.ReadString('\n')
	aFloat, erro:= strconv.ParseFloat(strings.TrimSpace(age), 64)

	if erro != nil {
		fmt.Println("Error: ", erro)
	} else {
		fmt.Println("Your age is: ", aFloat)
	}

}

