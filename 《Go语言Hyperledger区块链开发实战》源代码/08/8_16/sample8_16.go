package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"` // 表示不进行序列化
	Pagecount int    `json:" pagecount"`
	Price     uint   `json:"price"`
}

func main() {
	var data = `{"name":"Go语言Hyperledger区块链开发实战","publisher":"人民邮电出版社"," pagecount":320,"price":50}`
	b := &Book{}
	err := json.Unmarshal([]byte(data), b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*b)
	}
}
