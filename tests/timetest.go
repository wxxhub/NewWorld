package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Nanosecond()
	fmt.Println(t)
}
