package main

import "fmt"

type Book struct {
    Name string;
    Publisher string;
    Pagecount uint;
    Price uint;
}


func main() {

    // 创建一个新的结构体
    fmt.Println(Book{"Go语言Hyperledger区块链开发实战", "人民邮电出版社", 350, 90})
}
