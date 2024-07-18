package main
import "fmt"
func main() {    
    arr := make([]int, 3,6)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
    fmt.Printf("切片arr的的长度为：%d，容量为%d, slice=%v\r\n", len(arr), cap(arr), arr)
 }
