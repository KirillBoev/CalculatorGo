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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите матеманическое выражение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		a, _ := calc(text)
		fmt.Println(a)
	}
}

var num = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numInv = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var maxTable = []int{
	1000,
	900,
	500,
	400,
	100,
	90,
	50,
	40,
	10,
	9,
	5,
	4,
	1,
}

func highestDecimal(n int) int {
	for _, v := range maxTable {
		if v <= n {
			return v
		}
	}
	return 1
}

func toNumber(n string) int {
	out := 0
	ln := len(n)
	for i := 0; i < ln; i++ {
		c := string(n[i])
		vc := num[c]
		if i < ln-1 {
			cnext := string(n[i+1])
			vcnext := num[cnext]
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}
	return out
}

func toRoman(n int) string {
	out := ""
	for n > 0 {
		v := highestDecimal(n)
		out += numInv[v]
		n -= v
	}
	return out
}

func isRoman(number string) bool {
	_, ok := num[string(number[0])]
	return ok
}

func calc(input string) (text string, error error) {
	defer func() {
		if error != nil {
			fmt.Println(error)
			os.Exit(1)
		}
	}()

	operation := [4]string{"+", "-", "*", "/"}
	j := -1

	for i := 0; i < len(operation); i++ {
		if strings.Contains(input, operation[i]) {
			j = i
		}
	}

	if j == -1 {
		err := errors.New("Неверное матаматическое выражение, попробуйте снова")
		return "", err
	}

	arr := strings.Split(input, operation[j])
	if len(arr) != 2 {
		err := errors.New("Неверный формат математической операции - два операнда и один оператор")
		return "", err
	}

	if isRoman(arr[0]) == isRoman(arr[1]) {
		var a, b int
		isRoman := isRoman(arr[0])
		if isRoman {
			a = toNumber(arr[0])
			b = toNumber(arr[1])
		} else {
			a, _ = strconv.Atoi(arr[0])
			b, _ = strconv.Atoi(arr[1])
		}

		var result int
		if (a > 0 && b > 0) && (a <= 10 && b <= 10) {
			switch operation[j] {
			case "+":
				result = (a + b)
			case "-":
				result = (a - b)
			case "*":
				result = (a * b)
			default:
				result = (a / b)
			}

			if isRoman {
				if result > 0 {
					return toRoman(result), nil
				} else {
					err := errors.New("В римской системе нет отрицательных чисел")
					return "", err
				}
			} else {
				return strconv.Itoa(result), nil
			}
		} else {
			err := errors.New("Введите число от 1 до 10")
			return "", err
		}
	} else {
		err := errors.New("Вводимые числа должны быть в одной системе исчисления")
		return "", err
	}
}
