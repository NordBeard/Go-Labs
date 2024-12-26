package main

import "fmt"

func main() {
	//1
	hello()
	//2
	var firstNumber int
	var secondNumber int
	err := printEven(firstNumber, secondNumber)
	if err != nil {
		fmt.Println("error:", err)
	}
	//3
	var a float64
	var b float64
	var op string
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&op)
	ans, err := apply(a, b, op)
	fmt.Println(ans, err)
}

func apply(a, b float64, operator string) (float64, error) {
	if operator == "+" {
		return a + b, nil
	} else if operator == "-" {
		return a - b, nil
	} else if operator == "*" {
		return a * b, nil
	} else if operator == "/" {
		if b == 0 {
			return 0, fmt.Errorf("Divide by zero")
		} else {
		}
		return a / b, nil
	} else {
		return 0, fmt.Errorf("Invalid operator")
	}
}

func printEven(a, b int) error {
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a > b {
		return fmt.Errorf("Error")
	}
	for i := a; i < b+1; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	return nil
}

func hello() {
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello, " + name + "!")
}
