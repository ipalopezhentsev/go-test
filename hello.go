package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 4.0
	b := 5.0
	fmt.Println("Hello World! Square root of", a, "is", math.Sqrt(a))
	fmt.Println(b)
	c := PublicFun(4)
	fmt.Println(c)

	var p Person = Person{Name: "Ilya", Age: 40}
	fmt.Println("Person is", p)
	fmt.Println("Next year", p.Name, "will be", p.Age+1)

	passByValue(p)
	fmt.Println(p)
	passByRef(&p)
	fmt.Println(p)

	var arr []int = []int{1, 2}
	arr = append(arr, 3)
	fmt.Println(len(arr))

	if len(arr) == 3 {
		fmt.Println("right")
	}

	if l := len(arr); l == 3 {
		fmt.Println(l)
	}

	go func() int {
		fmt.Println("From goroutine")
		return 5
	}()
}

type Person struct {
	Name string
	Age  int
}

func passByValue(p Person) {
	p.Age = p.Age + 1
}

func passByRef(p *Person) {
	p.Age = p.Age + 1
}
