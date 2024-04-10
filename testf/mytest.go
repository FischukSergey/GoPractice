package mytest

import (
	"fmt"
	"strings"
	"unicode"
)

// func main() {
//fmt.Printf("Фраза наоборот %v\n", ReverseWords("This is 12345678"))
//fmt.Printf("Средний символ %s\n", GetMiddle("а"))
//fmt.Println(MyString("GGGHHHH").IsUpperCase())
// fmt.Println(ToAlternatingCase("Проверка ТЕКСТА 1234"))
// }

// функция меняет строчные буквы на заглавные
func ToAlternatingCase(str string) string {
	newstr := []rune(str)
	fmt.Println(newstr)
	reverseRune := ""
	for _, invert := range newstr {
		if unicode.IsLower(invert) {
			reverseRune += string(unicode.ToUpper(invert))
		} else {
			reverseRune += string(unicode.ToLower(invert))
		}
	}
	return reverseRune
}

// Проверка на наличие в строке только прописных букв
type MyString string

func (s MyString) IsUpperCase() bool {
	// Your code here!
	fmt.Println(strings.ToUpper(string(s)))
	return s == MyString(strings.ToUpper(string(s)))
}

/*
func Digit(n int) []int {
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
*/
// функция выбирает из строки 1 символ посредине если длина строки нечетная и 2 символа если четная
func GetMiddle(s string) string {
	simbol := []rune(s)
	chet := len(simbol) % 2 //проверка на чет нечет
	fmt.Println(s, chet, len(simbol))
	if chet == 0 {
		return string(simbol[len(simbol)/2-1 : len(simbol)/2+1])
	}
	return string(simbol[len(simbol)/2 : len(simbol)/2+1])
}

// функция реверсирует каждое слово в строке, но работает только если символ один байт
func ReverseWords(str string) string {
	words := strings.Fields(str)
	newstr := ""
	for _, word := range words {
		//fmt.Println(words, len(word))
		dwor := ""
		for i := len(word) - 1; i >= 0; i-- {
			//fmt.Println(i, word[i])
			dwor += string(word[i])
		}
		newstr += dwor + " "
	}
	return newstr // reverse those words
}

/* Пример двух способов разбивки целого числа на срез из цифр
func DigPow(n, p int) int {
	var sliceD []string
	var sum float64=0
	result:=0
	//   var sliceDR []int // первый способ через создание среза из цифр в обратном порядке путем получения
	//   for i:=0; n>0; i++ {  // остатка от деления на 10, а затем реверс среза
	//     sliceD=append(sliceD, n%10)
	//     n=n/10
	//   }
	//   for i:=len(sliceD); i>0; i-- {
	//     sliceDR=append(sliceDR,sliceD[i-1])
	//   }
	sliceD=strings.Split(strconv.Itoa(n),"") // второй способ через перевод в строку цифры и разделение строки
	fmt.Println(sliceD)
	for i, val:=range sliceD {  // на слайс
		base, _:=strconv.Atoi(val)
		if i==0 {
			sum=math.Pow(float64(base),float64(p))
		} else {
			sum+=math.Pow(float64(base),float64(p+i))
		}
	}
	if int(sum)%n==0 {
		result=int(sum)/n
	} else {
		result = -1
	}
	//   fmt.Println(sum, result)
	return result
*/
// Функция замены всех символов в строке, в данном случае, всех гласных
// func Disemvowel(comment string) string {
//   return regexp.MustCompile(`(?i)[aeiou]`).ReplaceAllString(comment, "")
// }
// или такой вариант
// func Disemvowel(comment string) string {
// var replacer = strings.NewReplacer(
// 		"a", "",
// 		"e", "",
// 		"i", "",
// 		"o", "",
// 		"u", "",
// 		"A", "",
// 		"E", "",
// 		"I", "",
// 		"O", "",
// 		"U", "",
// 	)
// 	return replacer.Replace(comment)

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


