package main

import (
	"fmt"
	"strconv"
)

func main() {
	task1()
	task2()
	task3()
	task4(10, 20)
}

func task1() {
	fmt.Println("Hello, Go")
	fmt.Println("Dmitry")
	fmt.Println("2026")
}

func task2() {
	name := "Dmitry"
	age := 2026
	city := "Dar Es Salaam"
	fmt.Println(name, strconv.Itoa(age), city)
}

func task3() {
	a, b := 10, 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}

func task4(a int, b int) {
	sum := a + b
	avg := sum / 2
	fmt.Println(avg)
}
