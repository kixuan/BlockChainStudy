
package main

import "fmt"

func main() {
	sum := 0
	i := 1
LOOP: sum += i
    	fmt.Println(i)
		i++
	    if i<=5 {
			fmt.Println("+")
			goto LOOP
		}	
	fmt.Println("=")
	fmt.Println(sum)
}
