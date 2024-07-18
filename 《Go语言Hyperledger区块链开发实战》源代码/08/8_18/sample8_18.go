package main

import "fmt"

func swap(x *int, y *int) {
   var temp int
   temp = *x    // 保存 x 地址上的值 
   *x = *y      // 将 y 值赋给 x 
   *y = temp    // 将 temp 值赋给 y 
}
func main() {
   var x int = 1
   var y int= 2

   fmt.Printf("交换前，x = %d, y=  %d\n", x, y)   
   swap(&x, &y)
   fmt.Printf("交换后，x = %d, y=  %d\n", x, y)   
}