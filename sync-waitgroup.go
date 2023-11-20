package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go callDatabase(&wg)
	go callExternalApi(&wg)
	go processInternal(&wg)
	wg.Wait()
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado callDatabase")
	wg.Done()
}

func callExternalApi(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado callExternalApi")
	wg.Done()
}

func processInternal(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizado processInternal")
	wg.Done()
}
