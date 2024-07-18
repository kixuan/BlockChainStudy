package main

import (
	"fmt"
	"time"
)

func main() {
    t := time.Now()
	var weekday int = int(t.Weekday())
	var str_weekday = ""
	switch weekday {
	case 1:
		str_weekday = "星期一"
	case 2:
		str_weekday = "星期二"
	case 3:
		str_weekday = "星期三"
	case 4:
		str_weekday = "星期四"
	case 5:
		str_weekday = "星期五"
	case 6:
		str_weekday = "星期六"
	case 7:
		str_weekday = "星期日"
	}
    fmt.Println("今天是："+str_weekday)
}
