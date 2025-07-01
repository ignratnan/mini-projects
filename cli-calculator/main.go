package calculator

import (
	"fmt"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	return a / b
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Main() {
	//Variable description
	var first_num float64
	var second_num float64
	var operator string
	var precision uint = 2

	fmt.Println("Welcome to Go Calculator")

	fmt.Print("Input the first number: ")
	_, err := fmt.Scanln(&first_num)
	if err != nil {
		fmt.Println("Error reading first number: ", err)
		return
	}

	fmt.Print("Input the operator ( + , - , * , / ): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Error reading first number: ", err)
		return
	}

	for operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		fmt.Println("Invalid operator. Please enter one of the following: +, -, *, /")
		fmt.Print("Input the operator ( + , - , * , / ): ")
		_, err = fmt.Scanln(&operator)
		if err != nil {
			fmt.Println("Error reading first number: ", err)
			return
		}
	}

	fmt.Print("Input the second number: ")
	_, err = fmt.Scanln(&second_num)
	if err != nil {
		fmt.Println("Error reading second number: ", err)
		return
	}

	// Switch Case Version
	switch operator {
	case "+":
		result := add(first_num, second_num)
		f_result := roundFloat(result, precision)
		fmt.Println("Result:", f_result)
	case "-":
		result := sub(first_num, second_num)
		f_result := roundFloat(result, precision)
		fmt.Println("Result:", f_result)
	case "*":
		result := mul(first_num, second_num)
		f_result := roundFloat(result, precision)
		fmt.Println("Result:", f_result)
	case "/":
		result := div(first_num, second_num)
		f_result := roundFloat(result, precision)
		fmt.Println("Result:", f_result)
	default:
		fmt.Println("Invalid operator")
	}

}
