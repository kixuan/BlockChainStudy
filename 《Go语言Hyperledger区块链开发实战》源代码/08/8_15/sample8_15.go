package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name      string
	Publisher string
	Pagecount uint
	Price     uint
}

func main() {
	b := Book{"Go语言Hyperledger区块链开发实战", "人民邮电出版社", 320, 50}
	data, _ := json.Marshal(b)
	fmt.Println(string(data))
}
