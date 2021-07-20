package main

import "fmt"

func main() {
	fmt.Println("Level 1 Exercises")

	//Exercise 1
	// x := 42
	// y := "James Bond"
	// z := true
	// fmt.Println(x, y, z)

	//Exercise 2
	// var x int
	// var y string
	// var z bool
	// fmt.Printf("%T\n", x)
	// fmt.Printf("%T\n", y)
	// fmt.Printf("%T\n", z)

	// Exercise 3
	// var x int = 42
	// var y string = "James Bond"
	// var z bool = true

	// s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	// fmt.Println(s)

	//Exercise 4
	type magic int
	var x magic
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
