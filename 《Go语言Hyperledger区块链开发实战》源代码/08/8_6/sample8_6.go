package main

import "fmt"

func main() {
	var scoreMap map[string] uint //记录学生的分数
	scoreMap = make(map[string] uint)
    scoreMap["小强"] = 100
	scoreMap["小刚"] = 95
    scoreMap["小红"] = 80
			
	fmt.Printf("小强的分数为：%d\r\n", scoreMap["小强"])
	fmt.Printf("小刚的分数为：%d\r\n", scoreMap["小刚"])
	fmt.Printf("小红的分数为：%d\r\n", scoreMap["小红"])
}
