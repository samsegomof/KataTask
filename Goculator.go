package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var RomMap = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40,
	"L":    50,
	"XC":   90,
	"C":    100,
}

var ConvToRoman = []int{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	40,
	50,
	90,
	100,
}

const (
	LOWCOUNT  = "Строка не является математической операцией"
	HIGHCOUNT = "Формат математической операции не удовлетворяет условию задачи"
	SCALE     = "Нельзя использовать одновременно разные системы счисления"
	NEGATIVE  = "В римской системе нет отрицательных чисел"
	ZERO      = "В римской системе нет числа ноль"
	RANGE     = "Калькулятор принимает числа от 0 до 10"
)

func OperandsError(mathTask string) {
	re := regexp.MustCompile("[+\\-*/]")
	operands := re.Split(mathTask, -1)
	if len(operands) < 2 {
		panic(LOWCOUNT)
	}
	if len(operands) > 2 {
		panic(HIGHCOUNT)
	}
}

func DetectOperation(mathTask string) string {
	if strings.Contains(mathTask, "+") {
		return "+"
	} else if strings.Contains(mathTask, "-") {
		return "-"
	} else if strings.Contains(mathTask, "*") {
		return "*"
	} else {
		return "/"
	}
}

func ParseRomToInt(result int) {
	var RomanElement string
	if result == 0 {
		panic(ZERO)
	} else if result < 0 {
		panic(NEGATIVE)
	} else {
		for _, elem := range ConvToRoman {
			for i := elem; i <= result; {
				for idx, value := range RomMap {
					if value == elem {
						RomanElement += idx
						result -= elem
					}
				}
			}
		}
	}
	fmt.Println(RomanElement)
}

func calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func parse(task string) {
	var romans []int
	var stringsFound int
	OperandsError(task)
	mathTask := strings.Split(task, " ")
	for i, elem := range mathTask {
		if i == 1 {
			continue
		}
		_, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
		}
	}
	switch stringsFound {
	case 0:
		num1, _ := strconv.Atoi(mathTask[0])
		num2, _ := strconv.Atoi(mathTask[2])
		errCheck := num1 < 0 && num1 > 11 && num2 < 0 && num2 > 11
		if !errCheck {
			fmt.Println(calculate(num1, num2, DetectOperation(mathTask[1])))
		} else {
			panic(RANGE)
		}
	case 1:
		panic(SCALE)
	case 2:
		for i, elem := range mathTask {
			if i == 1 {
				continue
			}
			if val, ok := RomMap[elem]; ok && val > 0 && val < 11 {
				romans = append(romans, val)
			} else {
				panic(RANGE)
			}
		}
		ParseRomToInt(calculate(romans[0], romans[1], DetectOperation(mathTask[1])))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите пример")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		parse(text)
	}
}
