package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* ------------------
	   fmt package

	   1. printf,%v (%v handles int,bool,string etc.)
	   	name := "Elon"
	   	fmt.Printf("Name = %v",name)

	   	>> Name = Elon

	   2. printf,%s %d
	   	name := "Elon"
	   	age := 69
	   	fmt.Printf("Name is %s",name)
	   	fmt.Printf("Age is %d",age)

	   	>> Name is Elon
	   	>> Age is 69

	   3. println
	   	fmt.Prinln("Hello", 123)

	   	fmt.Println() //new line

	   	items := []int{10,20,30}
	   	fmt.Println(items)

	   	println(items) (no 'fmt', prints ref value)

	   	>> Hello 123
	   	>> [10,20,30]
	   	>> [3/3]0xc083927498a7684

	   4. print
	   	elements := []int{999,99,9}

	   	for i := 0; i < len(elements); i++{
	   		fmt.Print(elements[i] + 1)
	   		fmt.Print(" ")
	   	}

	   	>> 1000 100 10

	   5. sprintf (like printf but returns a string, not printed out)

	   5. sprintln (like println but returns a string, not printed out)

	   ------------------ */

	/* ---------- Convert STRING <-> INT ------------

	   Package: strconv

	   STRING -> INT
	   strconv.Atoi()

	   INT -> STRING
	   strconv.Itoa()

	   ------------------------------------------------*/

	a, err := strconv.Atoi("123")
	if err != nil {
		panic(err.Error())
	}
	b := strconv.Itoa(123)

	fmt.Println(a, b)
	/* ------------------------------------------------*/

	x, y, z := 1, 2, 3
	fmt.Println(add(x, y))
	fmt.Println(add(y, z))
	fmt.Println("Hello World!")
	forloop()
}

func add(x int, y int) int {
	return x + y
}

func forloop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Printf("\n")
	}
}

func loop() {
	variables := []string{"A", "B", "C"}
	for i, j := range variables {
		fmt.Println(i, j)
	}
}

//Run
//go build
//go run hello.go
