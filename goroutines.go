package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	for i := 0; i < 10000; i++ {
		go showMessage(strconv.Itoa(i))
	}
	time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))
}

func showMessage(message string) {
	fmt.Println(message)
}
