package main

import (
	"fmt"
	"math"
)

var Global int = 1234
var AnotherGlobal = -1111

func main() {
	var j int
	i := Global + AnotherGlobal
	fmt.Println("Инициализация пустой переменной j: ", j)
	j = Global
	// math.Abs() требует параметр float64
	//соответсвенно мы приводим тип
	k := math.Abs(float64(AnotherGlobal))
	fmt.Printf("Global = %d, i = %d j = %d k = %.2f,\n", Global, i, j, k)
	//conv("молодой падаван", 50)
	newB := Digitize(100)
	fmt.Println(newB)

	fmt.Println(ToAlternatingCase("Проверка ТЕКСТА 1234"))
	s:="ПРОВЕРКА TEST"
	fmt.Println(MyString(s).IsUpperCase())
}

func Digitize(n int) []int {
	// your code here
	var b []int
	if n == 0 {
		b = append(b, 0)
	}
	for i := 0; n > 0; i++ {
		b = append(b, n%10)
		n = n / 10
	}
	return b
}
