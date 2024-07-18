package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 5; i++ {
			sum += i
			fmt.Println(i)
			if i<5 {
				fmt.Println("+")
			}
	}
	fmt.Println("=")
	fmt.Println(sum)
}
