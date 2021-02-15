package main

import (
	"fmt"
	"time"
)

func main() {
	retry := float32(10)
	wait := time.Duration(retry) * time.Second
	fmt.Println("waiting")
	time.Sleep(wait)
	fmt.Println("done")
}
