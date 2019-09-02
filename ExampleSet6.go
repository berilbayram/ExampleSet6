package main

import (
	"fmt"
	"math"
)

type person struct {
	first string
	last string
	age int
}
func (p person) speak (){
	fmt.Println("I am ", p.first, p.last, " and I am " , p.age, " years old")
}

type circle struct {
	radius float64
}
func (c circle ) area() float64 {
	return math.Phi * c.radius
}

type square struct {
	length float64
}

func (s square) area() float64{
	return math.Pow(2,s.length)
}

type shape interface{
	area() float64
}

func info (s shape)  {
	fmt.Println("Area of the shape is ", s.area())

}


func main() {

	fmt.Println("Exercise 1")
	fmt.Println(foo())
	fmt.Println(bar())

	fmt.Println("\nExercise 2")
	numberSet := []int {2,6,5,8,7,9,5,6,5,3,2,4,5,5}
	numberSet2 := []int {2,6,5,89,74,4,6,5,8,4,5,2,6,5,4,5}
	x := foo2(numberSet...)
	fmt.Println(x)
	y := bar2(numberSet2)
	fmt.Println(y)

	fmt.Println("\nExercise 3")
	defer foo3()
	bar3()

	fmt.Println("\nExercise 4")

	person1 := person{
		first: "Beril",
		last:  "Bayram",
		age:   21,
	}

	person1.speak()

	fmt.Println("\nExercise 5")

	s := square {
		length: 5.0,
	}
	fmt.Println("Area of the square is ", s.area())
	fmt.Println("With info: ")
	info(s)

	c := circle {
		radius:5.0,
	}
	fmt.Println("Area of the circle is ", c.area())
	fmt.Println("With info: ")
	info(c)

	fmt.Println("\nExercise 6")
	func() {
		for i := 0; i < 21 ; i++ {
			if i % 3 == 0 {
				fmt.Println(i)
			}
		}
	}()

	fmt.Println("\nExercise 7")
	f := func() {
		for i := 0; i < 5 ; i++ {
			fmt.Println(i)
		}
	}
	f()
	fmt.Printf("%T\n", f)

	fmt.Println("\nExercise 8")
	fun := foo4()
	fmt.Println(fun())

	fmt.Println("\nExercise 9")

	g := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi) - 1]
	}

	h := foo5(g, []int {1,2,3,4,5,6})
	fmt.Println(h)

	fmt.Println("\nExercise 10")
	k := increment()
	fmt.Println(k())
}



func foo() int  {
	x := 42
	return x
}

func bar() (string, int)  {
	return "Bond, James Bond", 007
}

func foo2 (xi ...int) int {
	sum := 0
	for _,v := range xi {
		sum = sum + v
	}
	return sum
}

func bar2 (xi []int) int {
	sum := 0
	for _,v := range xi {
		sum = sum + v
	}
	return sum
}

func foo3() {
	fmt.Println( "Hello World!")
}

func bar3()  {
	fmt.Println("Hello Mars!")
}

func foo4() func() string {
	return func() string  {
		return "Have you met Ted?."
	}
}
func foo5(f func(xi []int) int , ii [] int) int {
	number := f(ii)
	number = number + 1
	return number
}
func increment()  func() int {
	x := 0
	return func() int {
		x = x + 1
		return x
	}

}
