package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReaderValue(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func doCalculate(calculator *calculator) calculator {
	reader := bufio.NewReader(os.Stdin)
	numberOne, _ := ReaderValue("enterFirstNUmber : ", reader)
	numberTwo, _ := ReaderValue("enterSecondeNumber : ", reader)
	calNumb, err1 := strconv.Atoi(numberOne)
	calnum2, err2 := strconv.Atoi(numberTwo)
	if err1 != nil && err2 != nil {
		fmt.Println("not Ok")
	} else {
		calculator.setNumberOne(calNumb)
		calculator.setNumberTwo(calnum2)
	}
	option, _ := ReaderValue("input your oprand a)+   b)-   c)*   d)% : ", reader)
	switch option {
	case "a":
		fmt.Println(calculator.AddNumbers())
	case "b":
		fmt.Println(calculator.MinusNumbers())
	case "c":
		fmt.Println(calculator.MultiplyNumbers())
	case "d":
		fmt.Println(calculator.DivideNumbers())
	default:
		doCalculate(calculator)
	}
	return *calculator
}

func main() {
	calculator := newCalculator()
	calculator = doCalculate(&calculator)

}
