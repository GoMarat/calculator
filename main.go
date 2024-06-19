package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

var romeArabicMap = map[string]int{
    "I": 1,
    "II": 2,
    "III": 3,
    "IV": 4,
	"V": 5,
	"VI": 6,
    "VII": 7,
    "VIII": 8,
    "IX": 9,
	"X": 10,
}

var arabicRomeMap = map[int]string{
    1: "I",
    2: "II",
    3: "III",
    4: "IV",
	5: "V",
	6: "VI",
    7: "VII",
    8: "VIII",
    9: "IX",
	10: "X",
	11: "XI",
    12: "XII",
    13: "XIII",
    14: "XIV",
	15: "XV",
	16: "XVI",
    17: "XVII",
    18: "XVIII",
    19: "XIX",
	20: "XX",
	21: "XXI",
    24: "XXIV",
	25: "XXV",
	27: "XXVII",
    28: "XXVIII",
    30: "XXX",
	32: "XXXII",
    35: "XXXV",
	36: "XXXVI",
    40: "XL",
	42: "XLII",
    45: "XLV",
	48: "XLVIII",
    49: "XLIX",
	50: "L",
	54: "LIV",
	56: "LVI",
    60: "LX",
	63: "LXIII",
    64: "LXIV",
	70: "LXX",
	72: "LXXII",
    80: "LXXX",
	81: "LXXXI",
    90: "XC",
	100: "C",
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

	first, err := strconv.Atoi(firstString)
	if err != nil {
		first = romeArabicMap[firstString]
	}
	secondString := strings.TrimSpace(mathString[1])
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
	if len(el) > 2  {
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
		text = strings.ToUpper(text)
		if !strings.Contains(text, "+") && !strings.Contains(text, "-") && !strings.Contains(text, "*") && !strings.Contains(text, "/") {
			panic("выражение должно содержать арифмитическое действие")
		}
				
		if strings.Contains(text, "+") {
			elements := strings.Split(text, "+")
			elementCountError(elements)
			firstElement, secondElement := separatonOfElements(elements)
			NumberCheck(firstElement, secondElement)
			if romeArabicMap[strings.TrimSpace(elements[0])] == 0 {
			res :=	sum(firstElement, secondElement)
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
			_, ok := romeArabicMap[strings.TrimSpace(elements[0])]
			if ok  && romeArabicMap[strings.TrimSpace(elements[0])] <= romeArabicMap[strings.TrimSpace(elements[1])] {
				panic("Невозможно выдать результат так как нет римских 0 или отрицательных чисел")
			}
			NumberCheck(firstElement, secondElement)
			if romeArabicMap[strings.TrimSpace(elements[0])] == 0 {
				res :=	dif(firstElement, secondElement)
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
				res :=	mult(firstElement, secondElement)
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
				res :=	dev(firstElement, secondElement)
				fmt.Println(res)
				} else {
					res := arabicRomeMap[dev(firstElement, secondElement)]
					fmt.Println(res)
				}
		}
}
}