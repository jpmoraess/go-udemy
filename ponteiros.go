package main

import "fmt"

func main3() {
	x := 5
	y := &x
	*y = 10
	fmt.Println("main")
	fmt.Println(x, *y)
	fmt.Println(&x, y)
	ImprimirValores(&x, y)
}

func ImprimirValores(x *int, y *int) {
	fmt.Println("ImprimirValores")
	fmt.Println(x, y)
}
