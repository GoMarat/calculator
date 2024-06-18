package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romeArabicMap = map[string]int{
	"i":    1,
	"ii":   2,
	"iii":  3,
	"iv":   4,
	"v":    5,
	"vi":   6,
	"vii":  7,
	"viii": 8,
	"ix":   9,
	"x":    10,
}

var arabicRomeMap = map[int]string{
	1:   "i",
	2:   "ii",
	3:   "iii",
	4:   "iv",
	5:   "v",
	6:   "vi",
	7:   "vii",
	8:   "viii",
	9:   "ix",
	10:  "x",
	11:  "xi",
	12:  "xii",
	13:  "xiii",
	14:  "xiv",
	15:  "xv",
	16:  "xvi",
	17:  "xvii",
	18:  "xviii",
	19:  "xix",
	20:  "xx",
	21:  "xxi",
	24:  "xxiv",
	25:  "xxv",
	27:  "xxvii",
	28:  "xxviii",
	30:  "xxx",
	32:  "xxxii",
	35:  "xxxv",
	36:  "xxxvi",
	40:  "xl",
	42:  "xlii",
	45:  "xlv",
	48:  "xlviii",
	49:  "xlix",
	50:  "l",
	54:  "liv",
	56:  "lvi",
	60:  "lx",
	63:  "lxiii",
	64:  "lxiv",
	70:  "lxx",
	72:  "lxxii",
	80:  "lxxx",
	81:  "lxxxi",
	90:  "xc",
	100: "c",
}

func sum(x, y int) int {
	return x + y
}

func dif(x, y int) int {

	return x - y
}

func mult(x, y int) int {
	return x * y
}

func dev(x, y int) int {
	return x / y
}

func separatonOfElements(mathString []string) (first, second int) {
	firstString := strings.TrimSpace(mathString[0])
	firstString = strings.ToLower(firstString)
	first, err := strconv.Atoi(firstString)
	if err != nil {
		first = romeArabicMap[firstString]
	}
	secondString := strings.TrimSpace(mathString[1])
	secondString = strings.ToLower(secondString)
	if (romeArabicMap[firstString] == 0 && romeArabicMap[secondString] != 0) || (romeArabicMap[secondString] == 0 && romeArabicMap[firstString] != 0) {
		panic("Оба числа должны быть либо арабскими либо римскими")
	}
	second, err = strconv.Atoi(secondString)
	if err != nil {
		second = romeArabicMap[secondString]
	}
	return
}
func elementCountError(el []string) {
	if len(el) > 2 {
		panic("приложение работает только с двумя целыми числами")
	}
}

func NumberCheck(x, y int) {
	if x > 10 || x < 1 || y > 10 || y < 1 {
		panic("Калькулятор принимает на вход числа от 1 до 10 включительно")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите выражение для расчета")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if !strings.Contains(text, "+") && !strings.Contains(text, "-") && !strings.Contains(text, "*") && !strings.Contains(text, "/") {
			panic("выражение должно содержать арифмитическое действие")
		}

		if strings.Contains(text, "+") {
			elements := strings.Split(text, "+")
			elementCountError(elements)
			firstElement, secondElement := separatonOfElements(elements)
			NumberCheck(firstElement, secondElement)
			if romeArabicMap[strings.TrimSpace(elements[0])] == 0 {
				res := sum(firstElement, secondElement)
				fmt.Println(res)
			} else {
				res := arabicRomeMap[sum(firstElement, secondElement)]
				fmt.Println(res)
			}

		}
		if strings.Contains(text, "-") {
			elements := strings.Split(text, "-")
			elementCountError(elements)
			firstElement, secondElement := separatonOfElements(elements)
			if romeArabicMap[strings.TrimSpace(elements[0])] > 0 || romeArabicMap[strings.TrimSpace(elements[0])] <= romeArabicMap[strings.TrimSpace(elements[1])] {
				panic("Невозможно выдать результат так как нет римских 0 или отрицательных чисел")
			}
			NumberCheck(firstElement, secondElement)
			if romeArabicMap[strings.TrimSpace(elements[0])] == 0 {
				res := dif(firstElement, secondElement)
				fmt.Println(res)
			} else {
				res := arabicRomeMap[dif(firstElement, secondElement)]
				fmt.Println(res)
			}
		}
		if strings.Contains(text, "*") {
			elements := strings.Split(text, "*")
			elementCountError(elements)
			firstElement, secondElement := separatonOfElements(elements)
			NumberCheck(firstElement, secondElement)
			if romeArabicMap[strings.TrimSpace(elements[0])] == 0 {
				res := mult(firstElement, secondElement)
				fmt.Println(res)
			} else {
				res := arabicRomeMap[mult(firstElement, secondElement)]
				fmt.Println(res)
			}
		}
		if strings.Contains(text, "/") {
			elements := strings.Split(text, "/")
			elementCountError(elements)
			firstElement, secondElement := separatonOfElements(elements)
			NumberCheck(firstElement, secondElement)
			if romeArabicMap[strings.TrimSpace(elements[0])] == 0 {
				res := dev(firstElement, secondElement)
				fmt.Println(res)
			} else {
				res := arabicRomeMap[dev(firstElement, secondElement)]
				fmt.Println(res)
			}
		}
	}
}
