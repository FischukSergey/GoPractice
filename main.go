package main

import (
	"fmt"
	// "strconv"
	// "math"
)

var Global int = 1234
var AnotherGlobal = -1111

func main() {
	// var j int
	// i := Global + AnotherGlobal
	// fmt.Println("Инициализация пустой переменной j: ", j)
	// j = Global
	// // math.Abs() требует параметр float64
	// // соответственно мы приводим тип
	// k := math.Abs(float64(AnotherGlobal))
	// // fmt.Printf("Global = %d, i = %d j = %d k = %.2f,\n", Global, i, j, k)
	// //conv("молодой падаван", 50)
	newB := Digitize(100)
	fmt.Println(newB)

	// // fmt.Println(ToAlternatingCase("Проверка ТЕКСТА 1234"))
	// s:="ПРОВЕРКА TEST"
	// // fmt.Println(MyString(s).IsUpperCase())
	// fmt.Println(MultiplicationTable(1))
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

/* ТАБЛИЦА УМНОЖЕНИЯ
func MultiplicationTable(size int) [][]int {
  result:=make([][]int,size) //инициализируем строку в двумерном срезе
  for i := range result {    //инициализируем колонку в виде среза для каждого значения в строке
    result[i] = make([]int, size) //получается двумерный массив
}
	for row:=range result {
		for column:=range result[row] {
			result[row][column]=(column+1)*(row+1)
		}
	}
  return result
}

НУ ИЛИ ТАК
func MultiplicationTable(size int) [][]int {
  res := make([][]int, size)
  for i := 0 ; i < size ; i ++ {
    for x := 1 ; x < size + 1 ; x ++ {
      res[i] = append(res[i], (i + 1) * x)
      }}
  return res

}
*/
