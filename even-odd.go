package main

import "fmt"

func main() {

	new_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even_slice := []int{}
	odd_slice := []int{}
	for _, value := range new_slice {
		if value%2 == 0 {
			even_slice = append(even_slice, value)
		} else {
			odd_slice = append(odd_slice, value)
		}
	}
	fmt.Println("Even numbers:", even_slice)
	fmt.Println("Odd numbers:", odd_slice)
}
