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
	for {
		expression := inputUser()
		sentence := strings.Join(expression, " ")
		sentence = strings.ToUpper(sentence)
		check := checkIntOrRom(expression)

		if check == 7 {
			fmt.Println("inout error: expression must be of two operands and one operator")
		} else if check == 1 { // выражение с целыми числами
			resultInt, err := operations(intNum(expression))
			if err == nil {
				fmt.Println(sentence, " = ", resultInt)
			} else {
				fmt.Println(err)
			}

		} else if check == 2 { // выражение с римскими числами
			roman, err := operations(toRomanNum(expression))
			resultRoman := toIntNum(roman)
			if err == nil {
				fmt.Println(sentence, " = ", resultRoman)
			} else {
				fmt.Println(err)
			}

		} else if check == 0 { // в выражении разные типы чисел
			fmt.Println("input error: operands are different types")
		}
	}
}

func inputUser() []string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nCalc produces one action (+, -, *, /) with integer numbers (1-10) or roman numerals (I-X).   ")
	fmt.Println("Enter your expression separated by space \n(Example: 5 * 2, x - ii):")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	arr := strings.Split(text, " ")
	return arr
}

func intNum(arr []string) (a, b int, c string) {
	num1 := strings.TrimSpace(arr[0])
	num2 := strings.TrimSpace(arr[2])
	a, _ = strconv.Atoi(num1)
	b, _ = strconv.Atoi(num2)
	c = arr[1]
	return a, b, c

}

func checkIntOrRom(arr []string) int {

	if len(arr) != 3 {
		return 7
	}

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

func toRomanNum(arr []string) (a, b int, c string) {
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
	arr[0] = strings.ToLower(arr[0])
	arr[2] = strings.ToLower(arr[2])
	a = convertRoman[arr[0]]
	b = convertRoman[arr[2]]
	c = arr[1]
	return a, b, c
}

func toIntNum(res int) string {
	if res > 0 {
		convertInt1 := map[int]string{
			0: "",
			1: "I",
			2: "II",
			3: "III",
			4: "IV",
			5: "V",
			6: "VI",
			7: "VII",
			8: "IIX",
			9: "IX",
		}
		convertIntX := map[int]string{
			0:  "",
			1:  "X",
			2:  "XX",
			3:  "XXX",
			4:  "XL",
			5:  "L",
			6:  "LX",
			7:  "LXX",
			8:  "LXXX",
			9:  "XC",
			10: "C",
		}
		resX := res / 10
		res1 := res % 10
		res11 := convertInt1[res1]
		resXX := convertIntX[resX]

		fin := resXX + res11
		return fin
	} else {
		return "Error: Roman numerals cannot be less than zero"
	}
}

func operations(a, b int, c string) (int, error) {
	var result int
	if a <= 10 && a > 0 && b <= 10 && b > 0 {
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
	} else {
		return 0, errors.New("input error: incorrect operand (1-10 or I-X)")
	}
}
