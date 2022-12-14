package main

import (
	"fmt"
)

func main() {
	// var arr = [5]int{4, 2, 3, 5, 1}
	// var arr = [...]int{4, 2, 3, 5, 1}
	arr := [...]int{4, 2, 3, 5, 1, 0, 8, 9, 7}
	var temp int
	fmt.Println("[+] Original array =", arr)
	
	// bubble sort
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp = arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			}
		}
	}
	// print the sorted array
	fmt.Println("[+] Sorted array\t=", arr)

}
