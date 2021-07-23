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
	// counter := 0
	// for i := 2021; i > 1994; i-- {
	// 	fmt.Println(i)
	// 	counter++
	// }
	// fmt.Println("Marcus is", counter, "years old!")

	// count := 0
	// by := 1994
	// for {
	// 	if by == 2021 {
	// 		break
	// 	} else {
	// 		by++
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	//Exercise 4

	// for i := 10; i <= 100; i++ {
	// 	if i%4 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//Exercise 5
	// switch {
	// case (2 == 2):
	// 	fmt.Println("2 is 2!")
	// case (3 == 3):
	// 	fmt.Println("3 is 3!")
	// }

	//Exercise 6
	// 	favSport := "swimming"
	// 	switch favSport {
	// 	case "skiing":
	// 		fmt.Println("Skiing")
	// 	case "swimming":
	// 		fmt.Println("Glug glug")
	// 	}

	fmt.Println("Level 4 Exercises")

	//Exercise 1
	// arr := [5]int{1, 2, 3, 4, 5}

	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }
	// fmt.Printf("%T\n", arr)

	//Exercise 2
	// sl := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// for i := range sl {
	// 	fmt.Println(i)
	// }
	// fmt.Printf("%T\n", sl)

	//Exercise 3

	// sl2 := sl[1:6]
	// fmt.Println(sl2)
	// sl3 := sl[2:9]
	// fmt.Println(sl3)
	// sl4 := sl[2:len(sl)]
	// fmt.Println(sl4)

	//Exercise 4

	// x := []int{42, 43, 44, 45, 56, 47, 48, 49, 50, 51}
	// x = append(x, 52)
	// fmt.Println(x)
	// x = append(x, 53, 54, 55)
	// fmt.Println(x)
	// y := []int{56, 57, 58, 59, 60}
	// x = append(x, y...) //Can't forget the ... ellipse
	// fmt.Println(x)

	//Exercise 5
	// x := []int{42, 43, 44, 45, 46, 47, 38, 39, 50, 51}
	// x = append(x[:3], 48, 49)
	// fmt.Println(x)

	// Exercise 6
	// y := make([]int, 50, 50)
	// fmt.Println(len(y))
	// fmt.Println(cap(y))

	//Exercise 7
	// james := []string{"James", "Bond"}
	// mission := []string{"Tom", "Cruise"}
	// jamesMission := [][]string{james, mission}
	// fmt.Println(jamesMission)

	// for n, v := range jamesMission {
	// 	fmt.Println(n, v)
	// 	for j, v := range v {
	// 		fmt.Println(j, v)
	// 	}
	// }

	//Exercise 8
	// m := map[string]string{
	// 	"Joe":  "Schmo",
	// 	"Jane": "Dough",
	// }
	// fmt.Println(m)
	// for n, v := range m {
	// 	fmt.Println(n, v)
	// }

	fmt.Println("Level 5 Exercises")

	//Exercise 1
	type person struct {
		firstName       string
		lastName        string
		iceCreamFlavors []string
	}
	barb := person{
		firstName:       "Barbara",
		lastName:        "Grahmm",
		iceCreamFlavors: []string{"Chocolate", "Lime"},
	}
	fmt.Println(barb)
	rich := person{
		firstName:       "Richard",
		lastName:        "Washington",
		iceCreamFlavors: []string{"Vanilla", "Chocomint"},
	}
	fmt.Println(rich)

	//Exercise 2
	m := map[string]person{
		barb.firstName: barb,
		barb.lastName:  barb,
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.firstName)
	}
	//Exercise 3
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "grey",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}

	fmt.Println(t, s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)

	//Exercise 4

	d1 := struct {
		breed string
		age   int
		size  int
	}{
		breed: "Shiba Inu",
		age:   7,
		size:  27,
	}
	fmt.Println(d1)

}
