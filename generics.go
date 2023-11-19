package main

import "fmt"

func reverse(slice []int) []int {
	length := len(slice) - 1
	newInts := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		newInts[length] = slice[i]
		length--
	}
	return newInts
}

func reverse2(slice []string) []string {
	length := len(slice) - 1
	newStrings := make([]string, len(slice))
	for i := 0; i < len(slice); i++ {
		newStrings[length] = slice[i]
		length--
	}
	return newStrings
}

type AcceptedTypes interface { // constraint
	int | string
}

func reverseGenerics[T AcceptedTypes](slice []T) []T {
	length := len(slice) - 1
	newSlice := make([]T, len(slice))
	for i := 0; i < len(slice); i++ {
		newSlice[length] = slice[i]
		length--
	}
	return newSlice
}

func main() {
	slice := []int { 4, 6, 1, 2}
	slice2 := []string {"A", "N", "D"}
	/*
	fmt.Println(reverse(slice))
	fmt.Println(reverse2(slice2))
	 */
	fmt.Println(reverseGenerics(slice))
	fmt.Println(reverseGenerics(slice2))
}

