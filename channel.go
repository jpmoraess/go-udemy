package main

import (
	"fmt"
)

/*
func main() {
	var list []int
	go setList(&list)

	for _, v := range list {
		fmt.Println(v)
	}
}

func setList(list *[]int) {
	for i := 0; i < 100; i++ {
		*list = append(*list, i)
	}
}
 */

func main() {
	channel := make(chan int)

	go setList(channel)

	for v := range channel {
		fmt.Println(v)
	}
}

func setList(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}