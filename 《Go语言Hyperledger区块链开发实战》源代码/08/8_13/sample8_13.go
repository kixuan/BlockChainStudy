
package main
import (
	"fmt"
	"time"
)
 
func print(str string)  {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(str)
    }
} 

func main() {   
	go print("我是协程1")
	print("我是协程2")
}
