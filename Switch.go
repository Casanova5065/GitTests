package main

import "fmt"

func main() {
	fmt.Println("Switch-case demo !!")

	var num1, num2 float64
	var op string

	fmt.Print("Enter first number : ")
	fmt.Scanf("%f", &num1)
	fmt.Print("Enter second number : ")
	fmt.Scanf("%f", &num2)
	fmt.Print("Enter Operation : ")
	fmt.Scanf("%s", &op)
	fmt.Println("Values  =", num1, num2, "\n", "Operation =", op)

	switch op {
	case "+":
		fmt.Printf("%f + %f = %.2f", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%f - %f = %.2f", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%f * %f = %.2f", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%f / %f = %.2f", num1, num2, num1/num2)
	default:
		fmt.Println("[-] Operation undefined !!")
	}

}
