package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	delay, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("done")
}
