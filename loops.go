package main

import "fmt"

func main() {

	// 1st method .
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("i = ", i)
	// }

	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//table
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "X", j, "=", i*j)
		}
	}

}
