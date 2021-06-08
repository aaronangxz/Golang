package main

import "fmt"

func main() {
	x, y, z := 1, 2, 3
	fmt.Println(add(x, y))
	fmt.Println(add(y, z))
	fmt.Println("Hello World!")
}

func add(x int, y int) int {
	return x + y
}

//Run
//go build
//go run hello.go
