package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	expression := inputUser()
	check := checkInt(expression)

	if check == 1 { // выражение с целыми числами
		fmt.Println("int")

	} else if check == 2 { // выражение с римскими числами
		roman, err := operations(romanNum(expression))
		if err == nil {
			fmt.Println(expression, " = ", roman)
		} else {
			fmt.Println(err)
		}
		fmt.Println()

	} else if check == 0 { // в выражении разные типы чисел
		fmt.Println("input error: operands are different types")
	}
}

func inputUser() []string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Calc produces one action (+, -, *, /) with integer numbers (1-10) or roman numerals (I-X). \nExample: 5 * 2, x - ii  ")
	fmt.Println("Enter your expression: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	arr := strings.Split(text, " ")
	return arr
}

func checkInt(arr []string) int {

	num1 := strings.TrimSpace(arr[0])
	num2 := strings.TrimSpace(arr[2])
	var digit1 bool
	var digit2 bool
	if _, err := strconv.Atoi(num1); err == nil {
		digit1 = true
	} else {
		digit1 = false
	}

	if _, err := strconv.Atoi(num2); err == nil {
		digit2 = true
	} else {
		digit2 = false
	}

	if digit1 == true && digit2 == true {
		return 1
	} else if digit1 == false && digit2 == false {
		return 2
	} else {
		return 0
	}

}

func romanNum(arr []string) (a, b int, c string) {
	convertRoman := map[string]int{
		"i":   1,
		"ii":  2,
		"iii": 3,
		"iv":  4,
		"v":   5,
		"vi":  6,
		"vii": 7,
		"iix": 8,
		"ix":  9,
		"x":   10,
	}
	a = convertRoman[arr[0]]
	b = convertRoman[arr[2]]
	c = arr[1]
	return a, b, c
}

func operations(a, b int, c string) (int, error) {
	var result int
	switch c {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		return 0, errors.New("input error: incorrect operator")
	}
	return result, nil
}
