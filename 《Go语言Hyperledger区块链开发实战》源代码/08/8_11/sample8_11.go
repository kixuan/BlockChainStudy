
package main
import "fmt"
 
    type Animal interface {
		cry()
	}

	type Dog struct {
		name string	
	}

	func (d Dog) cry(){
	    fmt.Printf("wangwang~, My name is %s\r\n", d.name)		
	}

	type Cat struct {
		name string	
	}

	func (c Cat) cry(){
	    fmt.Printf("miao~, My name is %s\r\n", c.name)		
	}
	
	func main() {   
		d := Dog{"Snooby"}
		d.cry()	
		c := Cat{"Mimi"}
		c.cry()	
    }
