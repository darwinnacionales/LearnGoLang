package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	tf := t.Format(time.Kitchen)
	fmt.Println(tf)

	// For custom format, we use the following as reference:
	// 01/02 03:04:05PM '06 -0700

	fmt.Println(t.Format("2006-01-02"))

	fmt.Println(t.Format("January 2, 2006"))

	fmt.Println(t.Format("January 2, 2006 03:04:05PM"))
}
