package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	calculate := bufio.NewReader(os.Stdin)

	fmt.Println("Enter first number: ")
	firstNumber, _ := calculate.ReadString('\n')
	aFloat, erro:= strconv.ParseFloat(strings.TrimSpace(firstNumber), 64)
	if erro != nil {
		fmt.Println("Error: ", erro)
		panic("Value first number must be a number")
	} 
	fmt.Println("Enter second number: ")
	secondNumber, _ := calculate.ReadString('\n')
	bFloat, erro:= strconv.ParseFloat(strings.TrimSpace(secondNumber), 64)
	if erro != nil {
		fmt.Println("Error: ", erro)
		panic("Value second number must be a number")
	} 

	sum := aFloat + bFloat;
	sum = float64(int(sum * 100)) / 100
	fmt.Println("The sum is: ", sum)
}