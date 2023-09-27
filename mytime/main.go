package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")

	presentTime := time.Now()

	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006  Monday"))
	createdDate := time.Date(2022, time.May, 5, 5, 23, 43, 1, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("2006-01-02 Monday"))
}
