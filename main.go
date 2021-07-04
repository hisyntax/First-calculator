package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Cannot divide a number by zero: ")
			main()
		}
	}()

	functions := map[string]func(int, int) int{
		"add": add,
		"sub": sub,
		"mul": mul,
		"div": div,
	}
	var currentNumber int
	fmt.Println("Enter a number")
	fmt.Scanln(&currentNumber)

	for true {
		var functionName string
		var number int
		fmt.Println("What function do you want to use e.g mul = multiply, add = addition, div = division, sub = subtract ... and done = done to end the program")
		fmt.Scanln(&functionName)
		if functionName == "done" {
			break
		}
		//This is the feature which was added
		//The user should be able to know what "ARITHEMATIC OPERATOR" he/she is working on and also to see the "PREVIOUSLY IMPUTED NUMBER"
		fmt.Println(functionName, currentNumber, "by...")
		fmt.Scanln(&number)

		currentNumber := functions[functionName](currentNumber, number)
		fmt.Println("Your result is: ")
		fmt.Println(currentNumber)
	}

}
