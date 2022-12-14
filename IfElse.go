package main

import "fmt"

func check_greatest(num1, num2 int) {
	if num1 > num2 {
		fmt.Printf("%d is greater", num1)
	} else if num2 > num1 {
		fmt.Printf("%d is greater", num2)
	} else {
		fmt.Println("Draw")
	}
}

func main() {
	fmt.Println("Working well")
	check_greatest(1, 1)
}
