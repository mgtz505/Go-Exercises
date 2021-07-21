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
	// type magic int
	// var x magic
	// fmt.Println(x)        //Prints the zero value
	// fmt.Printf("%T\n", x) //Type
	// x = 42
	// fmt.Println(x)

	//Exercise 5
	// var y magic
	// y := int(x)

	// fmt.Printf("%T\n", y)
	// fmt.Println(y)

	fmt.Println("Level 2 Exercises")

	//Exercise 1
	// a := 42
	// fmt.Printf("%d\t%b\t%#x", a, a, a)

	//Exercise 2
	// a := (42 == 42)
	// b := (42 <= 42)
	// c := (42 >= 46)
	// d := (0 < 100)
	// e := (99 > 92)
	// fmt.Println(a, b, c, d, e)

	//Exercise 3
	// const (
	// 	a     = 42
	// 	b int = 43
	// )
	// fmt.Println(a, b)

	//Exercise 4
	// var x int
	// x = 32
	// fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	// y := x << 1
	// fmt.Printf("%d\t%b\t%#x\n", y, y, y)

	//Exercise 5
	// a := `Here is a "raw string literal"`
	// fmt.Println(a)

	//Exercise 6

	// const (
	// 	a = 2017 + iota
	// 	b
	// 	c
	// 	d
	// )
	// fmt.Println(a, b, c, d)

	fmt.Println("Level 3 Exercises")

	//Exercise 1
	// 	for i := 1; i <= 10_000; i++ {
	// 		fmt.Println(i)
	// 	}

	//Exercise 2
	// for i := 65; i <= 90; i++ {
	// 	fmt.Printf("%#U\n", i)
	// }

	// Exercise 3
	counter := 0
	for i := 2021; i > 1994; i-- {
		fmt.Println(i)
		counter++
	}
	fmt.Println("Marcus is", counter, "years old!")

	count := 0
	by := 1994
	for {
		if by == 2021 {
			break
		} else {
			by++
			count++
		}
	}
	fmt.Println(count)

	//Exercise 4

	for i := 10; i <= 100; i++ {
		if i%4 == 0 {
			fmt.Println(i)
		}
	}

	//Exercise 5
	switch {
	case (2 == 2):
		fmt.Println("2 is 2!")
	case (3 == 3):
		fmt.Println("3 is 3!")
	}

	//Exercise 6
	favSport := "swimming"
	switch favSport {
	case "skiing":
		fmt.Println("Skiing")
	case "swimming":
		fmt.Println("Glug glug")
	}
}
