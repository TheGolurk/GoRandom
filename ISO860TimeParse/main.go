package main

import (
	"fmt"
	"time"
)

func main() {
	dateSTR := "2020-04-13T00:00:00Z"
	t, e := time.Parse("2006-01-02T15:04:05Z0700", dateSTR)
	if e != nil {
		fmt.Printf("%v\n", t)
		return
	}
	fmt.Println(t.Format("02-01-2006"))
}
