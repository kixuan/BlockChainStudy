
package main
import "fmt"
 
	type Person struct {
		name string
		age int
	}

	func (p Person) printName(){
	    fmt.Printf("name=%s\r\n", p.name)		
	}

	func (p Person) printAge(){
	    fmt.Printf("age=%d\r\n", p.age)		
	}
	
	func main() {   
		person := Person{"小明", 18}
		person.printName()	
		person.printAge()
    }
