package main

import (
	"os"
)

func main() {
	_, err := os.Open("C://Users//settings.txt")

	if err != nil {
		panic(err)
	}
}
