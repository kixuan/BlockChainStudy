package main
import "fmt"
func main() {    
    var arr = [5]int{1, 2, 3, 4, 5}
    for _i, _num := range arr {
		fmt.Printf("数据arr的第%d个元素为：%d\r\n", _i, _num)
    }
 }