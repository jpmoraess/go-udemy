package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	i := 0
	for x := 0; x < 10000; x++ {
		go func() {
			m.Lock()
			i++
			m.Unlock()
		}()
	}
	time.Sleep(time.Second * 6)
	fmt.Println(i)
}

func ChangeNumber(i *int, newNumber int) {
	*i = newNumber
}