package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range s {
		if number%2 == 0 {
			fmt.Println("number %v is even", number)
		} else {
			fmt.Println("number %v is odd", number)
		}
	}
}
