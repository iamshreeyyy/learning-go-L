package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to the time video")

	presntTime := time.Now()
	fmt.Println(presntTime)

	fmt.Println(presntTime.Format("01-02-2006  15:04:05 Monday")) ///this how to use formmat this should be the exact numb

	createdDate := time.Date(2020, time.September, 20, 23, 23, 0, 23, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006  15:04:05 Monday"))

}
