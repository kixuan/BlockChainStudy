package main
import "fmt"	

func sum(x int, y int) int {
	return x+y
}
func main() {   
	var result = sum(1,2)
	fmt.Printf("result = %d", result)
}
