package main
import "fmt"
 
func main() {
    Any(1)
    Any("Hello")
    Any(false)
	Any(100.1)
}
 
func Any(v interface{})  {
	fmt.Printf("您传入的数据类型：%T；值： %v\n", v, v)
}

