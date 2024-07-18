
package main
import "fmt"
func main() {    
	var  a int = 5
    var ip *int = &a

	fmt.Printf("a 变量的地址是: %x\n", &a )
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	fmt.Printf("*ip 变量的值: %d\n", *ip )
}