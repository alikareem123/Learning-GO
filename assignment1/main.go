package main

import "fmt"

func main() {
	ints := []int{}
	for i := 0; i < 11; i++ {
		ints = append(ints, i)
	}
	for _, val := range ints {
		if val % 2 == 0 {
			fmt.Println("The number is even ", val)
		}else{
			fmt.Println("the number is odd ", val)
		}
	}
}
