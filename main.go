package main

import "fmt"

func main() {
	defer func() {
		r := recover(); r != nil {
			fmt.Println("Cannot be divided by zero...")

			main()
		}
	}()

	functions := map[string]func(int, int) int {
		"addition": add,
		"subtraction": sub,
		"multiplication": multiply,
		"division": mul,
	}

	var currentNumber int
	fmt.Println("Enter a number: ")
	fmt.Scanln(&currentNumber)

	for true {
		var functionName string
		fmt.Println("What funcion do you want to use: ")
		fmt.Scanln(&functionName)
		if functionName == "done" {
			break
		}
		var number int
		fmt.Println("Enter a number: ")
		fmt.Scanln(&number)

		currentNumber = functions[functionName](currentNumber,number)
		fmt.Println("Your result is: ")
		fmt.Println(currentNumber)
		
	}
}

func addition(x, y int) int {
	return x + y
}

func subtraction(x, y int) int {
	return x - y
}

func multiplication(x, y int) int {
	return x * y
}

func division(x, y int) int {
	return x / y
}
