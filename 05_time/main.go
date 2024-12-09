package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study of go lang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	date := time.Date(2020, time.December, 22, 23, 23, 0, 0, time.Local)
	fmt.Println(date)
	fmt.Println(date.Format("01-02-2006 Moday"))
}
