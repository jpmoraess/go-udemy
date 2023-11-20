package main

import (
	"fmt"
	"time"
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
	channel := make(chan int, 100)

	go setList(channel)

	for v := range channel {
		fmt.Println("Recebendo: ", v)
		time.Sleep(time.Second)
	}
}

func setList(channel chan <- int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("Enviando: ", i)
	}
	close(channel)
}