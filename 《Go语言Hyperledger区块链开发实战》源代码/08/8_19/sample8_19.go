package main

import "fmt"

func sort(x int, y int) (int, int) {
   if x>y {
   return y, x
   }else {
      return x, y
   }
}

func main() {
   a1, b1 := sort(1, 2)
   fmt.Println(a1, b1)
   a2, b2 := sort(2, 1)
   fmt.Println(a2, b2)
}
