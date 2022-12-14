package main

import "fmt"

func calculate(num1, num2 int) (int, int, int, int) {
	add := num1 + num2
	sub := num1 - num2
	product := num1 * num2
	divide := num1 / num2

	return add, sub, product, divide
}

func main() {
	a, b, c, d := calculate(12, 4)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
